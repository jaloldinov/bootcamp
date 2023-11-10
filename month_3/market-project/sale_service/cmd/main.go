package main

import (
	"sale_service/config"
	"sale_service/grpc"

	"context"
	"fmt"
	"log"
	"net"
	"sale_service/pkg/logger"
	"sale_service/storage/postgres"
)

func main() {
	cfg := config.Load()
	lg := logger.NewLogger(cfg.Environment, "debug")
	strg, err := postgres.NewStorage(context.Background(), cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	s := grpc.SetUpServer(cfg, lg, strg)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50054))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
