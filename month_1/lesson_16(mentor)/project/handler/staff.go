package handler

import (
	"fmt"
	"lesson_15/models"
)

func (h *handler) CreateStaff(BranchId, TariffId string, TypeId models.StaffType, Name string, Balance float64) {
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

func (h *handler) UpdateStaff(ID string, BranchId, TariffId string, TypeId models.StaffType, Name string, Balance float64) {
	resp, err := h.strg.Staff().UpdateStaff(models.Staff{
		Id:       ID,
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

func (h *handler) GetStaff(id string) {
	resp, err := h.strg.Staff().GetStaff(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetStaff: ", err.Error())
		return
	}
	fmt.Println("found staff with id: ", resp)
}

func (h *handler) GetAllStaff(page, limit int, staffType models.StaffType, name string, balanceFrom, balanceTo float64) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Staff().GetAllStaff(models.GetAllStaffRequest{
		Page:        page,
		Limit:       limit,
		Type:        staffType,
		Name:        name,
		BalanceFrom: balanceFrom,
		BalanceTo:   balanceTo,
	})

	if err != nil {
		fmt.Println("error from GetAllStaff: ", err.Error())
		return
	}
	fmt.Println("found all Staffes based on filter: ", resp)
}

func (h *handler) DeleteStaff(id string) {
	resp, err := h.strg.Staff().DeleteStaff(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteStaff: ", err.Error())
		return
	}
	fmt.Println("deleted staff with id: ", resp)
}

func (h *handler) ChangeBalance(ID string, Balance float64) {
	resp, err := h.strg.Staff().ChangeBalance(models.ChangeBalance{
		Id:      ID,
		Balance: Balance,
	})

	if err != nil {
		fmt.Println("error from change balance: ", err.Error())
		return
	}
	fmt.Println("Updated balance: ", resp)
}
