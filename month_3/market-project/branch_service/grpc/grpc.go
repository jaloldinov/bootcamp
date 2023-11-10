package grpc

import (
	"branch_service/config"
	branch_service "branch_service/genproto"

	"branch_service/grpc/service"
	"branch_service/pkg/logger"
	"branch_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	branch_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg))
	branch_service.RegisterBranchProductServiceServer(grpcServer, service.NewBranchProductService(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}
