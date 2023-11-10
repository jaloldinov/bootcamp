package main

import (
	"context"
	"example-grpc-server/config"
	"example-grpc-server/grpc"
	grpc_client "example-grpc-server/grpc/client"
	"example-grpc-server/pkg/logger"
	"example-grpc-server/storage/postgres"
	"fmt"
	"log"
	"net"
)

func main() {
	cfg := config.Load()
	lg := logger.NewLogger(cfg.Environment, "debug")
	strg, err := postgres.NewStorage(context.Background(), cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	clients, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to services: %v", err)
	}

	s := grpc.SetUpServer(lg, strg, clients)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
