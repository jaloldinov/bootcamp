package task12

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 12. har bir user qancha product kiritgani va chiqargani jadvali:
//          kiritgan  chiqargan
// 1 Akrom    12         84
// 2 Suhrob   54         33

func Task12() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	users, _ := readUsers("data/users.json")

	userPlusCount := make(map[int]int)
	userMinusCount := make(map[int]int)

	for _, t := range transactions {
		if t.Type == "plus" {
			userPlusCount[t.UserID] += t.Quantity
		} else {
			userMinusCount[t.UserID] += t.Quantity
		}
	}

	for _, u := range users {
		fmt.Printf("%s - kiritgan: %d chiqargan: %d\n", u.Name, userPlusCount[u.ID], userMinusCount[u.ID])
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
