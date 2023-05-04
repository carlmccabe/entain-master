package main

import (
	"database/sql"
	"flag"
	"log"
	"net"

	"sport/db"
	"sport/proto/sport"
	"sport/service"

	"google.golang.org/grpc"
)

var (
	sportEndpoint = flag.String("sport-endpoint", "localhost:9001", "sport gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}
}
func run() error {
	conn, err := net.Listen("tcp", ":9001")
	if err != nil {
		return err
	}

	eventsDB, err := sql.Open("sqlite3", "./db/events.db")
	if err != nil {
		return err
	}

	eventsRepo := db.NewEventsRepo(eventsDB)
	if err := eventsRepo.Init(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	sport.RegisterSportServer(
		grpcServer,
		service.NewSportService(
			eventsRepo,
		),
	)

	log.Printf("sport gRPC server listening on: %s\n", *sportEndpoint)

	if err := grpcServer.Serve(conn); err != nil {
		return err
	}

	return nil
}
