package service

import (
	"context"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type BranchProductTransactionsService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnimplementedBranchPrTransactionServiceServer
}

func NewBranchProductTransactionsService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *BranchProductTransactionsService {
	return &BranchProductTransactionsService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *BranchProductTransactionsService) Create(ctx context.Context, req *sale_service.CreateBranchPrTransactionRequest) (*sale_service.CreateBranchPrTransactionResponse, error) {
	id, err := b.storage.BranchProductTransactions().CreateBranchPrTran(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.CreateBranchPrTransactionResponse{Id: id}, nil
}

func (b *BranchProductTransactionsService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetBranchPrTransactionResponse, error) {
	branch, err := b.storage.BranchProductTransactions().GetBranchPrTran(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetBranchPrTransactionResponse{Branch: branch}, nil
}

func (b *BranchProductTransactionsService) List(ctx context.Context, req *sale_service.ListBranchPrTransactionRequest) (*sale_service.ListBranchPrTransactionResponse, error) {
	sales, err := b.storage.BranchProductTransactions().GetAllBranchPrTran(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.ListBranchPrTransactionResponse{Branches: sales.Branches,
		Count: sales.Count}, nil
}

func (s *BranchProductTransactionsService) Update(ctx context.Context, req *sale_service.UpdateBranchPrTransactionRequest) (*sale_service.Response, error) {
	resp, err := s.storage.BranchProductTransactions().UpdateBranchPrTran(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *BranchProductTransactionsService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.BranchProductTransactions().DeleteBranchPrTran(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
