package storage

import (
	"context"
	pb "sale_service/genproto"
	"time"
)

type StorageI interface {
	Sale() SaleI
	SaleProduct() SaleProductI
	StaffTransaction() StaffTransactionI
	BranchProductTransactions() BranchPrTransactionI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type SaleI interface {
	CreateSale(context.Context, *pb.CreateSaleRequest) (string, error)
	GetSale(context.Context, *pb.IdRequest) (*pb.Sale, error)
	GetAllSale(context.Context, *pb.ListSaleRequest) (*pb.ListSaleResponse, error)
	UpdateSale(context.Context, *pb.UpdateSaleRequest) (string, error)
	DeleteSale(context.Context, *pb.IdRequest) (string, error)
}

type SaleProductI interface {
	CreateSaleProduct(context.Context, *pb.CreateSaleProductRequest) (string, error)
	GetSaleProduct(context.Context, *pb.IdRequest) (*pb.SaleProduct, error)
	GetAllSaleProduct(context.Context, *pb.ListSaleProductRequest) (*pb.ListSaleProductResponse, error)
	UpdateSaleProduct(context.Context, *pb.UpdateSaleProductRequest) (string, error)
	DeleteSaleProduct(context.Context, *pb.IdRequest) (string, error)
}

type StaffTransactionI interface {
	CreateStaffTransaction(context.Context, *pb.CreateStaffTransactionRequest) (string, error)
	GetStaffTransaction(context.Context, *pb.IdRequest) (*pb.StaffTransaction, error)
	GetAllStaffTransaction(context.Context, *pb.ListStaffTransactionRequest) (*pb.ListStaffTransactionResponse, error)
	UpdateStaffTransaction(context.Context, *pb.UpdateStaffTransactionRequest) (string, error)
	DeleteStaffTransaction(context.Context, *pb.IdRequest) (string, error)
}

type BranchPrTransactionI interface {
	CreateBranchPrTran(context.Context, *pb.CreateBranchPrTransactionRequest) (string, error)
	GetBranchPrTran(context.Context, *pb.IdRequest) (*pb.BranchPrTransaction, error)
	GetAllBranchPrTran(context.Context, *pb.ListBranchPrTransactionRequest) (*pb.ListBranchPrTransactionResponse, error)
	UpdateBranchPrTran(context.Context, *pb.UpdateBranchPrTransactionRequest) (string, error)
	DeleteBranchPrTran(context.Context, *pb.IdRequest) (string, error)
}
