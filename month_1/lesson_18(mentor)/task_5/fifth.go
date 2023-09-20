package task5

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 5.har bir branchda har bir categorydan qancha transaction bo'lgani

func TopBranchTransactionCategory() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	categories, _ := readTCategory("data/categories.json")
	branches, _ := readBranches("data/branches.json")

	countMap := make(map[string]int)

	// Count the transactions
	for _, transaction := range transactions {
		key := fmt.Sprintf("Branch %d, Category %d", transaction.BranchID, transaction.ProductID)
		countMap[key]++
	}

	// Print the branches with non-zero transaction counts for each category
	for _, branch := range branches {
		fmt.Printf("Branch: %s\n", branch.Name)
		hasTransactions := false
		for _, category := range categories {
			key := fmt.Sprintf("Branch %d, Category %d", branch.ID, category.Id)
			count := countMap[key]
			if count > 0 {
				hasTransactions = true
				fmt.Printf("%s => transacted %d times\n", category.Name, count)
			}
		}
		if !hasTransactions {
			fmt.Println("No transactions found for any category.")
		}
	}
}

// ================================READERS======================================
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

func readTCategory(data string) ([]models.ProductTop, error) {
	var categories []models.ProductTop

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &categories)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return categories, nil
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
