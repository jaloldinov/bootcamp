package task11

// 11. har bir user kun bo'yicha nechta va necha sumlik transaction qilgani jadvali:

// 1 Akrom 2023-01-01  12  382 000
// 2 Suhrob 2023-03-05  32  745 000
func Task11() {

	query := `
		SELECT u.name, bt.created_at AS vaqt, COUNT(bt.quantity), SUM(bt.quantity * p.price)
		FROM  branch_transaction bt
  		JOIN product p ON bt.product_id =p.id
		JOIN "user" u ON u.id=bt.user_id
  		GROUP BY u.name, bt.created_at
	`
}
