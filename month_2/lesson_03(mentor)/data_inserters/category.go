package datainserters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

func CategoryDataInserter() {
	data, err := os.ReadFile("data/categories.json")
	if err != nil {
		panic(err)
	}
	var categories []models.Category
	err = json.Unmarshal(data, &categories)
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
	for _, category := range categories {
		_, err = db.Exec("INSERT INTO  category (id, name) VALUES ($1, $2)",
			category.ID, category.Name)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Category data inserted successfully.")
}
