package grpc

import (
	sale_service "example-grpc-server/genproto"
	grpc_client "example-grpc-server/grpc/client"
	"example-grpc-server/grpc/service"
	"example-grpc-server/pkg/logger"
	"example-grpc-server/storage"

	"google.golang.org/grpc"
)

func SetUpServer(log logger.LoggerI, strg storage.StorageI, grpcClient grpc_client.GrpcClientI) *grpc.Server {
	s := grpc.NewServer()
	sale_service.RegisterBranchServiceServer(s, service.NewBranchService(log, strg, grpcClient))
	return s
}
