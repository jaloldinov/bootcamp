package api

import (
	"app/config"
	"app/handler"
	"app/storage"
	"net/http"
)

func NewApi(cfg *config.Config, storage storage.StorageI) {
	handler := handler.NewHandler(cfg, storage)

	http.HandleFunc("/branch", handler.BranchHandler)
}
