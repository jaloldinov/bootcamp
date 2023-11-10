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

func (b *BranchService) Create(ctx context.Context, req *sale_service.CreateBranchRequest) (*sale_service.CreateBranchResponse, error) {
	id, err := b.storage.Branch().CreateBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.CreateBranchResponse{Id: id}, nil
}

func (b *BranchService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetBranchResponse, error) {
	branch, err := b.storage.Branch().GetBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetBranchResponse{Branch: branch}, nil
}

func (b *BranchService) List(ctx context.Context, req *sale_service.ListBranchRequest) (*sale_service.ListBranchResponse, error) {
	branches, err := b.storage.Branch().GetAllBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.ListBranchResponse{Branches: branches.Branches,
		Count: branches.Count}, nil
}

func (s *BranchService) Update(ctx context.Context, req *sale_service.UpdateBranchRequest) (*sale_service.Response, error) {
	resp, err := s.storage.Branch().UpdateBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *BranchService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.Branch().DeleteBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
