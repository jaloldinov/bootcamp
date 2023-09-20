package task2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 2.transactionlar summasi bo'yicha top branchess

func CalculateSumOfPriceTopBranches() {
	branches, _ := readBranches("data/branches.json")
	productes, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	branchSums := make(map[int]int)
	for _, p := range productes {
		for _, t := range transactions {
			for _, b := range branches {
				if t.BranchID == b.ID && t.ProductID == p.Id {
					for _, transaction := range transactions {
						// sum := t.Quantity * p.Price
						branchSums[transaction.BranchID] += t.Quantity * p.Price
					}
				}
			}
		}
	}

	var sortedBranches []models.BranchProductPrice
	for branchId, sum := range branchSums {
		for _, b := range branches {
			if branchId == b.ID {
				sortedBranches = append(sortedBranches, models.BranchProductPrice{
					BranchID:   branchId,
					BranchName: b.Name,
					Sum:        sum,
				})
			}
		}
	}

	sort.Slice(sortedBranches, func(i, j int) bool {
		return sortedBranches[i].Sum > sortedBranches[j].Sum
	})

	for _, v := range sortedBranches {
		fmt.Printf("Branch: %s: total sum: %d\n", v.BranchName, v.Sum)
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
