package service

import (
	"context"
	"user_service/config"
	user_service "user_service/genproto"
	"user_service/pkg/logger"
	"user_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BranchService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	user_service.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *BranchService {
	return &BranchService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *BranchService) Create(ctx context.Context, req *user_service.CreateBranchRequest) (*user_service.Response, error) {
	resp, err := b.storage.Branch().Create(context.Background(), req)
	if err != nil {
		b.log.Error("error while creating branch", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil

}

func (b *BranchService) Get(ctx context.Context, req *user_service.IdRequest) (*user_service.Branch, error) {
	resp, err := b.storage.Branch().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BranchService) List(ctx context.Context, req *user_service.ListBranchRequest) (*user_service.ListBranchResponse, error) {
	Branchs, err := b.storage.Branch().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return Branchs, nil
}

func (b *BranchService) GetListActive(ctx context.Context, req *user_service.ListBranchActiveRequest) (*user_service.ListBranchResponse, error) {
	Branchs, err := b.storage.Branch().GetListActive(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return Branchs, nil
}

func (s *BranchService) Update(ctx context.Context, req *user_service.UpdateBranchRequest) (*user_service.Response, error) {
	resp, err := s.storage.Branch().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (s *BranchService) Delete(ctx context.Context, req *user_service.IdRequest) (*user_service.Response, error) {
	resp, err := s.storage.Branch().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}
