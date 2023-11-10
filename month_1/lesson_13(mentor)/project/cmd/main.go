package main

import (
	"backend_bootcamp_17_07_2023/lesson_8/project/handler"
	"backend_bootcamp_17_07_2023/lesson_8/project/storage/memory"
	"fmt"
)

func main() {
	strg := memory.NewStorage()
	h := handler.NewHandler(strg)

	// fmt.Println("Enter name and adress: ")
	// name, adress := "", ""
	// fmt.Scan(&name, &adress)
	// h.CreateBranch(name, adress)

	// fmt.Print("To get Branch enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetBranch(id)

	// STAFF
	fmt.Println("Enter BranchId, TariffId, TypeId, Name and Balance: ")
	branchId, tariffId, typeId := 0, 0, 0
	name := ""
	balance := 0.0
	fmt.Scan(&branchId, &tariffId, &typeId, &name, &balance)
	h.CreateStaff(branchId, tariffId, typeId, name, balance)

	fmt.Print("To get Staff enter ID: ")
	id := 0
	fmt.Scan(&id)
	h.GetStaff(id)
}
