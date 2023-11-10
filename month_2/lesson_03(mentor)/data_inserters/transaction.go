package datainserters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

func TransactionDataInserter() {
	data, err := os.ReadFile("data/branch_pr_transaction.json")
	if err != nil {
		panic(err)
	}
	var transactions []models.Transaction
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		panic(err)
	}
	// Connect database
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert data
	for _, tr := range transactions {
		_, err = db.Exec("INSERT INTO branch_transaction (id, branch_id, product_id, user_id, type, quantity) VALUES ($1, $2, $3, $4, $5, $6)",
			tr.ID, tr.BranchID, tr.ProductID, tr.UserID, tr.Type, tr.Quantity)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Transaction data inserted successfully.")
}
