package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	pb "track-server-api/rpc"
)

func main() {
	client := pb.NewTrackDataServiceProtobufClient("http://localhost:8081", &http.Client{})

	trackServerId, err := client.UploadTrackData(context.Background(), &pb.TrackChunk{NumBytes: 100})
	if err != nil {
		fmt.Printf("Error on UploadTrackData: %v\n", err)
		os.Exit(1)
	}

	id := trackServerId.TrackServerId
	id = "bcdbd202-ffec-44cf-8dc0-f5cc8092961e"
	fmt.Printf("id: %v\n", id)

	trackChunkStream, err := client.StreamTrackData(context.Background(), &pb.UserTrack{TrackId: id, UserId: id})
	if err != nil {
		fmt.Printf("Error on StreamTrackData: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("ch: %v\n", trackChunkStream)

	for trackChunkOrErr := range trackChunkStream {
		if trackChunkOrErr.Err != nil {
			fmt.Printf("chunk err: %v\n", trackChunkOrErr.Err)
		}
		tc := trackChunkOrErr.Msg

		fmt.Printf("got chunk: %d %d\n", tc.StartPosition, tc.NumBytes)
	}
}
