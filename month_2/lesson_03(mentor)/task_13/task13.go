package task13

//  13. Har bir kunda o'rtacha qancha product kiritilgani va chiqarilgani bo'yicha jadval:
//     branch      o'rtacha+   o'rtacha-
//  1. Chilonzor      73         34
//  2. MGorkiy        60         75

func Task13() {
	query := `

		SELECT b.name, 
		SUM(CASE WHEN bt.type='plus' THEN (bt.quantity) ELSE 0 END)/COUNT(bt.created_at) as "plus",
		SUM(CASE WHEN bt.type='minus' THEN (bt.quantity) ELSE 0 END)/COUNT(bt.created_at) as "minus"
		FROM branch b
		JOIN branch_transaction bt ON bt.branch_id = b.id
		GROUP BY b.name
`
}
