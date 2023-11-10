package service

import (
	"context"
	"product_service/config"
	product_service "product_service/genproto"
	"product_service/pkg/logger"
	"product_service/storage"
)

type CategoryService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	product_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *CategoryService {
	return &CategoryService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *CategoryService) Create(ctx context.Context, req *product_service.CreateCategoryRequest) (*product_service.CreateCategoryResponse, error) {
	id, err := b.storage.Category().CreateCategory(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.CreateCategoryResponse{Id: id}, nil
}

func (b *CategoryService) Get(ctx context.Context, req *product_service.IdRequest) (*product_service.GetCategoryResponse, error) {
	branch, err := b.storage.Category().GetCategory(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.GetCategoryResponse{Category: branch}, nil
}

func (b *CategoryService) List(ctx context.Context, req *product_service.ListCategoryRequest) (*product_service.ListCategoryResponse, error) {
	branches, err := b.storage.Category().GetAllCategory(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.ListCategoryResponse{Categoryes: branches.Categoryes,
		Count: branches.Count}, nil
}

func (s *CategoryService) Update(ctx context.Context, req *product_service.UpdateCategoryRequest) (*product_service.Response, error) {
	resp, err := s.storage.Category().UpdateCategory(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.Response{Message: resp}, nil
}

func (s *CategoryService) Delete(ctx context.Context, req *product_service.IdRequest) (*product_service.Response, error) {
	resp, err := s.storage.Category().DeleteCategory(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.Response{Message: resp}, nil
}
