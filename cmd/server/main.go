package main

import (
	"context"
	"fmt"
	"net/http"

	playServer "track-server-api/internal/server"
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
	fmt.Printf("Track Play Service on :8081")

	newPlayServer := playServer.NewServer()
	playTwirpHandler := rpc.NewPlayServiceServer(newPlayServer, nil)

	mux := http.NewServeMux()
	mux.Handle(rpc.PlayServicePathPrefix, playTwirpHandler)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8081", handler)
}
