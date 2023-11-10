package datainserters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

func BranchProductDataInserter() {
	data, err := os.ReadFile("data/branch_products.json")
	if err != nil {
		panic(err)
	}
	var branches []models.BranchProduct
	err = json.Unmarshal(data, &branches)
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
	for _, branch := range branches {
		_, err = db.Exec("INSERT INTO branch_products (branch_id, product_id, quantity) VALUES ($1, $2, $3)",
			branch.BranchId, branch.ProductId, branch.Quantity)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Branch Product data inserted successfully.")
}
