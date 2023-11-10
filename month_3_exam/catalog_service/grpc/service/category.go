package service

import (
	"catalog_service/config"
	catalog_service "catalog_service/genproto"
	"catalog_service/pkg/logger"
	"catalog_service/storage"
	"context"
)

type CategoryService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	catalog_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *CategoryService {
	return &CategoryService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *CategoryService) Create(ctx context.Context, req *catalog_service.CreateCategoryRequest) (*catalog_service.Response, error) {
	id, err := b.storage.Category().Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &catalog_service.Response{Message: id}, nil
}

func (b *CategoryService) Get(ctx context.Context, req *catalog_service.IdRequest) (*catalog_service.Category, error) {
	reso, err := b.storage.Category().Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return reso, nil
}

func (b *CategoryService) List(ctx context.Context, req *catalog_service.ListCategoryRequest) (*catalog_service.ListCategoryResponse, error) {
	resp, err := b.storage.Category().GetList(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CategoryService) Update(ctx context.Context, req *catalog_service.UpdateCategoryRequest) (*catalog_service.Response, error) {
	resp, err := s.storage.Category().Update(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &catalog_service.Response{Message: resp}, nil
}

func (s *CategoryService) Delete(ctx context.Context, req *catalog_service.IdRequest) (*catalog_service.Response, error) {
	resp, err := s.storage.Category().Delete(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &catalog_service.Response{Message: resp}, nil
}
