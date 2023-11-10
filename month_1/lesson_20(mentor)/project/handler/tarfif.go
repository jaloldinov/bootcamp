package handler

import (
	"fmt"
	"lesson_20/models"
)

func (h *handler) CreateStaffTarif(name, typ string, amountCash, amountCard int) {
	resp, err := h.strg.StaffTarif().CreateStaffTarif(models.CreateStaffTarif{
		Name:          name,
		Type:          typ,
		AmountForCash: amountCash,
		AmountForCard: amountCard,
	})

	if err != nil {
		fmt.Println("error from CreateStaffTarif:", err.Error())
		return
	}
	fmt.Println("created new tarif with id:", resp)
}

func (h *handler) UpdateStaffTarif(id, name, typ string, amountCash, amountCard int) {
	resp, err := h.strg.StaffTarif().UpdateStaffTarif(models.StaffTarif{
		Id:            id,
		Name:          name,
		Type:          typ,
		AmountForCash: amountCash,
		AmountForCard: amountCard,
	})

	if err != nil {
		fmt.Println("error from UpdateStaffTarif:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetStaffTarif(id string) {
	resp, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetStaffTarif:", err.Error())
		return
	}

	fmt.Println(resp)
}

func (h *handler) GetAllStaffTarif(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.StaffTarif().GetAllStaffTarif(models.GetAllStaffTarifRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) DeleteStaffTarif(id string) {
	resp, err := h.strg.StaffTarif().DeleteStaffTarif(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteStaffTarif:", err.Error())
		return
	}
	fmt.Println(resp)
}
