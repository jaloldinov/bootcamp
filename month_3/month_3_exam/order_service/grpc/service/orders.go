package service

import (
	"context"
	"order_service/config"
	order_service "order_service/genproto"
	"order_service/pkg/logger"
	"order_service/storage"
)

type OrderService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *OrderService {
	return &OrderService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *OrderService) Create(ctx context.Context, req *order_service.CreateOrderRequest) (*order_service.Response, error) {
	id, err := b.storage.Order().Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: id}, nil
}

func (b *OrderService) Get(ctx context.Context, req *order_service.IdRequest) (*order_service.Order, error) {
	reso, err := b.storage.Order().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *OrderService) List(ctx context.Context, req *order_service.ListOrderRequest) (*order_service.ListOrderResponse, error) {
	Orders, err := b.storage.Order().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.ListOrderResponse{Orders: Orders.Orders,
		Count: Orders.Count}, nil
}

func (s *OrderService) Update(ctx context.Context, req *order_service.UpdateOrderRequest) (*order_service.Response, error) {
	resp, err := s.storage.Order().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}

func (s *OrderService) UpdateStatus(ctx context.Context, req *order_service.UpdateOrderStatusRequest) (*order_service.Response, error) {
	resp, err := s.storage.Order().UpdateStatus(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}

func (s *OrderService) Delete(ctx context.Context, req *order_service.IdRequest) (*order_service.Response, error) {
	resp, err := s.storage.Order().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: resp}, nil
}
