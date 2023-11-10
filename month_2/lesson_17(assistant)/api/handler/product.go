package handler

import (
	"app/config"
	"app/models"
	"app/pkg/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handler) Product(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		h.CreateProduct(w, r)

	case "GET":
		fmt.Println(r.URL.Query().Get("method"))
		var (
			method = r.URL.Query().Get("method")
		)

		if method == "GET_LIST" {
			h.GetListCategory(w, r)

		} else if method == "GET" {
			h.GetByIDCategory(w, r)
		}
	case "PUT":
		h.UpdateCategory(w, r)
	case "DELETE":
		h.DeleteCategory(w, r)
	}

}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.CreateProduct

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Product().Create(&product)
	if err != nil {
		h.handlerResponse(w, "error while storage product create:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	resp, err := h.strg.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage product get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	h.handlerResponse(w, "success", http.StatusOK, resp)

}

func (h *handler) GetListProduct(w http.ResponseWriter, r *http.Request) {
	var (
		offsetStr       = r.URL.Query().Get("offset")
		limitStr        = r.URL.Query().Get("limit")
		search          = r.URL.Query().Get("search")
		offset    int   = config.Load().DefaultOffset
		limit     int   = config.Load().DefaultLimit
		err       error = nil
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			h.handlerResponse(w, "error while offset: "+err.Error(), http.StatusBadRequest, nil)
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			h.handlerResponse(w, "error while limit: "+err.Error(), http.StatusBadRequest, nil)
			return
		}
	}

	resp, err := h.strg.Product().GetList(&models.ProductGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		h.handlerResponse(w, "error while storage product get list:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "success", http.StatusOK, resp)
}

func (h *handler) GetByIDProduct(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while give product id: invalid uuid ", http.StatusBadRequest, nil)
		return
	}

	resp, err := h.strg.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage product get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "success", http.StatusOK, resp)

}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var upProduct models.UpdateProduct

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &upProduct)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Product().Update(&upProduct)
	if err != nil {
		h.handlerResponse(w, "error while storage product upadate:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	resp, err := h.strg.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage product get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "success", http.StatusOK, resp)

}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while delete --> give product id: invalid uuid ", http.StatusBadRequest, nil)
		return
	}

	err := h.strg.Product().Delete(&models.ProductPrimaryKey{Id: id})

	if err != nil {
		h.handlerResponse(w, "error while delete product :"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	h.handlerResponse(w, "success", http.StatusOK, nil)

}
