package service

import (
	"context"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type SaleProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnsafeSaleProductServiceServer
}

func NewSaleProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *SaleProductService {
	return &SaleProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *SaleProductService) Create(ctx context.Context, req *sale_service.CreateSaleProductRequest) (*sale_service.CreateSaleProductResponse, error) {
	id, err := b.storage.SaleProduct().CreateSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.CreateSaleProductResponse{Id: id}, nil
}

func (b *SaleProductService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetSaleProductResponse, error) {
	branch, err := b.storage.SaleProduct().GetSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetSaleProductResponse{SaleProduct: branch}, nil
}

func (b *SaleProductService) List(ctx context.Context, req *sale_service.ListSaleProductRequest) (*sale_service.ListSaleProductResponse, error) {
	sales, err := b.storage.SaleProduct().GetAllSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.ListSaleProductResponse{SaleProducts: sales.SaleProducts,
		Count: sales.Count}, nil
}

func (s *SaleProductService) Update(ctx context.Context, req *sale_service.UpdateSaleProductRequest) (*sale_service.Response, error) {
	resp, err := s.storage.SaleProduct().UpdateSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *SaleProductService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.SaleProduct().DeleteSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
