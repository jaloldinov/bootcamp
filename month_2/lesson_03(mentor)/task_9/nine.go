package task9

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 9. Filialda qancha summalik product borligi jadvali:
// 1. Branch1        853 000
// 2. Branch2      1 982 000

func CalculateProductSum() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		SELECT b.name, SUM(p.price*bp.quantity) 
		FROM branch_products bp
		JOIN product p ON p.id = bp.product_id
		JOIN branch b ON b.id = bp.branch_id 
		GROUP BY b.name
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
