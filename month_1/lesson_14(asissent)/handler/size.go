package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"fmt"
)

func (h *handler) CreateSize(name string, price float64, created_at string) {
	resp, err := h.strg.Size().CreateSize(models.CreateSize{
		Name:       name,
		Price:      price,
		Created_at: created_at,
	})

	if err != nil {
		fmt.Println("error from CreatSize: ", err.Error())
		return
	}
	fmt.Println("created new size with id: ", resp)
}

func (h *handler) UpdateSize(id int, name string, price float64, created_at string) {
	resp, err := h.strg.Size().UpdateSize(models.Size{
		Id:         id,
		Name:       name,
		Price:      price,
		Created_at: created_at,
	})

	if err != nil {
		fmt.Println("error from UpdateSize: ", err.Error())
		return
	}
	fmt.Println("Updated size with id: ", resp)
}

func (h *handler) GetSize(id int) {
	resp, err := h.strg.Size().GetSize(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetSize: ", err.Error())
		return
	}
	fmt.Println("found size with id: ", resp)
}

func (h *handler) GetAllSize(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Size().GetAllSize(models.GetAllSizeRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllSize: ", err.Error())
		return
	}
	fmt.Println("found all sizes based on filter: ", resp)
}

func (h *handler) DeleteSize(id int) {
	resp, err := h.strg.Size().DeleteSize(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteSize: ", err.Error())
		return
	}
	fmt.Println("deleted size with id: ", resp)
}
