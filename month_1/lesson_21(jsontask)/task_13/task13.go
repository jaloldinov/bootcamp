package task13

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

//  13. Har bir kunda o'rtacha qancha product kiritilgani va chiqarilgani bo'yicha jadval:
//     branch      o'rtacha+   o'rtacha-
//  1. Chilonzor      73         34
//  2. MGorkiy        60         75

func Task13() {
	trasnsactions, _ := readTransaction("data/branch_pr_transaction.json")
	branches, _ := readProducts("data/branches.json")
	// branchId va name
	branchIdName := make(map[int]string)
	for _, b := range branches {
		branchIdName[b.ID] = b.Name
	}

	type PlusMinus struct {
		plusQuantity  int
		minusQuantity int
	}

	// branchid created count
	plusBranchIdTimeCount := make(map[int]map[string]PlusMinus)

	for _, tr := range trasnsactions {
		if tr.Type == "plus" {
			if _, ok := plusBranchIdTimeCount[tr.BranchID]; !ok {
				plusBranchIdTimeCount[tr.BranchID] = make(map[string]PlusMinus)
			}
			v := plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]]
			v.plusQuantity += tr.Quantity
			plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]] = v

		} else {
			if _, ok := plusBranchIdTimeCount[tr.BranchID]; !ok {
				plusBranchIdTimeCount[tr.BranchID] = make(map[string]PlusMinus)
			}
			v := plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]]
			v.minusQuantity += tr.Quantity
			plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]] = v
		}
	}

	for branch_id, innerMap := range plusBranchIdTimeCount {
		plusQuantity := 0
		minusQuantity := 0
		transactionCount := 0
		for _, PlusMinus := range innerMap {
			transactionCount++
			plusQuantity += PlusMinus.plusQuantity
			minusQuantity += PlusMinus.minusQuantity
		}
		fmt.Printf("%s, plus: %d minus: %d\n", branchIdName[branch_id], plusQuantity/transactionCount, minusQuantity/transactionCount)

	}
}

// ================================READERS======================================
func readProducts(data string) ([]models.Branch, error) {
	var products []models.Branch

	p, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(p, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return products, nil
}

func readTransaction(data string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &transactions)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return transactions, nil
}
