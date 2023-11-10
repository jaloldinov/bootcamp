package task14

/*
15. Har kuni o'rtacha user qancha summa product kiritgani va chiqargani bo'yicha jadval:
    branch      o'rtacha+        o'rtacha-
1. Anvar          370 000         435 000
2. Shuhrat        60 000          875 000
...    ...     ...    ...       ...
*/

func Task14() {

	query := `
			SELECT u.name, 
			SUM(CASE WHEN bt.type='plus' THEN (bt.quantity*p.price) ELSE 0 END)/COUNT(bt.created_at) as "plus",
			SUM(CASE WHEN bt.type='minus' THEN (bt.quantity*p.price) ELSE 0 END)/COUNT(bt.created_at) as "minus"
			FROM "user" u
			JOIN branch_transaction bt ON bt.user_id = u.id
			JOIN product p ON p.id = bt.product_id
			GROUP BY u.name
	`
}
