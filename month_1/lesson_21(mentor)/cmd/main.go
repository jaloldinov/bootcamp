package main

import (
	"fmt"
	"lesson_20/config"
	"lesson_20/handler"
	"lesson_20/models"
	"lesson_20/storage/memory"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage("data/branch.json", "data/staff.json", "data/sale.json", "data/transaction.json", "data/tariff.json")
	handler := handler.NewHandler(strg, *cfg)

	fmt.Println("Welcome to my Golang Project!")
	fmt.Println("Available methods:")

	for _, m := range cfg.Methods {
		fmt.Println("- ", m)
	}

	fmt.Println("Available objects:")

	for _, o := range cfg.Objects {
		fmt.Println("- ", o)
	}

	for {
		fmt.Print("enter methods and object: ")
		method, object := "", ""
		fmt.Scan(&method, &object)

		switch object {
		// BRANCH
		case "branch":
			switch method {
			case "create":
				fmt.Println("Enter name, adress and founded year: ")
				name, adress, year := "", "", ""
				fmt.Scan(&name, &adress, &year)
				handler.CreateBranch(name, adress, year)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handler.GetBranch(id)
			case "getAll":
				fmt.Print("Enter search text: ")
				var search string
				fmt.Scan(&search)
				handler.GetAllBranch(1, 10, search)
			case "update":
				fmt.Println("Enter ID, name, adress and founded year: ")
				id, name, adress, year := "", "", "", ""
				fmt.Scan(&id, &name, &adress, &year)
				handler.UpdateBranch(id, name, adress, year)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handler.DeleteBranch(id)
			}
		// STAFF
		case "staff":
			switch method {
			case "create":
				fmt.Println("Enter branchId, TariffId, type, Name and balance: ")
				var typId models.StaffType = ""
				branchId, TariffId, name := "", "", ""
				balance := 0.0
				fmt.Scan(&branchId, &TariffId, &typId, &name, &balance)
				handler.CreateStaff(branchId, TariffId, typId, name, balance)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handler.GetStaff(id)
			case "getAll":
				fmt.Println("Enter Type(cashier, shop_assistant), Name, balanceFrom and BalanceTo: ")
				var typId models.StaffType = ""
				name := ""
				balanceFrom, balanceTo := 0.0, 0.0
				fmt.Scan(&typId, &name, &balanceFrom, &balanceTo)
				handler.GetAllStaff(1, 10, typId, name, balanceFrom, balanceTo)
			case "update":
				fmt.Println("Enter ID, BranchId, TariffId, Type(cashier, shop_assistant), Name, Balance:")
				Name, BranchId, TariffId, id := "", "", "", ""
				var TypeId models.StaffType
				Balance := 0.0
				fmt.Scan(&id, &BranchId, &TariffId, &TypeId, &Name, &Balance)
				handler.UpdateStaff(id, BranchId, TariffId, TypeId, Name, Balance)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handler.DeleteStaff(id)
			}
		// SALE
		case "sale":
			switch method {
			case "create":
				fmt.Println("Client_name, Branch_id, Shop_asissent_id, Cashier_id, Price, Payment_Type(card, cash): ")
				Branch_id, Shop_asissent_id, Cashier_id, client_name := "", "", "", ""
				price := 0.0
				var payment int
				fmt.Scan(&client_name, &Branch_id, &Shop_asissent_id, &Cashier_id, &price, &payment)
				handler.CreateSale(client_name, Branch_id, Shop_asissent_id, Cashier_id, price, payment)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handler.GetSale(id)
			case "getAll":
				fmt.Println("Enter Client name: ")
				client_name := ""
				fmt.Scan(&client_name)
				handler.GetAllSale(1, 10, client_name)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handler.DeleteSale(id)
			}
			// TRANSACTION
		case "transaction":
			switch method {
			case "create":
				fmt.Println("type(withdraw,topup), amount, sourceType(sales,bonus), Text, saleId, staffId:")
				amount := 0
				typ, sourceType, text, saleId, staffId := "", "", "", "", ""
				fmt.Scan(&typ, &amount, &sourceType, &text, &saleId, &staffId)
				handler.CreateTransaction(typ, amount, sourceType, text, saleId, staffId)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handler.GetTransaction(id)
			case "getAll":
				// fmt.Print("Enter Text: ")
				// text := ""
				// fmt.Scan(&text)
				handler.GetAllTransaction(1, 10)
			case "update":
				fmt.Println("Enter ID, type(withdraw,topup), amount, sourceType(sales,bonus), Text, saleId, staffId:")
				amount := 0
				Id, typ, sourceType, text, saleId, staffId := "", "", "", "", "", ""
				fmt.Scan(&Id, &typ, &amount, &sourceType, &text, &saleId, &staffId)
				handler.UpdateTransaction(Id, typ, amount, sourceType, text, saleId, staffId)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handler.DeleteTransaction(id)
			case "getTopStaff":
				fmt.Println("Enter staff Type,  FromData and ToDate (2023-08-08)")
				tip, fromDate, toDate := "", "", ""
				fmt.Scan(&tip, &fromDate, &toDate)
				handler.GetTopStaffs(tip, fromDate, toDate)
			}

		// TARIFF
		case "tariff":
			switch method {
			case "create":
				fmt.Println("name, type(fixed, percent), amountCash, amountCard:")
				amountforCard, AmountForCash, typ := 0.0, 0.0, 0
				name := ""
				fmt.Scan(&name, &typ, &amountforCard, &AmountForCash)
				handler.CreateStaffTarif(name, typ, amountforCard, AmountForCash)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(id)
				handler.GetStaffTarif(id)
			case "getAll":
				fmt.Print("Enter Name: ")
				text := ""
				fmt.Scan(&text)
				handler.GetAllStaffTarif(1, 10, "")
			case "update":
				fmt.Println("Enter ID, name, type(fixed, percent), amountCash, amountCard:")
				amountforCard, AmountForCash, typ := 0.0, 0.0, 0
				id, name := "", ""
				fmt.Scan(&id, &name, &typ, &amountforCard, &AmountForCash)
				handler.UpdateStaffTarif(id, name, typ, amountforCard, AmountForCash)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handler.DeleteTransaction(id)
			}
		}
	}

}
