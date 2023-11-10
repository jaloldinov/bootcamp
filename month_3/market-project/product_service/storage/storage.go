package storage

import (
	"context"
	pb "product_service/genproto"
	"time"
)

type StorageI interface {
	Product() ProductI
	Category() CategoryI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type ProductI interface {
	CreateProduct(context.Context, *pb.CreateProductRequest) (string, error)
	GetProduct(context.Context, *pb.IdRequest) (*pb.Product, error)
	GetAllProduct(context.Context, *pb.ListProductRequest) (*pb.ListProductResponse, error)
	UpdateProduct(context.Context, *pb.UpdateProductRequest) (string, error)
	DeleteProduct(context.Context, *pb.IdRequest) (string, error)
}

type CategoryI interface {
	CreateCategory(context.Context, *pb.CreateCategoryRequest) (string, error)
	GetCategory(context.Context, *pb.IdRequest) (*pb.Category, error)
	GetAllCategory(context.Context, *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error)
	UpdateCategory(context.Context, *pb.UpdateCategoryRequest) (string, error)
	DeleteCategory(context.Context, *pb.IdRequest) (string, error)
}
