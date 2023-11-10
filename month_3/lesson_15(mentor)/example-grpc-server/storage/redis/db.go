package redis

import (
	"context"
	"example-grpc-server/config"
	"example-grpc-server/storage"
	"fmt"

	"github.com/go-redis/cache/v9"
	goRedis "github.com/redis/go-redis/v9"
)

type cacheStrg struct {
	db     *cache.Cache
	cacheR *cacheRepo
}

func NewCache(ctx context.Context, cfg config.Config) (storage.CacheI, error) {
	redisClient := goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       2,
	})
	redisCache := cache.New(&cache.Options{
		Redis: redisClient,
	})
	return &cacheStrg{
		db: redisCache,
	}, nil
}

func (d *cacheStrg) Cache() storage.RedisI {
	if d.cacheR == nil {
		d.cacheR = NewCacheRepo(d.db)

	}
	return d.cacheR
}
