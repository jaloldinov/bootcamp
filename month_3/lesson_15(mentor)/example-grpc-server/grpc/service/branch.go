package service

import (
	"context"
	sale_service "example-grpc-server/genproto"
	grpc_client "example-grpc-server/grpc/client"
	"example-grpc-server/pkg/logger"
	"example-grpc-server/storage"
)

type BranchService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	clients grpc_client.GrpcClientI
	sale_service.UnimplementedBranchServiceServer
}

func NewBranchService(log logger.LoggerI, strg storage.StorageI, grpcClients grpc_client.GrpcClientI) *BranchService {
	return &BranchService{
		logger:  log,
		storage: strg,
		clients: grpcClients,
	}
}

func (b *BranchService) Create(ctx context.Context, req *sale_service.CreateBranch) (*sale_service.CreateResponse, error) {
	// id, err := b.storage.Branch().Create(req)
	// if err != nil {
	// 	return nil, err
	// }

	return &sale_service.CreateResponse{Id: "hello world"}, nil
}
