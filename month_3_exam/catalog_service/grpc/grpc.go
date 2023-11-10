package grpc

import (
	"catalog_service/config"
	catalog_service "catalog_service/genproto"
	"catalog_service/grpc/service"

	"catalog_service/pkg/logger"
	"catalog_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	catalog_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg))
	catalog_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
