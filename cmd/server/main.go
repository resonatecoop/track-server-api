package main

import (
	"context"
	"fmt"
	"net/http"

	"track-server-api/internal/database"
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
	testing := true
	db := database.Connect(testing)
	server := trackDataServer.NewServer(db)
	trackDataTwirpHandler := WithURLQuery(rpc.NewTrackDataServiceServer(server, nil))

	mux := http.NewServeMux()
	mux.Handle(rpc.TrackDataServicePathPrefix, trackDataTwirpHandler)

	handler := cors.Default().Handler(mux)

	fmt.Printf("TrackData Service on :8081")

	http.ListenAndServe(":8081", handler)
}
