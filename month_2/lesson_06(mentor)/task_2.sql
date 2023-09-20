-- 2. json bo'yicha masalalarni birinchi 5tasini function bilan qilish(xohlaganlar hammasini qilsin)


-- ==================================================================================================
-- // 1.transactionlar soni bo'yicha top branches
CREATE OR REPLACE FUNCTION task1()
    RETURNS TABLE(branch_name VARCHAR, tr_count INT)
    LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY (
    SELECT name AS branch_name, COUNT(t.branch_id)::INT AS tr_count
	    FROM branch 
	    JOIN branch_transaction t ON branch.id = t.branch_id
	    GROUP BY name
	    ORDER BY tr_count DESC
    );
END;
$$;

-- ==================================================================================================
-- // 2.transactionlar summasi bo'yicha top branchess
CREATE OR REPLACE FUNCTION task2()
    RETURNS TABLE(branch_name VARCHAR, SUMMA INT)
    LANGUAGE plpgsql
AS $$
DECLARE
    total_sum INT;
BEGIN 
    RETURN QUERY (
    SELECT b.name as branch_name, CAST(SUM(p.price*t.quantity) AS int) as total_sum from branch b 
	    join branch_transaction t on t.branch_id = b.id 
	    join product p on p.id = t.product_id
	    group by branch_name order by total_sum desc
    );
END;
$$;

-- ==================================================================================================
-- // 3.transactionda bo'lgan top productlar
CREATE OR REPLACE FUNCTION task3() 
    RETURNS TABLE (product_name VARCHAR, tr_count int)
    LANGUAGE plpgsql
AS $$
BEGIN 
    RETURN QUERY (
        SELECT p.name as product_name, CAST(count(t.product_id) as int) as tr_count from branch_transaction t 
	    join product p on p.id = t.product_id 
	    group by product_name order by tr_count desc
    );
END;
$$;

-- ==================================================================================================
-- // 4.transactionda bo'lgan top categorylar
CREATE OR REPLACE FUNCTION task4()
    RETURNS TABLE (category_name VARCHAR, tr_count int) 
    LANGUAGE plpgsql
AS $$
BEGIN 
    RETURN QUERY (
    SELECT c.name AS category_name, cast(COUNT(p.id) as int) AS transaction_count
	    FROM branch_transaction bt
	    JOIN product p ON bt.product_id = p.id
	    JOIN category c ON p.category_id = c.id
	    GROUP BY c.id
	    ORDER BY transaction_count DESC
    );
END;
$$;

-- ==================================================================================================
-- // 5.har bir branchda har bir categorydan qancha transaction bo'lgani
CREATE OR REPLACE FUNCTION task5()
    RETURNS TABLE(branch_name VARCHAR, category_name VARCHAR, tr_count int) 
    LANGUAGE plpgsql
AS $$
BEGIN 
    RETURN QUERY (
    SELECT b.name as branch_name, c.name as cat_name, cast(count(p.id) as int) as tr_count
	    FROM branch b
	    JOIN branch_transaction t ON b.id = t.branch_id
	    JOIN product p on p.id = t.product_id
	    JOIN category c on c.id = p.category_id
	    GROUP BY branch_name, cat_name
    );
END;
$$;