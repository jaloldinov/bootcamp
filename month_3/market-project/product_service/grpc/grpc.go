package grpc

import (
	"product_service/config"
	product_service "product_service/genproto"
	"product_service/grpc/service"

	"product_service/pkg/logger"
	"product_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg))
	product_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}
