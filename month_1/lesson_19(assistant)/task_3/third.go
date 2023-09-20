package task3

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 3.transactionda bo'lgan top productlar
func TopTransactionProducts() {
	productes, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	transactionCount := make(map[int]int) //[productId]count
	prodNameCount := make(map[int]string)
	for _, t := range transactions {
		transactionCount[t.ProductID]++
	}
	for _, p := range productes {
		prodNameCount[p.Id] = p.Name
	}

	for id, t := range transactionCount {
		fmt.Printf("%s - %d\n", prodNameCount[id], t)
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
