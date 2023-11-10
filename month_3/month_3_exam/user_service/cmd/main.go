package main

import (
	"user_service/config"
	"user_service/grpc"

	"context"
	"log"
	"net"
	"user_service/pkg/logger"
	"user_service/storage/postgres"
)

func main() {
	cfg := config.Load()
	lg := logger.NewLogger(cfg.Environment, "debug")
	strg, err := postgres.NewStorage(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	s := grpc.SetUpServer(cfg, lg, strg)
	lis, err := net.Listen("tcp", cfg.UserServicePort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
