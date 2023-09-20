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

func (h *handler) Category(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		h.CreateCategory(w, r)

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

func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.CreateCategory

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &category)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Category().Create(&category)
	if err != nil {
		h.handlerResponse(w, "error while storage category create:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	resp, err := h.strg.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage category get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	h.handlerResponse(w, "Success", http.StatusOK, resp)

}
func (h *handler) GetListCategory(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.strg.Category().GetList(&models.CategoryGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		h.handlerResponse(w, "error while storage category get list:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "Ssuccess", http.StatusOK, resp)
}

func (h *handler) GetByIDCategory(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while give category id: invalid uuid ", http.StatusBadRequest, nil)
		return
	}

	resp, err := h.strg.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage category get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)

}

func (h *handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	var upCategory models.UpdateCategory

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &upCategory)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Category().Update(&upCategory)
	if err != nil {
		h.handlerResponse(w, "error while storage category upadate:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	resp, err := h.strg.Category().GetByID(&models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage category get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)

}

func (h *handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while delete --> give category id: invalid uuid ", http.StatusBadRequest, nil)
		return
	}

	err := h.strg.Category().Delete(&models.CategoryPrimaryKey{Id: id})

	if err != nil {
		h.handlerResponse(w, "error while delete category :"+err.Error(), http.StatusInternalServerError, nil)
		return
	}
	h.handlerResponse(w, "Success", http.StatusOK, nil)

}
