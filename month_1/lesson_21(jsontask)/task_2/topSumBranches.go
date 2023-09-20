package task2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 2.transactionlar summasi bo'yicha top branchess

func CalculateSumOfPriceTopBranches() {
	branches, _ := readBranches("data/branches.json")
	productes, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	branchName := make(map[int]string)
	productPrice := make(map[int]int)
	branchSum := make(map[string]int)

	for _, p := range productes {
		productPrice[p.Id] = p.Price
	}
	for _, b := range branches {
		branchName[b.ID] = b.Name
	}
	for _, t := range transactions {
		branchSum[branchName[t.BranchID]] += (productPrice[t.ProductID] * t.Quantity)
	}

	for n, s := range branchSum {
		fmt.Printf("%s - %d\n", n, s)
	}

}

// ================================READERS======================================

func readProducts(data string) ([]models.Products, error) {
	var products []models.Products

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

func readBranches(data string) ([]models.Branch, error) {
	var branches []models.Branch

	branch, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read branch: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(branch, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal branch: %+v", err)
		return nil, err
	}
	return branches, nil
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
