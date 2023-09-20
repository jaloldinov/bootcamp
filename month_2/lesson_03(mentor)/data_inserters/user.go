package datainserters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

func UserDataInserter() {
	data, err := os.ReadFile("data/users.json")
	if err != nil {
		panic(err)
	}
	var users []models.User
	err = json.Unmarshal(data, &users)
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
	for _, user := range users {
		_, err = db.Exec("INSERT INTO  \"user\" (id, name) VALUES ($1, $2)",
			user.ID, user.Name)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("User data inserted successfully.")
}
