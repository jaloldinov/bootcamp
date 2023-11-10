package handler

import (
	"app/models"
	"app/pkg/helper"
	"encoding/json"
	"io"
	"net/http"
)

func (h *handler) BranchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateBranch(w, r)

	case "GET":
		var (
			method = r.URL.Query().Get("method")
		)

		if method == "GET_LIST" {
			// h.GetAllBranch(w, r)

		} else if method == "GET" {
			h.GetBranch(w, r)
		}
	case "PUT":
		// h.UpdateBranch(w, r)
	case "DELETE":
		// h.DeleteBranch(w, r)
	}
}

func (h *handler) CreateBranch(w http.ResponseWriter, r *http.Request) {
	var branch models.CreateBranch
	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.handlerResponse(w, "error while read body:  "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = json.Unmarshal(body, &branch)
	if err != nil {
		h.handlerResponse(w, "error while unmarshal body: "+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Branch().CreateBranch(&branch)
	if err != nil {
		h.handlerResponse(w, "error while storage branch create:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "success", http.StatusOK, id)
}

func (h *handler) GetBranch(w http.ResponseWriter, r *http.Request) {

	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "error while give branch id: invalid uuid ", http.StatusBadRequest, nil)
		return
	}

	resp, err := h.strg.Branch().GetBranch(&models.IdRequest{Id: id})
	if err != nil {
		h.handlerResponse(w, "error while storage branch get by id:"+err.Error(), http.StatusInternalServerError, nil)
		return
	}

	h.handlerResponse(w, "success", http.StatusOK, resp)

}

// func (h *handler) GetAllBranch(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		offsetStr       = r.URL.Query().Get("offset")
// 		limitStr        = r.URL.Query().Get("limit")
// 		search          = r.URL.Query().Get("search")
// 		offset    int   = config.Load().DefaultOffset
// 		limit     int   = config.Load().DefaultLimit
// 		err       error = nil
// 	)

// 	if offsetStr != "" {
// 		offset, err = strconv.Atoi(offsetStr)
// 		if err != nil {
// 			h.handlerResponse(w, "error while offset: "+err.Error(), http.StatusBadRequest, nil)
// 			return
// 		}
// 	}

// 	if limitStr != "" {
// 		limit, err = strconv.Atoi(limitStr)
// 		if err != nil {
// 			h.handlerResponse(w, "error while limit: "+err.Error(), http.StatusBadRequest, nil)
// 			return
// 		}
// 	}

// 	resp, err := h.strg.Branch().GetAllBranch(&models.GetAllBranchRequest{
// 		Page:  offset,
// 		Limit: limit,
// 		Name:  search,
// 	})

// 	if err != nil {
// 		h.handlerResponse(w, "error while storage branch get list:"+err.Error(), http.StatusInternalServerError, nil)
// 		return
// 	}

// 	h.handlerResponse(w, "success", http.StatusOK, resp)
// }

// func (h *handler) UpdateBranch(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to read request body", http.StatusBadRequest)
// 		return
// 	}

// 	var branch models.Branch
// 	err = json.Unmarshal(body, &branch)
// 	if err != nil {
// 		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
// 		return
// 	}

// 	resp, err := h.strg.Branch().UpdateBranch(branch)
// 	if err != nil {
// 		http.Error(w, "Failed to update branch", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "Updated branch with ID: %s", resp)
// }

// func (h *handler) DeleteBranch(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Path[len("/branch/"):]
// 	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{
// 		Id: id,
// 	})

// 	if err != nil {
// 		http.Error(w, "Failed to delete branch", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "deleted branch with ID: %s", resp)
// }
