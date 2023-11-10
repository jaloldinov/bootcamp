package service

import (
	"catalog_service/config"
	catalog_service "catalog_service/genproto"
	"catalog_service/pkg/logger"
	"catalog_service/storage"
	"context"
)

type ProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	catalog_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *ProductService {
	return &ProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *ProductService) Create(ctx context.Context, req *catalog_service.CreateProductRequest) (*catalog_service.Response, error) {
	id, err := b.storage.Product().Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &catalog_service.Response{Message: id}, nil
}

func (b *ProductService) Get(ctx context.Context, req *catalog_service.IdRequest) (*catalog_service.Product, error) {
	reso, err := b.storage.Product().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *ProductService) List(ctx context.Context, req *catalog_service.ListProductRequest) (*catalog_service.ListProductResponse, error) {
	Products, err := b.storage.Product().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return Products, nil
}

func (s *ProductService) Update(ctx context.Context, req *catalog_service.UpdateProductRequest) (*catalog_service.Response, error) {
	resp, err := s.storage.Product().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &catalog_service.Response{Message: resp}, nil
}

func (s *ProductService) Delete(ctx context.Context, req *catalog_service.IdRequest) (*catalog_service.Response, error) {
	resp, err := s.storage.Product().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &catalog_service.Response{Message: resp}, nil
}
