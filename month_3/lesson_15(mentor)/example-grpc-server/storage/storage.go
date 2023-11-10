package storage

import (
	"context"
	sale_service "example-grpc-server/genproto"

	"time"
)

type StorageI interface {
	Branch() BranchI
}
type CacheI interface {
	Cache() RedisI
}
type BranchI interface {
	Create(*sale_service.CreateBranch) (string, error)
	// Update(models.Person) (string, error)
	// Get(req models.RequestByID) (*models.Person, error)
	// // GetByUsername(req models.RequestByUsername) (*models.Person, error)
	// GetAll(models.GetAllPersonsRequest) (*models.GetAllPersonsResponse, error)
	// Delete(req models.RequestByID) (string, error)
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, res interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}
