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

type ClientsService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	user_service.UnimplementedClientServiceServer
}

func NewClientService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *ClientsService {
	return &ClientsService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *ClientsService) Create(ctx context.Context, req *user_service.CreateClientsRequest) (*user_service.Response, error) {
	resp, err := b.storage.Clients().Create(context.Background(), req)
	if err != nil {
		b.log.Error("error while creating clients", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil

}

func (b *ClientsService) Get(ctx context.Context, req *user_service.IdRequest) (*user_service.Clients, error) {
	reso, err := b.storage.Clients().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *ClientsService) List(ctx context.Context, req *user_service.ListClientsRequest) (*user_service.ListClientsResponse, error) {
	Clients, err := b.storage.Clients().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.ListClientsResponse{Clients: Clients.Clients,
		Count: Clients.Count}, nil
}

func (s *ClientsService) Update(ctx context.Context, req *user_service.UpdateClientsRequest) (*user_service.Response, error) {
	resp, err := s.storage.Clients().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (s *ClientsService) Delete(ctx context.Context, req *user_service.IdRequest) (*user_service.Response, error) {
	resp, err := s.storage.Clients().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (s *ClientsService) UpdateOrder(ctx context.Context, req *user_service.UpdateClientsOrderRequest) (*user_service.Response, error) {
	resp, err := s.storage.Clients().UpdateOrder(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}
