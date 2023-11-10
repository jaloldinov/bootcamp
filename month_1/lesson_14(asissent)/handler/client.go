package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"fmt"
)

func (h *handler) CreateClient(name string, card_id int, created_at string) {
	resp, err := h.strg.Client().CreateClient(models.CreateClient{
		Name:       name,
		Card_Id:    card_id,
		Created_at: created_at,
	})
	if err != nil {
		fmt.Println("error from CreateClient: ", err.Error())
		return
	}
	fmt.Println("created new Client with id: ", resp)
}

func (h *handler) UpdateClient(name string, card_id int, created_at string) {
	resp, err := h.strg.Client().UpdateClient(models.Client{
		Name:       name,
		Card_Id:    card_id,
		Created_at: created_at,
	})

	if err != nil {
		fmt.Println("error from UpdateClient: ", err.Error())
		return
	}
	fmt.Println("Updated Client with id: ", resp)
}

func (h *handler) GetClient(id int) {
	resp, err := h.strg.Client().GetClient(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetClient: ", err.Error())
		return
	}
	fmt.Println("found Client with id: ", resp)
}

func (h *handler) GetAllClient(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Client().GetAllClient(models.GetAllClientRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllClient: ", err.Error())
		return
	}
	fmt.Println("found all Clientes based on filter: ", resp)
}

func (h *handler) DeleteClient(id int) {
	resp, err := h.strg.Client().DeleteClient(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteClient: ", err.Error())
		return
	}
	fmt.Println("deleted Client with id: ", resp)
}
