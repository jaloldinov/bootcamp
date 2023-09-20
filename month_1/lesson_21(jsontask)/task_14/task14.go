package task14

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

/*
15. Har kuni o'rtacha user qancha summa product kiritgani va chiqargani bo'yicha jadval:
    branch      o'rtacha+        o'rtacha-
1. Anvar          370 000         435 000
2. Shuhrat        60 000          875 000
...    ...     ...    ...       ...
*/

func Task14() {
	trasnsactions, _ := readTransaction("data/branch_pr_transaction.json")
	users, _ := readUser("data/users.json")
	products, _ := readProducts("data/products.json")

	// branchId va name
	userName := make(map[int]string)
	for _, u := range users {
		userName[u.ID] = u.Name
	}

	productPrice := make(map[int]int)
	for _, p := range products {
		productPrice[p.Id] = p.Price
	}

	type PlusMinus struct {
		plusSumOfPrice  float64
		minusSumOfPrice float64
	}

	// branchid created count
	plusBranchIdTimeCount := make(map[int]map[string]PlusMinus)

	for _, tr := range trasnsactions {
		if tr.Type == "plus" {
			if _, ok := plusBranchIdTimeCount[tr.UserID]; !ok {
				plusBranchIdTimeCount[tr.UserID] = make(map[string]PlusMinus)
			}
			v := plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]]
			v.plusSumOfPrice += float64((tr.Quantity * productPrice[tr.ProductID]))
			plusBranchIdTimeCount[tr.UserID][tr.CreatedAt[:11]] = v

		} else if tr.Type == "minus" {
			if _, ok := plusBranchIdTimeCount[tr.UserID]; !ok {
				plusBranchIdTimeCount[tr.UserID] = make(map[string]PlusMinus)
			}
			v := plusBranchIdTimeCount[tr.UserID][tr.CreatedAt[:11]]
			v.minusSumOfPrice += float64((tr.Quantity * productPrice[tr.ProductID]))
			plusBranchIdTimeCount[tr.UserID][tr.CreatedAt[:11]] = v
		}
	}

	for userID, innerMap := range plusBranchIdTimeCount {
		plusSumOfPrice := 0.0
		minusSumOfPrice := 0.0
		transactionCount := 0.0
		for _, PlusMinus := range innerMap {
			transactionCount++
			plusSumOfPrice += PlusMinus.plusSumOfPrice
			minusSumOfPrice += PlusMinus.minusSumOfPrice
		}
		fmt.Printf("%s, plus: %f minus: %f\n", userName[userID], plusSumOfPrice/transactionCount, minusSumOfPrice/transactionCount)

	}
}

// ================================READERS======================================
func readUser(data string) ([]models.User, error) {
	var products []models.User

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
