package task7

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish:
//         kun         soni
//     1. 2023-08-04   789
//     2. 2023-08-12   634

func CalculateProductIncome() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	incomeCount := make(map[string]int)
	for _, transaction := range transactions {
		if transaction.Type == "plus" {
			incomeCount[transaction.CreatedAt[:11]] += transaction.Quantity
		}
	}

	var sortedBranches []models.ProductIncome
	for day, count := range incomeCount {
		sortedBranches = append(sortedBranches, models.ProductIncome{
			Day:   day,
			Count: count,
		})
	}

	sort.Slice(sortedBranches, func(i, j int) bool {
		return sortedBranches[i].Count > sortedBranches[j].Count
	})

	for _, v := range sortedBranches {
		fmt.Printf("Day: %s=> income count: %d\n", v.Day, v.Count)
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
