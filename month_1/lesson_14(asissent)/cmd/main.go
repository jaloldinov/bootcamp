package main

import (
	"backend_bootcamp_17_07_2023/lesson_14/config"
	"backend_bootcamp_17_07_2023/lesson_14/handler"
	"backend_bootcamp_17_07_2023/lesson_14/storage/memory"
	"fmt"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage()
	h := handler.NewHandler(strg, *cfg)

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

		if object == "branch" && method == "getAll" {
			h.GetAllBranch(1, 10, "")
		}
	}
	
	type branchRepo struct {
	fileName string
}

func NewBranchRepo(fileName string) *branchRepo {
	return &branchRepo{fileName: fileName}
}
	//================== BRANCH ==================
	// fmt.Println("Enter name and adress: ")
	// name, adress := "", ""
	// fmt.Scan(&name, &adress)
	// h.CreateBranch(name, adress)

	// fmt.Print("To get Branch enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetBranch(id)

	//================== STAFF ==================
	// fmt.Println("Enter BranchId, TariffId, TypeId, Name and Balance: ")
	// branchId, tariffId, typeId := 0, 0, 0
	// name := ""
	// balance := 0.0
	// fmt.Scan(&branchId, &tariffId, &typeId, &name, &balance)
	// h.CreateStaff(branchId, tariffId, typeId, name, balance)
	// fmt.Print("To get Staff enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetStaff(id)

	//================== PRODUCT ==================
	// fmt.Println("Enter name, card_id, size_id, created_at: ")
	// card_id, size_id := 0, 0
	// name, created_at := "", ""
	// fmt.Scan(&name, &card_id, &size_id, &created_at)
	// h.CreateProduct(name, card_id, size_id, created_at)
	// fmt.Print("To get Product enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetProduct(id)

	//================== CLIENT ==================
	// fmt.Println("Enter Name, Card_id and Created_at: ")
	// card_id := 0
	// name, created_at := "", ""
	// fmt.Scan(&name, &card_id, &created_at)
	// h.CreateClient(name, card_id, created_at)
	// fmt.Print("To get Staff enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetClient(id)

	// ================== CARD ==================
	// fmt.Println("Enter Name, Quantity, Product_Id and Created_at: ")
	// quantity, product_id := 122, 2
	// name, created_at := "apple", "04-07-2023"
	// // fmt.Scan(&name, &quantity, &product_id, created_at)
	// h.CreateCard(name, quantity, product_id, created_at)
	// fmt.Print("To get Card enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetCard(id)

	// ================== SIZE ==================
	// fmt.Println("Enter Name, Price and Created_at: ")
	// price := 1002.
	// name, created_at := "1litr", "04-07-2023"
	// h.CreateSize(name, price, created_at)
	// fmt.Print("To get Size enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetSize(id)

	// ================== SALES ==================
	// fmt.Println("Enter Name, Price, payment_type, status, client_id, branch_id, shop_ass_id, cashier_id and Created_at: ")
	// name := "nimadir"
	// Price := 121.2
	// Payment_Type := 1
	// Status := 1
	// Client_id := 1
	// Branch_id := 1
	// Shop_asissent_id := 1
	// Cashier_id := 1
	// Created_at := "2002-20-23"

	// h.CreateSales(name, Price, Payment_Type, Status, Client_id, Branch_id, Shop_asissent_id, Cashier_id, Created_at)
	// fmt.Print("To get sales enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetSales(id)

	// Name             string
	// Price            float64
	// Payment_Type     int
	// Status           int
	// Client_id        int
	// Branch_id        int
	// Shop_asissent_id int
	// Cashier_id       int
	// Created_at       string

	// ================== TRANSACTION ==================
	// fmt.Println("Enter Name, Price, payment_type, status, client_id, branch_id, shop_ass_id, cashier_id and Created_at: ")
	// // Type, Amount int, source_type, text string, sale_id float64, staff_id int, created_at string

	// h.CreateSales(name, Price, Payment_Type, Status, Client_id, Branch_id, Shop_asissent_id, Cashier_id, Created_at)
	// fmt.Print("To get sales enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetSales(id)
}
