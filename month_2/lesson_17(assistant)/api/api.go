package api

import (
	"app/api/handler"
	"app/config"
	"app/storage"
	"net/http"
)

func NewApi(cfg *config.Config, storage storage.StorageI) {
	handler := handler.NewHandler(cfg, storage)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/product", handler.Product)

}
