package task4

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 4.transactionda bo'lgan top categorylar
func TopTransactionCategory() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	categories, _ := readTCategory("data/categories.json")

	var transactionCount = make(map[int]int) //[productId]count

	for _, t := range transactions {
		transactionCount[t.ProductID]++
	}

	var sortedTopProducts []models.ProductTop
	for _, c := range categories {
		for k, t := range transactionCount {
			if c.Id == k {
				sortedTopProducts = append(sortedTopProducts, models.ProductTop{
					Id:    k,
					Name:  c.Name,
					Count: t,
				})
			}
		}
	}

	sort.Slice(sortedTopProducts, func(i, j int) bool {
		return sortedTopProducts[i].Count > sortedTopProducts[j].Count
	})

	for _, v := range sortedTopProducts {
		fmt.Printf("%s: %d times transacted\n", v.Name, v.Count)
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
