package task8

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 8. Product qancha kiritilgan va chiqarilganligi jadvali:
//     Name    Kiritilgan  Chiqarilgan
//     Olma     345            847
//     Cola     374            219
//     ....     ...       ...   ....
// products, transaction

func Task8() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// =========== EXAMPLE ==============
	// SELECT purchases.seller_id,
	//   SUM(CASE WHEN state IN ('authorized', 'reversed') THEN 1 ELSE 0 END) AS sales_count,
	//   SUM(CASE WHEN state = 'authorized' THEN 1 ELSE 0 END) AS successful_sales_count
	// FROM purchases
	// GROUP BY purchases.seller_id
	//=========================================================

	query := `
		SELECT p.name,
		SUM(CASE WHEN t.type = 'plus' THEN t.quantity ELSE 0 END) as kiritilgan,
		SUM(CASE WHEN t.type = 'minus' THEN t.quantity ELSE 0 END) as chiqarilgan
		FROM product AS p 
	 	JOIN branch_transaction AS t ON p.id = t.product_id
		GROUP BY p.name;
	`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task8
		err := rows.Scan(&t.Name, &t.PlusCount, &t.MinusCount)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d - %d\n", t.Name, t.PlusCount, t.MinusCount)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
