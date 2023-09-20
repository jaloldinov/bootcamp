package handler

import (
	"fmt"
	"lesson_20/models"
)

func (h *handler) CreateBranch(name, adress, year string) {
	resp, err := h.strg.Branch().CreateBranch(models.CreateBranch{
		Name:      name,
		Adress:    adress,
		FoundedAt: year,
	})

	if err != nil {
		fmt.Println("error from CreatBranch: ", err.Error())
		return
	}
	fmt.Println("created new branch with id: ", resp)
}

func (h *handler) UpdateBranch(id string, name, adress, founded_at string) {
	resp, err := h.strg.Branch().UpdateBranch(models.Branch{
		Id:        id,
		Name:      name,
		Adress:    adress,
		FoundedAt: founded_at,
	})

	if err != nil {
		fmt.Println("error from UpdateBranch: ", err.Error())
		return
	}
	fmt.Println("Updated branch with id: ", resp)
}

func (h *handler) GetBranch(id string) {
	resp, err := h.strg.Branch().GetBranch(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetBranch: ", err.Error())
		return
	}
	fmt.Println("found branch with id: ", resp)
}

func (h *handler) GetAllBranch(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllBranch: ", err.Error())
		return
	}
	fmt.Println("found all branchs based on filter: ", resp)
}

func (h *handler) DeleteBranch(id string) {
	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteBranch: ", err.Error())
		return
	}
	fmt.Println("deleted branch with id: ", resp)
}
