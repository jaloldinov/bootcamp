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

type CourierService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	user_service.UnimplementedCourierServiceServer
}

func NewCourierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *CourierService {
	return &CourierService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}
func (b *CourierService) Create(ctx context.Context, req *user_service.CreateCouriersRequest) (*user_service.Response, error) {
	resp, err := b.storage.Couriers().Create(context.Background(), req)
	if err != nil {
		b.log.Error("error while creating courier", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil

}

func (b *CourierService) Get(ctx context.Context, req *user_service.IdRequest) (resp *user_service.Couriers, err error) {
	resp, err = b.storage.Couriers().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *CourierService) List(ctx context.Context, req *user_service.ListCouriersRequest) (*user_service.ListCouriersResponse, error) {
	Couriers, err := b.storage.Couriers().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.ListCouriersResponse{Couriers: Couriers.Couriers,
		Count: Couriers.Count}, nil
}

func (s *CourierService) Update(ctx context.Context, req *user_service.UpdateCouriersRequest) (*user_service.Response, error) {
	resp, err := s.storage.Couriers().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (s *CourierService) Delete(ctx context.Context, req *user_service.IdRequest) (*user_service.Response, error) {
	resp, err := s.storage.Couriers().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &user_service.Response{Message: resp}, nil
}

func (b *CourierService) GetByLogin(ctx context.Context, req *user_service.IdRequest) (resp *user_service.Couriers, err error) {
	resp, err = b.storage.Couriers().GetByLogin(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
