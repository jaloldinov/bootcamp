package handler

import (
	"app/pkg/logger"
	"app/storage"
)

type Handler struct {
	storage storage.StorageI
	log     logger.LoggerI
	redis   storage.RedisI
}

func NewHandler(strg storage.StorageI, strgR storage.RedisI, loger logger.LoggerI) *Handler {
	return &Handler{storage: strg, redis: strgR, log: loger}
}
