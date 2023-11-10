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
	products, _ := readProduct("data/products.json")
	branches, _ := readBranches("data/branches.json")

	categoryMap := make(map[int]models.Category)
	productCategoryMap := make(map[int]int)
	branchMap := make(map[int]string)

	for _, b := range branches {
		branchMap[b.ID] = b.Name
	}

	for _, category := range categories {
		categoryMap[category.Id] = category
	}

	for _, product := range products {
		productCategoryMap[product.Id] = product.CategoryID
	}

	countMap := make(map[int]map[int]int)

	for _, tr := range transactions {
		if countMap[tr.BranchID] == nil {
			countMap[tr.BranchID] = make(map[int]int)
		}
		countMap[tr.BranchID][productCategoryMap[tr.ProductID]]++
	}

	for branchID, catCount := range countMap {
		fmt.Println(branchMap[branchID])
		for catID, count := range catCount {
			category := categoryMap[catID]
			fmt.Printf("%s: %d\n", category.Name, count)
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

func readTCategory(data string) ([]models.Category, error) {
	var categories []models.Category

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

func readProduct(data string) ([]models.Products, error) {
	var productes []models.Products

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &productes)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return productes, nil
}
