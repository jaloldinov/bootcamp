package task1

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 1.transactionlar soni bo'yicha top branches
func CalculateTranTopBranches() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		SELECT name, COUNT(t.branch_id) AS tr_count
		FROM branch 
		JOIN branch_transaction t ON branch.id = t.branch_id
		GROUP BY name
		ORDER BY tr_count DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bt models.BranchTransactionCount
		err := rows.Scan(&bt.BranchName, &bt.Count)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d\n", bt.BranchName, bt.Count)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
