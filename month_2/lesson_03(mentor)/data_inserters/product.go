package datainserters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

func ProductDataInserter() {
	data, err := os.ReadFile("data/products.json")
	if err != nil {
		panic(err)
	}
	var products []models.Products
	err = json.Unmarshal(data, &products)
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
	for _, product := range products {
		_, err = db.Exec("INSERT INTO  product (id, name, price, category_id) VALUES ($1, $2, $3, $4)",
			product.Id, product.Name, product.Price, product.CategoryID)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Product data inserted successfully.")
}
