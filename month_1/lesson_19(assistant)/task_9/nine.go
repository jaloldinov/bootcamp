package task9

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 9. Filialda qancha summalik product borligi jadvali:
// 1. Branch1        853 000
// 2. Branch2      1 982 000

func CalculateProductSum() {
	branch_products, _ := readBranches("data/branch_products.json")
	productes, _ := readProducts("data/products.json")
	branches, _ := readBranch("data/branches.json")

	prodIdPrice := make(map[int]int)
	branchIdName := make(map[int]string)

	for _, p := range productes {
		prodIdPrice[p.Id] = p.Price
	}
	for _, b := range branches {
		branchIdName[b.ID] = b.Name
	}

	branchSum := make(map[string]int)
	for _, v := range branch_products {
		branchSum[branchIdName[v.BranchId]] += (v.Quantity * prodIdPrice[v.ProductId])
	}
	for name, sum := range branchSum {
		fmt.Printf("%s - %d\n", name, sum)
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

func readBranches(data string) ([]models.Product7task, error) {
	var branches []models.Product7task

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
func readBranch(data string) ([]models.Branch, error) {
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
