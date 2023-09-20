package task2

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 2.transactionlar summasi bo'yicha top branchess
func CalculateSumOfPriceTopBranches() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
	SELECT b.name as branch_name, CAST(SUM(p.price*t.quantity) AS int) as total_sum from branch b 
	join branch_transaction t on t.branch_id = b.id 
	join product p on p.id = t.product_id
	group by branch_name order by total_sum desc
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
