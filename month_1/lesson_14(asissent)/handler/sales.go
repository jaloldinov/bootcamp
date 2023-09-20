package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"fmt"
)

func (h *handler) CreateSales(name string, price float64, payment_type, status, client_id, branch_id, shop_ass_id, cashier_id int, created_at string) {
	resp, err := h.strg.Sales().CreateSales(models.CreateSales{
		Name:             name,
		Price:            price,
		Payment_Type:     payment_type,
		Status:           status,
		Client_id:        client_id,
		Branch_id:        branch_id,
		Shop_asissent_id: shop_ass_id,
		Cashier_id:       cashier_id,
		Created_at:       created_at,
	})
	if err != nil {
		fmt.Println("error from CreateSales: ", err.Error())
		return
	}
	fmt.Println("created new Sales with id: ", resp)
}

func (h *handler) UpdateSales(id int, name string, price float64, payment_type, status, client_id, branch_id, shop_ass_id, cashier_id int, created_at string) {
	resp, err := h.strg.Sales().UpdateSales(models.Sales{
		Id:               id,
		Name:             name,
		Price:            price,
		Payment_Type:     payment_type,
		Status:           status,
		Client_id:        client_id,
		Branch_id:        branch_id,
		Shop_asissent_id: shop_ass_id,
		Cashier_id:       cashier_id,
		Created_at:       created_at,
	})

	if err != nil {
		fmt.Println("error from UpdateSales: ", err.Error())
		return
	}
	fmt.Println("Updated Sales with id: ", resp)
}

func (h *handler) GetSales(id int) {
	resp, err := h.strg.Sales().GetSales(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetSales: ", err.Error())
		return
	}
	fmt.Println("found Sales with id: ", resp)
}

func (h *handler) GetAllSales(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Sales().GetAllSales(models.GetAllSalesRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllSales: ", err.Error())
		return
	}
	fmt.Println("found all Saleses based on filter: ", resp)
}

func (h *handler) DeleteSales(id int) {
	resp, err := h.strg.Sales().DeleteSales(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteSales: ", err.Error())
		return
	}
	fmt.Println("deleted Sales with id: ", resp)
}
