package trackdataserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/twitchtv/twirp"

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
func (server *Server) StreamTrackData(ctx context.Context, userTrackPB *pb.UserTrack) (<-chan pb.TrackDataOrError, error) {

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

	sc, err := openConnection()
	if err != nil {
		return nil, err
	}

	const bytesPerRead = 12000 // temporary fake value

	trackData := &pb.TrackData{
		StartPosition: 0,
		NumBytes:      bytesPerRead,
	}

	ch := make(chan pb.TrackDataOrError, bytesPerRead)
	go func() {
		defer close(ch)
		for {
			td, err := getTrackData(res.Tracks[0].TrackServerId, trackData, sc)
			select {
			case <-ctx.Done():
				return
			case ch <- pb.TrackDataOrError{Msg: td, Err: err}:
			}
			time.Sleep(5000 * time.Millisecond)
			td.StartPosition += bytesPerRead
		}
	}()

	// TODO: Preauth payment with payment-api

	return ch, nil
}
