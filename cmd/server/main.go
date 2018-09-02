package main

import (
	"context"
	"fmt"
	"net/http"

	trackDataServer "track-server-api/internal/server"
	rpc "track-server-api/rpc"

	"github.com/rs/cors"
)

func WithURLQuery(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		query := r.URL.Query()
		ctx = context.WithValue(ctx, "query", query)
		r = r.WithContext(ctx)

		base.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Printf("TrackData Service on :8081")

	newTrackDataServer := trackDataServer.NewServer()
	trackDataTwirpHandler := rpc.NewTrackDataServiceServer(newTrackDataServer, nil)

	mux := http.NewServeMux()
	mux.Handle(rpc.TrackDataServicePathPrefix, trackDataTwirpHandler)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8081", handler)
}
