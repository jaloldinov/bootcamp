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

type UserService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *UserService {
	return &UserService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *UserService) Create(ctx context.Context, req *user_service.CreateUsersRequest) (*user_service.Response, error) {
	resp, err := b.storage.Users().Create(context.Background(), req)
	if err != nil {
		b.log.Error("error while creating users", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil

}

func (b *UserService) Get(ctx context.Context, req *user_service.IdRequest) (*user_service.Users, error) {
	reso, err := b.storage.Users().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *UserService) List(ctx context.Context, req *user_service.ListUsersRequest) (*user_service.ListUsersResponse, error) {
	Users, err := b.storage.Users().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.ListUsersResponse{Users: Users.Users,
		Count: Users.Count}, nil
}

func (s *UserService) Update(ctx context.Context, req *user_service.UpdateUsersRequest) (*user_service.Response, error) {
	resp, err := s.storage.Users().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (s *UserService) Delete(ctx context.Context, req *user_service.IdRequest) (*user_service.Response, error) {
	resp, err := s.storage.Users().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (b *UserService) GetByLogin(ctx context.Context, req *user_service.IdRequest) (*user_service.Users, error) {
	reso, err := b.storage.Users().GetByLogin(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return reso, nil
}
