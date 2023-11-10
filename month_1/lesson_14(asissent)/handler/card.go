package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"fmt"
)

func (h *handler) CreateCard(name string, quantity, product_id int, created_at string) {
	resp, err := h.strg.Card().CreateCard(models.CreateCard{
		Name:       name,
		Quantity:   quantity,
		Product_id: product_id,
		Created_at: created_at,
	})

	if err != nil {
		fmt.Println("error from CreatCard: ", err.Error())
		return
	}
	fmt.Println("created new card with id: ", resp)
}

func (h *handler) UpdateCard(id int, name string, quantity, product_id int, created_at string) {
	resp, err := h.strg.Card().UpdateCard(models.Card{
		Id:         id,
		Name:       name,
		Quantity:   quantity,
		Product_id: product_id,
		Created_at: created_at,
	})

	if err != nil {
		fmt.Println("error from UpdateCard: ", err.Error())
		return
	}
	fmt.Println("Updated card with id: ", resp)
}

func (h *handler) GetCard(id int) {
	resp, err := h.strg.Card().GetCard(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetCard: ", err.Error())
		return
	}
	fmt.Println("found card with id: ", resp)
}

func (h *handler) GetAllCard(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Card().GetAllCard(models.GetAllCardRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllCard: ", err.Error())
		return
	}
	fmt.Println("found all cards based on filter: ", resp)
}

func (h *handler) DeleteCard(id int) {
	resp, err := h.strg.Card().DeleteCard(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteCard: ", err.Error())
		return
	}
	fmt.Println("deleted card with id: ", resp)
}
