package handler

import (
	"fmt"
	"lesson_15/models"
)

func (h *handler) CreateTransaction(typ string, amount float64, sourceType, text, saleId string, staffId string) {
	resp, err := h.strg.Transaction().CreateTransaction(models.CreateTransaction{
		Type:        typ,
		Amount:      amount,
		Source_type: sourceType,
		Text:        text,
		Sale_id:     saleId,
		Staff_id:    staffId,
	})
	if err != nil {
		fmt.Println("error from CreateTransaction: ", err.Error())
		return
	}
	fmt.Println("created new transaction with id: ", resp)
}

func (h *handler) UpdateTransaction(id, typ string, amount float64, sourceType, text, saleId string, staffId string) {
	resp, err := h.strg.Transaction().UpdateTransaction(models.Transaction{
		Id:          id,
		Type:        typ,
		Amount:      amount,
		Source_type: sourceType,
		Text:        text,
		Sale_id:     saleId,
		Staff_id:    staffId,
	})

	if err != nil {
		fmt.Println("error from UpdateTransaction: ", err.Error())
		return
	}
	fmt.Println("Updated transaction with id: ", resp)
}

func (h *handler) GetTransaction(id string) {
	resp, err := h.strg.Transaction().GetTransaction(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetTransaction: ", err.Error())
		return
	}
	fmt.Println("found transaction with id: ", resp)
}

func (h *handler) GetAllTransaction(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.Transaction().GetAllTransaction(models.GetAllTransactionRequest{
		Page:  page,
		Limit: limit,
		Text:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllTransaction: ", err.Error())
		return
	}
	fmt.Println("found all Transactiones based on filter: ", resp)
}

func (h *handler) DeleteTransaction(id string) {
	resp, err := h.strg.Transaction().DeleteTransaction(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteTransaction: ", err.Error())
		return
	}
	fmt.Println("deleted transaction with id: ", resp)
}
