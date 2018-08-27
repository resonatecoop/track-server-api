package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	pb "track-server-api/rpc"
)

func main() {
	client := pb.NewPlayServiceProtobufClient("http://localhost:8081", &http.Client{})

	s, err := client.Play(context.Background(), &pb.UserTrack{TrackId: "0", UserId: "0"})
	if err != nil {
		fmt.Printf("Error on track play: %v", err)
		os.Exit(1)
	}
}
