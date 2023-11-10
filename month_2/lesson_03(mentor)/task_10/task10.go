package task10

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 10. har bir user transaction qilgan summasi jadvali:

// 1  Akrom   1 349 000
// 3  Ilhom   2 974 000
func Task10() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		SELECT u.name, SUM(p.price*bt.quantity) 
		FROM branch_transaction bt
		JOIN product p ON p.id = bt.product_id
		JOIN branch b ON b.id = bt.branch_id 
		JOIN "user" u ON u.id = bt.user_id
		GROUP BY u.name
	`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task2
		err := rows.Scan(&t.BranchName, &t.Sum)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d\n", t.BranchName, t.Sum)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
