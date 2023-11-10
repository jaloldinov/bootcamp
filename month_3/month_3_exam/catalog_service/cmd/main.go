package main

import (
	"catalog_service/config"
	"catalog_service/grpc"

	"catalog_service/pkg/logger"
	"catalog_service/storage/postgres"
	"context"
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
	lis, err := net.Listen("tcp", cfg.CatalogServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
