package service

import (
	"context"
	"staff_service/config"
	staff_service "staff_service/genproto"
	"staff_service/pkg/logger"
	"staff_service/storage"
)

type TariffService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	staff_service.UnimplementedTariffServiceServer
}

func NewTariffService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *TariffService {
	return &TariffService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *TariffService) Create(ctx context.Context, req *staff_service.CreateTariffRequest) (*staff_service.CreateTariffResponse, error) {
	id, err := b.storage.Tariff().CreateTariff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.CreateTariffResponse{Id: id}, nil
}

func (b *TariffService) Get(ctx context.Context, req *staff_service.IdRequest) (*staff_service.GetTariffResponse, error) {
	branch, err := b.storage.Tariff().GetTariff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.GetTariffResponse{Tariff: branch}, nil
}

func (b *TariffService) List(ctx context.Context, req *staff_service.ListTariffRequest) (*staff_service.ListTariffResponse, error) {
	tariffs, err := b.storage.Tariff().GetAllTariff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.ListTariffResponse{Tariffs: tariffs.Tariffs,
		Count: tariffs.Count}, nil
}

func (s *TariffService) Update(ctx context.Context, req *staff_service.UpdateTariffRequest) (*staff_service.Response, error) {
	resp, err := s.storage.Tariff().UpdateTariff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.Response{Message: resp}, nil
}

func (s *TariffService) Delete(ctx context.Context, req *staff_service.IdRequest) (*staff_service.Response, error) {
	resp, err := s.storage.Tariff().DeleteTariff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.Response{Message: resp}, nil
}
