package main

import (
	"app/api"
	"app/api/handler"
	"app/config"
	"app/storage/postgres"
	"app/storage/redis"
	"fmt"

	"app/pkg/logger"
	"context"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := postgres.NewStorage(context.Background(), cfg)
	if err != nil {
		return
	}
	redisStrg, err := redis.NewCache(context.Background(), cfg)
	if err != nil {
		return
	}

	h := handler.NewHandler(strg, redisStrg, log)

	r := api.NewServer(h)
	r.Run(fmt.Sprintf(":%s", cfg.Port))

}
