package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson_20/models"
	"net/http"
	"strconv"
)

func (h *handler) BranchHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		h.CreateBranch(w, r)
	case http.MethodGet:
		if r.URL.Path[len("/branch")] == '/' && r.URL.Path[len("/branch")+1:] != "" {
			h.GetBranch(w, r)
		}
		if r.URL.Path == "/branch/" {
			h.GetAllBranch(w, r)
		}
	case http.MethodPut:
		h.UpdateBranch(w, r)
	case http.MethodDelete:
		h.DeleteBranch(w, r)
	}

}
func (h *handler) CreateBranch(w http.ResponseWriter, r *http.Request) {
	var branch models.CreateBranch
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error io.ReadAll:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &branch)
	if err != nil {
		fmt.Println("error Unmarshal:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := h.strg.Branch().CreateBranch(branch)
	if err != nil {
		fmt.Println("error CreateBranch:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("created new branch with id: %s", resp)))
}

func (h *handler) GetBranch(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/branch/"):]
	resp, err := h.strg.Branch().GetBranch(models.IdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("error GetBranch:", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error Marshal:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (h *handler) GetAllBranch(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	limitN, _ := strconv.Atoi(limit)

	resp, err := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{
		Limit: limitN,
		Name:  search,
	})

	if err != nil {
		fmt.Println("Error from GetAllBranch:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error marshaling response:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func (h *handler) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var branch models.Branch
	err = json.Unmarshal(body, &branch)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	resp, err := h.strg.Branch().UpdateBranch(branch)
	if err != nil {
		http.Error(w, "Failed to update branch", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Updated branch with ID: %s", resp)
}

func (h *handler) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/branch/"):]
	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{
		Id: id,
	})

	if err != nil {
		http.Error(w, "Failed to delete branch", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "deleted branch with ID: %s", resp)
}
