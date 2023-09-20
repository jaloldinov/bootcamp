package handler

import (
	"lesson_20/config"
	"lesson_20/storage"
)

type handler struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewHandler(strg storage.StorageI, conf config.Config) *handler {
	return &handler{strg: strg, cfg: conf}
}
