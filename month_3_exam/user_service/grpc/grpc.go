package grpc

import (
	"user_service/config"
	user_service "user_service/genproto"

	"user_service/grpc/service"
	"user_service/pkg/logger"
	"user_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg))
	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg))
	user_service.RegisterCourierServiceServer(grpcServer, service.NewCourierService(cfg, log, strg))
	user_service.RegisterClientServiceServer(grpcServer, service.NewClientService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
