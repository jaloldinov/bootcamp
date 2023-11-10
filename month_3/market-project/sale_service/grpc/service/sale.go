package service

import (
	"context"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type SaleService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnsafeSaleServiceServer
}

func NewSaleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *SaleService {
	return &SaleService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *SaleService) Create(ctx context.Context, req *sale_service.CreateSaleRequest) (*sale_service.CreateSaleResponse, error) {
	id, err := b.storage.Sale().CreateSale(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.CreateSaleResponse{Id: id}, nil
}

func (b *SaleService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetSaleResponse, error) {
	branch, err := b.storage.Sale().GetSale(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetSaleResponse{Sale: branch}, nil
}

func (b *SaleService) List(ctx context.Context, req *sale_service.ListSaleRequest) (*sale_service.ListSaleResponse, error) {
	sales, err := b.storage.Sale().GetAllSale(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.ListSaleResponse{Sales: sales.Sales,
		Count: sales.Count}, nil
}

func (s *SaleService) Update(ctx context.Context, req *sale_service.UpdateSaleRequest) (*sale_service.Response, error) {
	resp, err := s.storage.Sale().UpdateSale(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *SaleService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.Sale().DeleteSale(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
