package grpc

import (
	"staff_service/config"
	staff_service "staff_service/genproto"
	"staff_service/grpc/service"

	"staff_service/pkg/logger"
	"staff_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	staff_service.RegisterStaffServiceServer(grpcServer, service.NewStaffService(cfg, log, strg))
	staff_service.RegisterTariffServiceServer(grpcServer, service.NewTariffService(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}
