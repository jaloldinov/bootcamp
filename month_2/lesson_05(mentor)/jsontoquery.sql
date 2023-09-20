-- // 1.transactionlar soni bo'yicha top branches
SELECT name, COUNT(t.branch_id) AS tr_count
	FROM branch 
	JOIN branch_transaction t ON branch.id = t.branch_id
	GROUP BY name
	ORDER BY tr_count DESC

-- // 2.transactionlar summasi bo'yicha top branchess
SELECT b.name as branch_name, CAST(SUM(p.price*t.quantity) AS int) as total_sum from branch b 
	join branch_transaction t on t.branch_id = b.id 
	join product p on p.id = t.product_id
	group by branch_name order by total_sum desc
	

-- // 3.transactionda bo'lgan top productlar
SELECT p.name as product_name, count(t.product_id) as tr_count from branch_transaction t 
	join product p on p.id = t.product_id 
	group by product_name order by tr_count desc
	

-- // 4.transactionda bo'lgan top categorylar
SELECT c.name AS category_name, COUNT(p.id) AS transaction_count
	FROM branch_transaction bt
	JOIN product p ON bt.product_id = p.id
	JOIN category c ON p.category_id = c.id
	GROUP BY c.id
	ORDER BY transaction_count DESC

-- // 5.har bir branchda har bir categorydan qancha transaction bo'lgani
SELECT b.name as branch_name, c.name as cat_name, count(p.id) as tr_count
	FROM branch b
	JOIN branch_transaction t ON b.id = t.branch_id
	JOIN product p on p.id = t.product_id
	JOIN category c on c.id = p.category_id
	GROUP BY branch_name, cat_name

-- // 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
-- //                     Transactions            Summ
-- //                     plus   minus        plus     minus
-- //     1. Branch1      53      20          853 000  278 000
-- //     2. Branch2      38      185         492 000  1 982 000
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

-- //  7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish:
-- //     kun         soni
-- //  1. 2023-08-04   789
-- //  2. 2023-08-12   634
SELECT TO_CHAR(created_at, 'YYYY-MM-DD') AS Day, sum(quantity) as soni 
	FROM branch_transaction 
	WHERE type = 'plus' 
	GROUP BY TO_CHAR(created_at, 'YYYY-MM-DD')
	

-- // 8. Product qancha kiritilgan va chiqarilganligi jadvali:
-- //     Name    Kiritilgan  Chiqarilgan
-- //     Olma     345            847
-- //     Cola     374            219
-- //     ....     ...       ...   ....
-- // products, transaction
SELECT p.name,
	SUM(CASE WHEN t.type = 'plus' THEN t.quantity ELSE 0 END) as kiritilgan,
	SUM(CASE WHEN t.type = 'minus' THEN t.quantity ELSE 0 END) as chiqarilgan
	FROM product AS p 
	JOIN branch_transaction AS t ON p.id = t.product_id
	GROUP BY p.name;


-- // 9. Filialda qancha summalik product borligi jadvali:
-- // 1. Branch1        853 000
-- // 2. Branch2      1 982 000
SELECT b.name, SUM(p.price*bp.quantity) 
	FROM branch_products bp
	JOIN product p ON p.id = bp.product_id
	JOIN branch b ON b.id = bp.branch_id 
    GROUP BY b.name

-- // 10. har bir user transaction qilgan summasi jadvali:
-- // 1  Akrom   1 349 000
-- // 3  Ilhom   2 974 000
SELECT u.name, SUM(p.price*bt.quantity) 
	FROM branch_transaction bt
	JOIN product p ON p.id = bt.product_id
	JOIN branch b ON b.id = bt.branch_id 
	JOIN "user" u ON u.id = bt.user_id
	GROUP BY u.name

-- // 11. har bir user kun bo'yicha nechta va necha sumlik transaction qilgani jadvali:
-- // 1 Akrom 2023-01-01  12  382 000
-- // 2 Suhrob 2023-03-05  32  745 000
SELECT u.name, bt.created_at AS vaqt, COUNT(bt.quantity), SUM(bt.quantity * p.price)
	FROM  branch_transaction bt
	JOIN product p ON bt.product_id =p.id
	JOIN "user" u ON u.id=bt.user_id
	GROUP BY u.name, bt.created_at

-- // 12. har bir user qancha product kiritgani va chiqargani jadvali:
-- //          kiritgan  chiqargan
-- // 1 Akrom    12         84
-- // 2 Suhrob   54         33
SELECT u.name, 
	SUM(CASE WHEN bt.type = 'plus' THEN bt.quantity ELSE 0 END) as kiritilgan, 
	SUM(CASE WHEN bt.type = 'minus' THEN bt.quantity ELSE 0 END) as chiqarilgan
	FROM branch_transaction bt
	JOIN "user" u ON u.id = bt.user_id
	GROUP BY u.name

-- //  13. Har bir kunda o'rtacha qancha product kiritilgani va chiqarilgani bo'yicha jadval:
-- //     branch      o'rtacha+   o'rtacha-
-- //  1. Chilonzor      73         34
-- //  2. MGorkiy        60         75
SELECT b.name, 
	SUM(CASE WHEN bt.type='plus' THEN (bt.quantity) ELSE 0 END)/COUNT(bt.created_at) as "plus",
	SUM(CASE WHEN bt.type='minus' THEN (bt.quantity) ELSE 0 END)/COUNT(bt.created_at) as "minus"
	FROM branch b
	JOIN branch_transaction bt ON bt.branch_id = b.id
	GROUP BY b.name

-- 15. Har kuni o'rtacha user qancha summa product kiritgani va chiqargani bo'yicha jadval:
--     branch      o'rtacha+        o'rtacha-
-- 1. Anvar          370 000         435 000
-- 2. Shuhrat        60 000          875 000
SELECT u.name, 
	SUM(CASE WHEN bt.type='plus' THEN (bt.quantity*p.price) ELSE 0 END)/COUNT(bt.created_at) as "plus",
	SUM(CASE WHEN bt.type='minus' THEN (bt.quantity*p.price) ELSE 0 END)/COUNT(bt.created_at) as "minus"
	FROM "user" u
	JOIN branch_transaction bt ON bt.user_id = u.id
	JOIN product p ON p.id = bt.product_id
	GROUP BY u.name