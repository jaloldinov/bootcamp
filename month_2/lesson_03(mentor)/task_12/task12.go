package task12

// 12. har bir user qancha product kiritgani va chiqargani jadvali:
//          kiritgan  chiqargan
// 1 Akrom    12         84
// 2 Suhrob   54         33

func Task12() {
	query := `
		SELECT u.name, 
		SUM(CASE WHEN bt.type = 'plus' THEN bt.quantity ELSE 0 END) as kiritilgan, 
		SUM(CASE WHEN bt.type = 'minus' THEN bt.quantity ELSE 0 END) as chiqarilgan
		FROM branch_transaction bt
		JOIN "user" u ON u.id = bt.user_id
		GROUP BY u.name
	`
}
