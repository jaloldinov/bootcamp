package task6

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
//                     Transactions            Summ
//                     plus   minus        plus     minus
//     1. Branch1      53      20          853 000  278 000
//     2. Branch2      38      185         492 000  1 982 000

func PlusMinus() {

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
SELECT 
    b.name AS branch_name,
    COUNT(CASE WHEN t.type = 'plus' THEN t.id END) AS tran_plus,
    COUNT(CASE WHEN t.type = 'minus' THEN t.id END) AS tran_minus,
    cast( SUM(CASE WHEN t.type = 'plus' THEN t.quantity * p.price END) as int) AS sum_plus, -- cast as int vscode uchun ishlatdim
    cast( SUM(CASE WHEN t.type = 'minus' THEN t.quantity * p.price END) as int) AS sum_minus
 	FROM branch AS b 
	JOIN branch_transaction AS t ON b.id = t.branch_id
	JOIN product AS p ON t.product_id = p.id
	GROUP BY b.name;
	`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task6
		err := rows.Scan(&t.BranchaName, &t.TranPlusCount, &t.TranMinusCount, &t.TranPlusSum, &t.TranMinusSum)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: t_plus: %d t_minus: %d - s_plus: %d s_minus: %d\n",
			t.BranchaName, t.TranPlusCount, t.TranMinusCount, t.TranPlusSum, t.TranMinusSum)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

}
