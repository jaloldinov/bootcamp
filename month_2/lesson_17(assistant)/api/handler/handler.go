package handler

import (
	"app/config"
	"app/storage"
	"encoding/json"
	"log"
	"net/http"
)

type handler struct {
	cfg  *config.Config
	strg storage.StorageI
}

type response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, storage storage.StorageI) *handler {
	return &handler{
		cfg:  cfg,
		strg: storage,
	}
}
func (h *handler) handlerResponse(w http.ResponseWriter, msg string,code int, data interface{}) {
	resp := response{
		Status:      code,
		Description: msg,
		Data:        data,
	}

	log.Printf("%+v", resp)

	body,err:= json.Marshal(resp)

	if err!= nil{
		log.Printf("error while marshalling: %+v",err)
	}

	w.WriteHeader(code)
	w.Write([]byte(body))

}