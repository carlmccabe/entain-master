package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"git.neds.sh/matty/entain/api/proto/sport"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	apiEndpoint    = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	racingEndpoint = flag.String("racing-endpoint", "localhost:9000", "racing gRPC server endpoint")
	sportEndpoint  = flag.String("sport-endpoint", "localhost:9001", "sport gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*racingEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}
	if err := sport.RegisterSportHandlerFromEndpoint(
		ctx,
		mux,
		*sportEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}

	log.Printf("API server listening on: %s\n", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}
