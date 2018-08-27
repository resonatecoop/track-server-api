package playserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "track-server-api/rpc"
	track_pb "user-api/rpc/track"
)

// Server implements the PlayService
type Server struct {
}

// NewServer creates an instance of our server
func NewServer() *Server {
	fmt.Printf("new server")
	return &Server{}
}

// Request a track stream
func (server *Server) Play(ctx context.Context, userTrackPB *pb.UserTrack) (<-chan pb.TrackDataOrError, error) {

	// Get track object for TrackServerID and CreatorID

	trackClient := track_pb.NewTrackServiceProtobufClient("http://localhost:8080", &http.Client{})

	track, err := trackClient.GetTrack(context.Background(), &track_pb.Track{Id: userTrackPB.TrackId})
	if err != nil {
		return nil, err
	}

	sc, err := openConnection()
	if err != nil {
		return nil, err
	}

	const bytesPerRead = 12000 // temporary fake value

	trackData := &pb.TrackData{
		TrackServerId: track.TrackServerId,
		StartPosition: 0,
		NumBytes: bytesPerRead,
	}

	ch := make(chan pb.TrackDataOrError,bytesPerRead)
	go func () {
		defer close(ch)
		for {
			td, err := getTrackData(trackData, sc)
			select {
			case <-ctx.Done():
				return
			case ch <- pb.TrackDataOrError{Msg:td, Err:err}:
			}
			time.Sleep(5000 * time.Millisecond)
			td.StartPosition += bytesPerRead
		}
	}()

	// TODO: Preauth payment with payment-api

	return ch, nil
}

