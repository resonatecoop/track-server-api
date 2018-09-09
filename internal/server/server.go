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

	"track-server-api/internal"
	"track-server-api/internal/database/models"
	pb "track-server-api/rpc"

	"github.com/go-pg/pg"
)

const Bitrate int32 = 96000
const SecondsPerRead time.Duration = 5
const BytesPerRead int32 = Bitrate * 8 / int32(SecondsPerRead)

// Server implements the TrackDataService
type Server struct {
	db *pg.DB
}

// NewServer creates an instance of our server
func NewServer(db *pg.DB) *Server {
	//fmt.Printf("new server")

	return &Server{db: db}
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

	trackId, twerr := internal.GetUuidFromString(userTrackPB.TrackId)
	if twerr != nil {
		return nil, twerr
	}

	_, twerr = internal.GetUuidFromString(userTrackPB.UserId)
	if twerr != nil {
		return nil, twerr
	}

	trackData := &models.TrackData{TrackId: trackId}
	pgerr := server.db.Model(trackData).Where("track_id = ?", trackId).Select()
	if pgerr != nil {
		return nil, internal.CheckError(pgerr, "track_data")
	}

	sc, err := OpenStorageConnection()
	if err != nil {
		return nil, err
	}

	trackChunk := &pb.TrackChunk{
		StartPosition: 0,
		NumBytes:      BytesPerRead,
	}

	ch := make(chan pb.TrackChunkOrError, BytesPerRead)
	go func() {
		defer close(ch)
		for {
			td, err := GetTrackChunkFromStorage(trackData.StorageId, trackChunk, sc)
			if err == io.EOF {
				break
			}
			select {
			case <-ctx.Done():
				return
			case ch <- pb.TrackChunkOrError{Msg: td, Err: err}:
			}
			time.Sleep(SecondsPerRead * time.Second)
			trackChunk.StartPosition += BytesPerRead
		}
	}()

	// TODO: Preauth payment with payment-api

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

	sc, err := OpenStorageConnection()
	if err != nil {
		return nil, err
	}

	uploadUrl, err := GetUploadUrl(sc)
	if err != nil {
		return nil, err
	}

	storageFileInfo, err := UploadTrackToStorage(newTrackUpload, uploadUrl, sc)
	if err != nil {
		return nil, err
	}

	userId, twerr := internal.GetUuidFromString(newTrackUpload.UserId)
	if twerr != nil {
		return nil, twerr
	}

	newTrackData := &models.TrackData{
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
