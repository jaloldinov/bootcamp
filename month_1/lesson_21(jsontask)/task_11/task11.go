package task11

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 11. har bir user kun bo'yicha nechta va necha sumlik transaction qilgani jadvali:

// 1 Akrom 2023-01-01  12  382 000
// 2 Suhrob 2023-03-05  32  745 000
func Task11() {
	trasnsactions, _ := readTransaction("data/branch_pr_transaction.json")
	products, _ := readProducts("data/products.json")
	users, _ := readUsers("data/users.json")

	// userid va userName
	userIdName := make(map[int]string)
	// productID va price
	prodIdPrice := make(map[int]int)
	// userID: time va count
	userTimeCount := make(map[int]map[string]int)

	for _, u := range users {
		userIdName[u.ID] = u.Name
	}
	for _, p := range products {
		prodIdPrice[p.Id] = p.Price
	}
	totalSumMap := make(map[string]int)

	for _, tr := range trasnsactions {
		if _, ok := userTimeCount[tr.UserID]; !ok {
			userTimeCount[tr.UserID] = make(map[string]int)
		}
		userTimeCount[tr.UserID][tr.CreatedAt[:11]]++
		totalSumMap[tr.CreatedAt[:11]] = prodIdPrice[tr.ProductID]
	}

	for userId, innerMap := range userTimeCount {
		for time, count := range innerMap {
			fmt.Printf("%s - %s - %d - %d\n", userIdName[userId], time, count, (totalSumMap[time] * count))
		}
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

func readUsers(data string) ([]models.User, error) {
	var users []models.User
	p, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(p, &users)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return users, nil
}
