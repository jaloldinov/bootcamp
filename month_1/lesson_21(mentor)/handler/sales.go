package handler

import (
	"fmt"
	"lesson_20/config"
	"lesson_20/models"
	"sort"
)

func (h *handler) CreateSale(Client_name string, Branch_id, Shop_asissent_id, Cashier_id string, Price float64, Payment_Type int) {
	id, err := h.strg.Sales().CreateSale(models.CreateSales{
		Client_name:      Client_name,
		Branch_id:        Branch_id,
		Shop_asissent_id: Shop_asissent_id,
		Cashier_id:       Cashier_id,
		Price:            Price,
		Payment_Type:     Payment_Type,
	})
	if err != nil {
		fmt.Println("error from CreateSale:", err.Error())
		return
	}
	// ============================================
	amount := 0.0
	if Shop_asissent_id != "" {
		shopAssistant, err := h.strg.Staff().GetStaff(models.IdRequest{Id: Shop_asissent_id})
		if err == nil {
			tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: shopAssistant.TariffId})
			if err != nil {
				fmt.Println("error while get staff tarif")
				fmt.Println(err)
				return
			}

			if tarif.Type == config.Fixed {
				if Payment_Type == 2 {
					amount = tarif.AmountForCash
				} else {
					amount = tarif.AmountForCard
				}
			} else if tarif.Type == config.Percent {
				if Payment_Type == 2 {
					amount = Price * tarif.AmountForCash / 100
				} else {
					amount = Price * tarif.AmountForCard / 100
				}
			}

			// shop assitant uchun transaction
			_, err = h.strg.Transaction().CreateTransaction(models.CreateTransaction{
				Sale_id:     id,
				Staff_id:    Shop_asissent_id,
				Type:        "shop_assistant",
				Source_type: "sales",
				Amount:      int(Price),
				Text:        fmt.Sprintf("transcatiion successfull, summa: %.2f", Price),
			})
			if err != nil {
				fmt.Println("Error while create transaction")
				return
			}

			_, err = h.strg.Staff().ChangeBalance(models.ChangeBalance{Id: shopAssistant.Id, Balance: amount})
			if err != nil {
				fmt.Println("Error while change balance")
				return
			}
		} else {
			fmt.Println("shopAssistant not found")
		}

		cashier, err := h.strg.Staff().GetStaff(models.IdRequest{Id: Cashier_id})
		if err == nil {
			tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: cashier.TariffId})
			if err != nil {
				fmt.Println(err)
				return
			}

			if tarif.Type == config.Fixed {
				if Payment_Type == 2 {
					amount = tarif.AmountForCash
				} else {
					amount = tarif.AmountForCard
				}
			} else if tarif.Type == config.Percent {
				if Payment_Type == 2 {
					amount = Price * tarif.AmountForCash / 100
				} else {
					amount = Price * tarif.AmountForCard / 100
				}
			}
		}

		// cashier uchun transaction create qilish
		_, err = h.strg.Transaction().CreateTransaction(models.CreateTransaction{
			Sale_id:     id,
			Staff_id:    Cashier_id,
			Type:        "cashier",
			Source_type: "sales",
			Amount:      int(Price),
			Text:        fmt.Sprintf("transcatiion successfull, summa: %.2f", Price),
		})
		if err != nil {
			fmt.Println("Error while create transaction")
			return
		}
		// cashier change balance
		_, err = h.strg.Staff().ChangeBalance(models.ChangeBalance{Id: cashier.Id, Balance: amount})
		if err != nil {
			fmt.Println("Error while change balance")
			return
		}
	} else {
		fmt.Println("error while get cashier data")
		return
	}

	fmt.Println("created new sale with id:", id)
}

func (h *handler) GetSale(id string) {
	resp, err := h.strg.Sales().GetSale(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetSale:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllSale(page, limit int, clientName string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.Sales().GetAllSale(models.GetAllSalesRequest{
		Page:        page,
		Limit:       limit,
		Client_name: clientName,
	})

	if err != nil {
		fmt.Println("error from GetAllSale:", err.Error())
		return
	}
	fmt.Println(resp)
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

func (h *handler) CancelSale(id string) {
	resp, err := h.strg.Sales().CancelSale(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from CreateSale:", err.Error())
		return
	}

	// shop assistant change balance
	sale, err := h.strg.Sales().GetSale(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error while read data", err)
		return
	}
	shopAsistant, err := h.strg.Staff().GetStaff(models.IdRequest{Id: sale.Shop_asissent_id})
	if err == nil {
		amount := 0.0
		tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: shopAsistant.TariffId})
		if err != nil {
			fmt.Println("error while get staff tarif")
			fmt.Println(err)
			return
		}

		if tarif.Type == config.Fixed {
			if sale.Payment_Type == 2 {
				amount = tarif.AmountForCash
			} else {
				amount = tarif.AmountForCard
			}
		} else if tarif.Type == config.Percent {
			if sale.Payment_Type == 2 {
				amount = sale.Price * tarif.AmountForCash / 100
			} else {
				amount = sale.Price * tarif.AmountForCard / 100
			}
		}

		// shop assitant uchun transaction
		_, err = h.strg.Transaction().CreateTransaction(models.CreateTransaction{
			Sale_id:     resp,
			Staff_id:    shopAsistant.Id,
			Type:        "shop_assistant",
			Source_type: "sales",
			Amount:      int(sale.Price),
			Text:        fmt.Sprintf("transcatiion cancelled, summa: %.2f", sale.Price),
		})
		if err != nil {
			fmt.Println("Error while create transaction")
			return
		}

		_, err = h.strg.Staff().ChangeBalance(models.ChangeBalance{Id: shopAsistant.Id, Balance: -amount})
		if err != nil {
			fmt.Println("Error while change balance")
			return
		}
	} else {
		fmt.Println("shopAssistant not found")
	}

	cashier, err := h.strg.Staff().GetStaff(models.IdRequest{Id: sale.Cashier_id})
	if err == nil {
		amount := 0.0
		tarif, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: cashier.TariffId})
		if err != nil {
			fmt.Println(err)
			return
		}

		if tarif.Type == config.Fixed {
			if sale.Payment_Type == 2 {
				amount = tarif.AmountForCash
			} else {
				amount = tarif.AmountForCard
			}
		} else if tarif.Type == config.Percent {
			if sale.Payment_Type == 2 {
				amount = sale.Price * tarif.AmountForCash / 100
			} else {
				amount = sale.Price * tarif.AmountForCard / 100
			}
		}

		// cashier uchun transaction create qilish
		_, err = h.strg.Transaction().CreateTransaction(models.CreateTransaction{
			Sale_id:     resp,
			Staff_id:    cashier.Id,
			Type:        "cashier",
			Source_type: "sales",
			Amount:      int(sale.Price),
			Text:        fmt.Sprintf("transcatiion cancelled, summa: %.2f", sale.Price),
		})
		if err != nil {
			fmt.Println("Error while create transaction")
			return
		}
		// cashier change balance
		_, err = h.strg.Staff().ChangeBalance(models.ChangeBalance{Id: cashier.Id, Balance: -amount})
		if err != nil {
			fmt.Println("Error while change balance")
			return
		}
	} else {
		fmt.Println("error while get cashier data")
		return
	}

	fmt.Println(resp)
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
