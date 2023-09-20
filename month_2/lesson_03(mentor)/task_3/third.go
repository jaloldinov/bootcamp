package task3

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 3.transactionda bo'lgan top productlar
func TopTransactionProducts() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
	SELECT p.name as product_name, count(t.product_id) as tr_count from branch_transaction t 
	join product p on p.id = t.product_id 
	group by product_name order by tr_count desc
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
