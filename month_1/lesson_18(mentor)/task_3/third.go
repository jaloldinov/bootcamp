package task3

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 3.transactionda bo'lgan top productlar
func TopTransactionProducts() {
	productes, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	var transactionCount = make(map[int]int) //[productId]count
	for _, p := range productes {
		for _, t := range transactions {
			if t.ProductID == p.Id {
				transactionCount[p.Id]++
			}
		}
	}

	var sortedTopProducts []models.ProductTop
	for productId, count := range transactionCount {
		for _, p := range productes {
			if productId == p.Id {
				sortedTopProducts = append(sortedTopProducts, models.ProductTop{
					Id:    productId,
					Name:  p.Name,
					Count: count,
				})
			}
		}
	}

	sort.Slice(sortedTopProducts, func(i, j int) bool {
		return sortedTopProducts[i].Count > sortedTopProducts[j].Count
	})

	for _, v := range sortedTopProducts {
		fmt.Printf("%s: transacted %d times\n", v.Name, v.Count)
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
