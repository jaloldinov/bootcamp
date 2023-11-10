package service

import (
	"context"
	"order_service/config"
	order_service "order_service/genproto"
	"order_service/pkg/logger"
	"order_service/storage"
)

type DeliveryTariffService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	order_service.UnimplementedDeliveryTariffServiceServer
}

func NewDeliveryTariffService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *DeliveryTariffService {
	return &DeliveryTariffService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *DeliveryTariffService) Create(ctx context.Context, req *order_service.CreateDeliveryTariffRequest) (*order_service.Response, error) {
	id, err := b.storage.DeliveryTariff().Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: id}, nil
}

func (b *DeliveryTariffService) Get(ctx context.Context, req *order_service.IdRequest) (*order_service.DeliveryTariff, error) {
	reso, err := b.storage.DeliveryTariff().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *DeliveryTariffService) List(ctx context.Context, req *order_service.ListDeliveryTariffRequest) (*order_service.ListDeliveryTariffResponse, error) {
	DeliveryTariffs, err := b.storage.DeliveryTariff().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.ListDeliveryTariffResponse{DeliveryTariffs: DeliveryTariffs.DeliveryTariffs,
		Count: DeliveryTariffs.Count}, nil
}

func (s *DeliveryTariffService) Update(ctx context.Context, req *order_service.UpdateDeliveryTariffRequest) (*order_service.Response, error) {
	resp, err := s.storage.DeliveryTariff().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}

func (s *DeliveryTariffService) Delete(ctx context.Context, req *order_service.IdRequest) (*order_service.Response, error) {
	resp, err := s.storage.DeliveryTariff().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}
