package task7

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

//  7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish:
//     kun         soni
//  1. 2023-08-04   789
//  2. 2023-08-12   634
func CalculateProductIncome() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
		SELECT TO_CHAR(created_at, 'YYYY-MM-DD') AS Day, sum(quantity) as soni 
			FROM branch_transaction 
		WHERE type = 'plus' 
		GROUP BY TO_CHAR(created_at, 'YYYY-MM-DD')
	`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task7
		err := rows.Scan(&t.Date, &t.Sum)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d\n", t.Date, t.Sum)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
