package task10

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 10. har bir user transaction qilgan summasi jadvali:

// 1  Akrom   1 349 000
// 3  Ilhom   2 974 000
func Task10() {
	trasnsactions, _ := readTransaction("data/branch_pr_transaction.json")
	products, _ := readProducts("data/products.json")
	users, _ := readUsers("data/users.json")
	/*
		Avval qaysidir objectni ma'lumotlarini saqlaydigan maplarni
		birinchi qivolish kerak. Keyin sanaydigan yoki summani hisoblaydigan mapni
		 qilsangiz,bu mapni qilishda narigi maplardan foydalansa bo'ladi ya'ni hisoblash osonroq bo'ladi
		Masalan bunda products bilan usersni birinchi qivolish kerak
		Keyin transaction bo'ylab aylanib keyda userId valueda summni saqlaydigan map qilsa bo'ladi
		Buni qilayotganda sizda uje productsMap bor transaction ichida quantity bor
		bir yo'la summani hisoblab qo'shaverasiz
	*/

	// userid va userName
	userIdName := make(map[int]string)
	// productID va price
	prodIdPrice := make(map[int]int)

	for _, u := range users {
		userIdName[u.ID] = u.Name
	}
	for _, p := range products {
		prodIdPrice[p.Id] = p.Price
	}

	nameSum := make(map[string]int)
	for _, tr := range trasnsactions {
		nameSum[userIdName[tr.UserID]] += tr.Quantity * prodIdPrice[tr.ProductID]
	}
	for name, sum := range nameSum {
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
