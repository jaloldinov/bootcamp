package handler

import (
	"fmt"
	"lesson_15/models"
)

func (h *handler) CreateSale(Client_name string, Branch_id, Shop_asissent_id, Cashier_id int, Price float64, Payment_Type models.Payment, Status models.Status) {
	resp, err := h.strg.Sales().CreateSale(models.CreateSales{
		Client_name:      Client_name,
		Branch_id:        Branch_id,
		Shop_asissent_id: Shop_asissent_id,
		Cashier_id:       Cashier_id,
		Price:            Price,
		Payment_Type:     Payment_Type,
		Status:           Status,
	})
	if err != nil {
		fmt.Println("error from CreateSales: ", err.Error())
		return
	}
	fmt.Println("created new staff with id: ", resp)
}

func (h *handler) UpdateSale(Id, Client_name string, Branch_id, Shop_asissent_id, Cashier_id int, Price float64, Payment_Type models.Payment, Status models.Status) {
	resp, err := h.strg.Sales().UpdateSale(models.Sales{
		Id:               Id,
		Client_name:      Client_name,
		Branch_id:        Branch_id,
		Shop_asissent_id: Shop_asissent_id,
		Cashier_id:       Cashier_id,
		Price:            Price,
		Payment_Type:     Payment_Type,
		Status:           Status,
	})

	if err != nil {
		fmt.Println("error from UpdateSales: ", err.Error())
		return
	}
	fmt.Println("Updated staff with id: ", resp)
}

func (h *handler) GetSale(id string) {
	resp, err := h.strg.Sales().GetSale(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetSales: ", err.Error())
		return
	}
	fmt.Println("found staff with id: ", resp)
}

func (h *handler) GetAllSale(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.Sales().GetAllSale(models.GetAllSalesRequest{
		Page:        page,
		Limit:       limit,
		Client_name: search,
	})

	if err != nil {
		fmt.Println("error from GetAllSales: ", err.Error())
		return
	}
	fmt.Println("found all Saleses based on filter: ", resp)
}

func (h *handler) DeleteSale(id string) {
	resp, err := h.strg.Sales().DeleteSale(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteSales: ", err.Error())
		return
	}
	fmt.Println("deleted staff with id: ", resp)
}
