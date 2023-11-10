package service

import (
	"context"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type StaffTransactionService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnsafeStaffTransactionServiceServer
}

func NewStaffTransactionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *StaffTransactionService {
	return &StaffTransactionService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *StaffTransactionService) Create(ctx context.Context, req *sale_service.CreateStaffTransactionRequest) (*sale_service.CreateStaffTransactionResponse, error) {
	id, err := b.storage.StaffTransaction().CreateStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.CreateStaffTransactionResponse{Id: id}, nil
}

func (b *StaffTransactionService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetStaffTransactionResponse, error) {
	branch, err := b.storage.StaffTransaction().GetStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetStaffTransactionResponse{StaffTransaction: branch}, nil
}

func (b *StaffTransactionService) List(ctx context.Context, req *sale_service.ListStaffTransactionRequest) (*sale_service.ListStaffTransactionResponse, error) {
	sales, err := b.storage.StaffTransaction().GetAllStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.ListStaffTransactionResponse{StaffTransactions: sales.StaffTransactions,
		Count: sales.Count}, nil
}

func (s *StaffTransactionService) Update(ctx context.Context, req *sale_service.UpdateStaffTransactionRequest) (*sale_service.Response, error) {
	resp, err := s.storage.StaffTransaction().UpdateStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *StaffTransactionService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.StaffTransaction().DeleteStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
