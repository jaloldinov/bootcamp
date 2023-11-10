package storage

import (
	pb "branch_service/genproto"
	"context"
	"time"
)

type StorageI interface {
	Branch() BranchI
	BranchProduct() BranchProductI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type BranchI interface {
	CreateBranch(context.Context, *pb.CreateBranchRequest) (string, error)
	GetBranch(context.Context, *pb.IdRequest) (*pb.Branch, error)
	GetAllBranch(context.Context, *pb.ListBranchRequest) (*pb.ListBranchResponse, error)
	UpdateBranch(context.Context, *pb.UpdateBranchRequest) (string, error)
	DeleteBranch(context.Context, *pb.IdRequest) (string, error)
}

type BranchProductI interface {
	CreateBranchProduct(context.Context, *pb.CreateBranchProductRequest) (string, error)
	GetBranchProduct(context.Context, *pb.IdRequest) (*pb.BranchProduct, error)
	GetAllBranchProduct(context.Context, *pb.ListBranchProductRequest) (*pb.ListBranchProductResponse, error)
	UpdateBranchProduct(context.Context, *pb.UpdateBranchProductRequest) (string, error)
	DeleteBranchProduct(context.Context, *pb.IdRequest) (string, error)
}
