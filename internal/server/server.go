package trackdataserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"

	"track-server-api/internal/database/models"
	pb "track-server-api/rpc"
	track_pb "user-api/rpc/track"

	"github.com/go-pg/pg"
)

// Server implements the TrackDataService
type Server struct {
	db *pg.DB
}

// NewServer creates an instance of our server
func NewServer(db *pg.DB) *Server {
	fmt.Printf("new server")
	return &Server{db: db}
}

// Request a track stream
func (server *Server) StreamTrackData(ctx context.Context, userTrackPB *pb.UserTrack) (<-chan pb.TrackChunkOrError, error) {

	// Get track object for TrackServerID and CreatorID

	trackClient := track_pb.NewTrackServiceProtobufClient("http://localhost:8080", &http.Client{})

	track := &track_pb.Track{Id: userTrackPB.TrackId}
	tracks := &track_pb.TracksList{Tracks: []*track_pb.Track{track}}
	res, err := trackClient.GetTracks(context.Background(), tracks)
	if err != nil {
		return nil, err
	}
	if len(res.Tracks) == 0 {
		return nil, twirp.NotFoundError("track id not found")
	}

	sc, err := OpenStorageConnection()
	if err != nil {
		return nil, err
	}

	const bytesPerRead = 12000 // temporary fake value

	trackChunk := &pb.TrackChunk{
		StartPosition: 0,
		NumBytes:      bytesPerRead,
	}

	ch := make(chan pb.TrackChunkOrError, bytesPerRead)
	go func() {
		defer close(ch)
		for {
			td, err := GetTrackChunkFromStorage(res.Tracks[0].TrackServerId, trackChunk, sc)
			select {
			case <-ctx.Done():
				return
			case ch <- pb.TrackChunkOrError{Msg: td, Err: err}:
			}
			time.Sleep(5000 * time.Millisecond)
			td.StartPosition += bytesPerRead
		}
	}()

	// TODO: Preauth payment with payment-api

	return ch, nil
}

// Upload a track stream
func (server *Server) UploadTrackData(ctx context.Context, trackChunks <-chan pb.TrackChunkOrError) (*pb.TrackServerId, error) {

	newTrackData := &models.TrackData{TrackId: uuid.NewV4(), UserId: uuid.NewV4()}
	err := server.db.Insert(newTrackData)
	if err != nil {
		return nil, err
	}

	return &pb.TrackServerId{TrackServerId: newTrackData.Id.String()}, nil
}
