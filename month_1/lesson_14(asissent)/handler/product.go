package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"fmt"
)

func (h *handler) CreateProduct(name string, card_id, size_id int, created_at string) {
	resp, err := h.strg.Product().CreateProduct(models.CreateProduct{
		Name:       name,
		Card_Id:    card_id,
		Size_Id:    size_id,
		Created_at: created_at,
	})
	if err != nil {
		fmt.Println("error from CreateProduct: ", err.Error())
		return
	}
	fmt.Println("created new Product with id: ", resp)
}

func (h *handler) UpdateProduct(name string, card_id, size_id int, created_at string) {
	resp, err := h.strg.Product().UpdateProduct(models.Product{
		Name:       name,
		Card_Id:    card_id,
		Size_Id:    size_id,
		Created_at: created_at,
	})

	if err != nil {
		fmt.Println("error from UpdateProduct: ", err.Error())
		return
	}
	fmt.Println("Updated Product with id: ", resp)
}

func (h *handler) GetProduct(id int) {
	resp, err := h.strg.Product().GetProduct(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetProduct: ", err.Error())
		return
	}
	fmt.Println("found Product with id: ", resp)
}

func (h *handler) GetAllProduct(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Product().GetAllProduct(models.GetAllProductRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllProduct: ", err.Error())
		return
	}
	fmt.Println("found all Productes based on filter: ", resp)
}

func (h *handler) DeleteProduct(id int) {
	resp, err := h.strg.Product().DeleteProduct(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteProduct: ", err.Error())
		return
	}
	fmt.Println("deleted Product with id: ", resp)
}
