package task8

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 8. Product qancha kiritilgan va chiqarilganligi jadvali:
//     Name    Kiritilgan  Chiqarilgan
//     Olma     345            847
//     Cola     374            219
//     ....     ...       ...   ....
// products, transaction

func Task8() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	products, _ := readProducts("data/products.json")

	inputCount := make(map[int]int)
	outputCount := make(map[int]int)

	for _, t := range transactions {
		if t.Type == "plus" {
			inputCount[t.ProductID] += t.Quantity
		} else {
			outputCount[t.ProductID] += t.Quantity
		}
	}

	for _, p := range products {
		if inputCount[p.Id] > 0 || outputCount[p.Id] > 0 {
			fmt.Printf("%s - kiritilgan: %d chiqarilgan: %d\n", p.Name, inputCount[p.Id], outputCount[p.Id])
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
