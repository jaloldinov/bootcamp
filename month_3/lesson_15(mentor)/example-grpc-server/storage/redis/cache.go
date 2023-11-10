package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/pkg/errors"
)

type cacheRepo struct {
	cache *cache.Cache
}

func NewCacheRepo(cache *cache.Cache) *cacheRepo {
	return &cacheRepo{cache: cache}
}

func (u cacheRepo) Create(ctx context.Context, id string, obj interface{}, ttl time.Duration) error {
	err := u.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   id,
		Value: obj,
		TTL:   ttl,
	})
	if err != nil {
		println("redis.Create.Error:", err.Error(), "\nkey:", id)
		return errors.Wrap(err, "error while creating cache in redis")
	}

	fmt.Println("create in redis", id)
	return nil
}

func (u cacheRepo) Get(ctx context.Context, id string, response interface{}) (bool, error) {
	// var response interface{}

	err := u.cache.Get(ctx, id, response)
	if err != nil {
		println("redis.Get.Error:", err.Error(), "\nkey:", id)
		return false, err
	}
	fmt.Println("get from redis", id)
	return true, nil
}

func (u cacheRepo) Delete(ctx context.Context, id string) error {

	err := u.cache.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "error while deleting cache in redis")
	}
	fmt.Println("delete from redis", id)
	return nil
}
