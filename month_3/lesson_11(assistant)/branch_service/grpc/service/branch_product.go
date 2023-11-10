package service

import (
	"branch_service/config"
	branch_service "branch_service/genproto"
	"branch_service/pkg/logger"
	"branch_service/storage"
	"context"
)

type BranchProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	branch_service.UnimplementedBranchProductServiceServer
}

func NewBranchProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *BranchProductService {
	return &BranchProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *BranchProductService) Create(ctx context.Context, req *branch_service.CreateBranchProductRequest) (*branch_service.CreateBranchProductResponse, error) {
	id, err := b.storage.BranchProduct().CreateBranchProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.CreateBranchProductResponse{Id: id}, nil
}

func (b *BranchProductService) Get(ctx context.Context, req *branch_service.IdRequest) (*branch_service.GetBranchProductResponse, error) {
	branch, err := b.storage.BranchProduct().GetBranchProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.GetBranchProductResponse{BranchProduct: branch}, nil
}

func (b *BranchProductService) List(ctx context.Context, req *branch_service.ListBranchProductRequest) (*branch_service.ListBranchProductResponse, error) {
	branches, err := b.storage.BranchProduct().GetAllBranchProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.ListBranchProductResponse{BranchProductes: branches.BranchProductes,
		Count: branches.Count}, nil
}

func (s *BranchProductService) Update(ctx context.Context, req *branch_service.UpdateBranchProductRequest) (*branch_service.Response, error) {
	resp, err := s.storage.BranchProduct().UpdateBranchProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.Response{Message: resp}, nil
}

func (s *BranchProductService) Delete(ctx context.Context, req *branch_service.IdRequest) (*branch_service.Response, error) {
	resp, err := s.storage.BranchProduct().DeleteBranchProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.Response{Message: resp}, nil
}
