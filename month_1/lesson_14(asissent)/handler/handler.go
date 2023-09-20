package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/config"
	"backend_bootcamp_17_07_2023/lesson_14/storage"
)

type handler struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewHandler(strg storage.StorageI, conf config.Config) *handler {
	return &handler{strg: strg, cfg: conf}
}
