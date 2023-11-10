package main

import (
	"branch_service/config"
	"branch_service/grpc"

	"branch_service/pkg/logger"
	"branch_service/storage/postgres"
	"context"
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

	s := grpc.SetUpServer(cfg, lg, strg)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
