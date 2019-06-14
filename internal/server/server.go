package trackdataserver

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	uuid "github.com/satori/go.uuid"
	errorpkg "user-api/pkg/error"
	uuidpkg "user-api/pkg/uuid"

	"track-server-api/internal/model"
	"track-server-api/internal/pkg/storage"
	pb "track-server-api/rpc"

	"github.com/go-pg/pg"
)

const Bitrate int32 = 96000
const SecondsPerRead time.Duration = 5
const BytesPerRead int32 = Bitrate * 8 / int32(SecondsPerRead)

const SecondsBeforeAuthRequired = 10

// Server implements the TrackDataService
type Server struct {
	db *pg.DB
	sc Storage
}

type Storage interface {
	GetTrackChunkFromStorage(string, *pb.TrackChunk) (*pb.TrackChunk, error)
	GetUploadUrl() (*storage.UploadUrl, error)
	UploadTrackToStorage(*pb.TrackUpload, *storage.UploadUrl) (*storage.StorageFileInfo, error)
}

// NewServer creates an instance of our server
func NewServer(db *pg.DB, sc Storage) *Server {
	return &Server{db: db, sc: sc}
}

// Request a track stream
func (server *Server) StreamTrackData(ctx context.Context, userTrackPB *pb.UserTrack) (<-chan pb.TrackChunkOrError, error) {

	// Get Track from TrackService for CreatorID

	// How will we do multi-API setup for testing? ??

	// trackClient := track_pb.NewTrackServiceProtobufClient("http://localhost:8080", &http.Client{})
	// track := &track_pb.Track{Id: userTrackPB.TrackId}
	// tracks := &track_pb.TracksList{Tracks: []*track_pb.Track{track}}
	// res, err := trackClient.GetTracks(context.Background(), tracks)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(res.Tracks) == 0 {
	// 	return nil, twirp.NotFoundError("track id not found")
	// }

	trackId, twerr := uuidpkg.GetUuidFromString(userTrackPB.TrackId)
	if twerr != nil {
		return nil, twerr
	}

	_, twerr = uuidpkg.GetUuidFromString(userTrackPB.UserId)
	if twerr != nil {
		return nil, twerr
	}

	trackData := &model.TrackData{TrackId: trackId}
	pgerr := server.db.Model(trackData).Where("track_id = ?", trackId).Select()
	if pgerr != nil {
		return nil, errorpkg.CheckError(pgerr, "track_data")
	}

	trackChunk := &pb.TrackChunk{
		StartPosition: 0,
		NumBytes:      BytesPerRead,
	}

	//auth_ch := make(chan bool)
	auth := false
	start := time.Now()
	prev := time.Now()
	once := false

	ch := make(chan pb.TrackChunkOrError)
	go func() {
		defer close(ch)

		go func() {
			// request pre-auth
			time.Sleep(13 * time.Second)
			auth = true
		}()

		for {
			td, err := server.sc.GetTrackChunkFromStorage(trackData.StorageId, trackChunk)
			if err == io.EOF {
				break
			}
			select {
			case <-ctx.Done():
				return
			case ch <- pb.TrackChunkOrError{Msg: td, Err: err}:
			}

			t := time.Now()
			elapsed := t.Sub(prev)
			next := t.Sub(start) + elapsed

			fmt.Printf("auth %v next %v t\n", auth, next)

			if !auth && next > SecondsBeforeAuthRequired*time.Second {
				return
			}

			if once {
				time.Sleep(SecondsPerRead*time.Second - elapsed)
			}
			once = true

			trackChunk.StartPosition += BytesPerRead
		}
	}()

	return ch, nil
}

// Upload a track stream
func (server *Server) UploadTrackData(ctx context.Context, trackUpload *pb.TrackUpload) (*pb.TrackServerId, error) {

	tmpfile, err := ioutil.TempFile("", trackUpload.Name)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(trackUpload.Data); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}

	cmdStr := "ffmpeg -y -v error -i %s -c:a aac -vn -movflags +faststart -ar 48000 -b:a 96k %s"
	outFileName := fmt.Sprintf("%s.m4a", trackUpload.TrackId)
	cmd := exec.Command("sh", "-c", fmt.Sprintf(cmdStr, tmpfile.Name(), outFileName))

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ffmpeg error: " + fmt.Sprint(err) + ": " + stderr.String())
	}

	defer os.Remove(outFileName)

	data, err := ioutil.ReadFile(outFileName)
	if err != nil {
		return nil, err
	}

	newTrackUpload := &pb.TrackUpload{
		Name:    outFileName,
		UserId:  trackUpload.UserId,
		TrackId: trackUpload.TrackId,
		Data:    data,
	}

	uploadUrl, err := server.sc.GetUploadUrl()
	if err != nil {
		return nil, err
	}

	storageFileInfo, err := server.sc.UploadTrackToStorage(newTrackUpload, uploadUrl)
	if err != nil {
		return nil, err
	}

	userId, twerr := uuidpkg.GetUuidFromString(newTrackUpload.UserId)
	if twerr != nil {
		return nil, twerr
	}

	newTrackData := &model.TrackData{
		TrackId:   uuid.NewV4(), // this should be linked to user-track api, but how?
		UserId:    userId,
		StorageId: storageFileInfo.FileId,
	}

	err = server.db.Insert(newTrackData)
	if err != nil {
		return nil, err
	}

	return &pb.TrackServerId{TrackServerId: newTrackData.Id.String()}, nil
}
