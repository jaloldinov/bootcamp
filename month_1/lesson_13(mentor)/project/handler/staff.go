package handler

import (
	"backend_bootcamp_17_07_2023/lesson_8/project/models"
	"fmt"
)

func (h *handler) CreateStaff(BranchId, TariffId, TypeId int, Name string, Balance float64) {
	resp, err := h.strg.Staff().CreateStaff(models.CreateStaff{
		BranchId: BranchId,
		TariffId: TariffId,
		TypeId:   TypeId,
		Name:     Name,
		Balance:  Balance,
	})
	if err != nil {
		fmt.Println("error from CreateStaff: ", err.Error())
		return
	}
	fmt.Println("created new staff with id: ", resp)
}

func (h *handler) UpdateStaff(BranchId, TariffId, TypeId int, Name string, Balance float64) {
	resp, err := h.strg.Staff().UpdateStaff(models.Staff{
		BranchId: BranchId,
		TariffId: TariffId,
		TypeId:   TypeId,
		Name:     Name,
		Balance:  Balance,
	})

	if err != nil {
		fmt.Println("error from UpdateStaff: ", err.Error())
		return
	}
	fmt.Println("Updated staff with id: ", resp)
}

func (h *handler) GetStaff(id int) {
	resp, err := h.strg.Staff().GetStaff(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetStaff: ", err.Error())
		return
	}
	fmt.Println("found staff with id: ", resp)
}

func (h *handler) GetAllStaff(page, limit int, search string) {
	resp, err := h.strg.Staff().GetAllStaff(models.GetAllStaffRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllStaff: ", err.Error())
		return
	}
	fmt.Println("found all Staffes based on filter: ", resp)
}

func (h *handler) DeleteStaff(id int) {
	resp, err := h.strg.Staff().DeleteStaff(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteStaff: ", err.Error())
		return
	}
	fmt.Println("deleted staff with id: ", resp)
}
