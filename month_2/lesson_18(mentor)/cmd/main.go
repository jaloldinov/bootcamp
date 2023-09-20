package main

import (
	"app/api"
	"app/api/handler"
	"app/config"
	"app/storage/postgres"

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

	h := handler.NewHandler(strg, log)

	r := api.NewServer(h)
	r.Run(":8000")
}
