package handler

import "playground/cpp-bootcamp/storage"

type handler struct {
	storage storage.StorageI
}

func NewHandler(strg storage.StorageI) *handler {
	return &handler{storage: strg}
}
