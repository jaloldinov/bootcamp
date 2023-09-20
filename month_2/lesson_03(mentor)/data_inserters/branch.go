package datainserters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

func BranchDataInserter() {
	data, err := os.ReadFile("data/branches.json")
	if err != nil {
		panic(err)
	}
	var branches []models.Branch
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
		_, err = db.Exec("INSERT INTO branch (id, name, adress) VALUES ($1, $2, $3)",
			branch.ID, branch.Name, branch.Adress)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Branch data inserted successfully.")
}
