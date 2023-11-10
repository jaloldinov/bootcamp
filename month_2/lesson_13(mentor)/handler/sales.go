package handler

import (
	"fmt"
	"lesson_20/models"
	"sort"
)

func (h *handler) CreateSale(Client_name string, Branch_id, Shop_asissent_id, Cashier_id string, Price float64, Payment_Type models.Payment) {
	id, err := h.strg.Sales().CreateSale(models.CreateSales{
		Client_name:      Client_name,
		Branch_id:        Branch_id,
		Shop_asissent_id: Shop_asissent_id,
		Cashier_id:       Cashier_id,
		Price:            Price,
		Payment_Type:     Payment_Type,
	})
	// ============================================
	// cashier, err := h.strg.Staff().GetStaff(models.IdRequest{Id: Cashier_id})
	// if err != nil {
	// 	fmt.Println("error while getting cashier")
	// 	return
	// }
	// tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: Cashier_id})
	// if err != nil {
	// 	fmt.Println("error while getting tariff")
	// 	return
	// }

	// if Payment_Type == "cash" {
	// 	if tarif.Type == "fixed" {

	// 	}
	// }

	// if tarif.Type==h.cfg.Fixed {
	// amount:=0
	// if paymentType=="cash"{
	// amount=tarif.AmountForCash
	// }else{
	//  amount=tarif.AmountForCard
	// }
	// }else{
	// amount:=0
	// if paymentType=="cash"{
	// amount=price*tarif.AmountForCash/100
	// }else{
	//  amount=price*tarif.AmountForCard/100
	// }
	// }
	// h.strg.Staff().ChangeBalance(amount)
	//  h.strg.StaffTransaction().Create()// type="topup",sourceType="sales" staffId=cashierId, saleId=id

	// if shopAssisstantId!=""{
	// shopAssisstant, err := h.strg.Staff().GetStaff(models.IdRequest{Id: cashierId})
	// if err != nil {
	// 	fmt.Println("error")
	// 	return
	// }
	// }
	// ============================================
	if err != nil {
		fmt.Println("error from CreateSales: ", err.Error())
		return
	}
	fmt.Println("created new staff with id: ", id)
}

func (h *handler) UpdateSale(Id, Client_name string, Branch_id, Shop_asissent_id, Cashier_id string, Price float64, Payment_Type models.Payment, Status models.Status) {
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

func (h *handler) GetTopSaleBranch() {
	resp, err := h.strg.Sales().GetTopSaleBranch()
	if err != nil {
		fmt.Println(err)
		return
	}
	branches, _ := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{})
	branchName := make(map[string]string)

	for _, v := range branches.Branches {
		branchName[v.Id] = v.Name
	}
	for _, structs := range resp {
		fmt.Printf("%s - %s - %f\n", structs.Day, branchName[structs.BranchId], structs.SalesAmount)
	}
}

func (h *handler) GetSaleCountBranch() {
	resp, err := h.strg.Sales().GetSaleCountBranch()
	if err != nil {
		fmt.Println(err)
		return
	}
	branches, _ := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{})
	branchName := make(map[string]string)

	for _, v := range branches.Branches {
		branchName[v.Id] = v.Name
	}
	var sortedSlice []models.SaleCountSumBranch

	for _, structs := range resp {
		sortedSlice = append(sortedSlice, structs)
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].SalesAmount > sortedSlice[j].SalesAmount
	})

	for _, v := range sortedSlice {
		fmt.Printf("%s - %f - %d\n", branchName[v.BranchId], v.SalesAmount, v.Count)

	}

}
