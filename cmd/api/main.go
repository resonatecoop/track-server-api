package main

import (
	"flag"
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	// "user-api/pkg/zerolog"
	"user-api/pkg/mw"
	"user-api/pkg/config"
	"user-api/pkg/postgres"

	"track-server-api/internal/pkg/storage"
	trackDataServer "track-server-api/internal/server"
	rpc "track-server-api/rpc"

)

func main() {
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	router := mux.NewRouter().StrictSlash(true)
	registerRoutes(router, cfg)

	mws := alice.New(mw.CORS, mw.AuthContext, mw.WithURLQuery)

	http.ListenAndServe(cfg.Server.Port, mws.Then(router))
}

func registerRoutes(r *mux.Router, cfg *config.Configuration) {
	db, err := pgsql.New(cfg.DB.Dev.PSN, cfg.DB.Dev.LogQueries, cfg.DB.Dev.TimeoutSeconds)
	checkErr(err)

	sc, err := storage.New(
		cfg.Storage.AccountId,
		cfg.Storage.Key,
		cfg.Storage.AuthEndpoint,
		cfg.Storage.FileEndpoint,
		cfg.Storage.UploadEndpoint,
		cfg.Storage.BucketId,
		cfg.Storage.Timeout)
	checkErr(err)

	server := trackDataServer.NewServer(db, sc)
	trackDataTwirpHandler := rpc.NewTrackDataServiceServer(server, nil))
	r.PathPrefix(rpc.TrackDataServicePathPrefix).Handler(trackDataTwirpHandler)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
