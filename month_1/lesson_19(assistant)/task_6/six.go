package task6

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
//                     Transactions            Summ
//                     plus   minus        plus     minus
//     1. Branch1      53      20          853 000  278 000
//     2. Branch2      38      185         492 000  1 982 000
// { BranchName: transactionPlus: 23, minus: 12:
//     "id": 1,
//     "branch_id": 4,
//     "product_id": 3,
//     "type": "plus",
//     "quantity": 74,
//     "created_at": "2023-08-09 20:05:37"
//   },

func PlusMinus() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	productes, _ := readProduct("data/products.json")
	branches, _ := readBranches("data/branches.json")

	plusBranchIdTransCount := make(map[int]int)
	minusBranchIdTransCount := make(map[int]int)
	plusBranchIdSum := make(map[int]int)
	minusBranchIdSum := make(map[int]int)

	BranchMap := make(map[int]string)
	ProductMap := make(map[int]int)

	for _, br := range branches {
		BranchMap[br.ID] = br.Name
	}
	for _, pr := range productes {
		ProductMap[pr.Id] = pr.Price
	}

	for _, tr := range transactions {
		if tr.Type == "plus" {
			plusBranchIdTransCount[tr.BranchID]++
			plusBranchIdSum[tr.BranchID] += tr.Quantity * ProductMap[tr.ProductID]
		} else {
			minusBranchIdTransCount[tr.BranchID]++
			minusBranchIdSum[tr.BranchID] += tr.Quantity * ProductMap[tr.ProductID]
		}
	}

	for BranchID, BranchName := range BranchMap {
		fmt.Printf("%s TranPlus: %d, Tranminus: %d, SumPlus: %d, SumMinus: %d\n", BranchName,
			plusBranchIdTransCount[BranchID],
			minusBranchIdTransCount[BranchID],
			plusBranchIdSum[BranchID],
			minusBranchIdSum[BranchID],
		)
	}

}

// // ================================READERS======================================
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
