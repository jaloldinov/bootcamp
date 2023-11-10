package task5

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 5.har bir branchda har bir categorydan qancha transaction bo'lgani
func TopBranchTransactionCategory() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		SELECT b.name as branch_name, c.name as cat_name, count(p.id) as tr_count
		FROM branch b
		JOIN branch_transaction t ON b.id = t.branch_id
		JOIN product p on p.id = t.product_id
		JOIN category c on c.id = p.category_id
		GROUP BY branch_name, cat_name
	`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task5
		err := rows.Scan(&t.BranchName, &t.CategoryName, &t.Count)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %s - %d\n", t.BranchName, t.CategoryName, t.Count)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

}
