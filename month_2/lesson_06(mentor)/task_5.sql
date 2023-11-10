-- 5.Block/function qiling: branchni productni sotilgan soniga qarab 50dan ko'p bo'lsa 'xit', 100dan ko'p bo'lsa 'top' aks holda 'yangi' deb chiqarsin.tartibi: top->xit->yangi
-- Example:

--  product       |     status
--  Cola          |     top
--  Fanta         |     xit
--  Rio           |     yangi
CREATE OR REPLACE FUNCTION check_product() 
    RETURNS TABLE (branch_name VARCHAR, product VARCHAR, status VARCHAR) 
    LANGUAGE plpgsql
    AS 
$$
BEGIN
    RETURN QUERY
    SELECT b.name AS branch_name, p.name AS product,
        CASE WHEN bt.type = 'minus' THEN
            CASE WHEN SUM(bt.quantity) > 100 THEN 'TOP'
                 WHEN SUM(bt.quantity) > 20 THEN 'XIT'
                 ELSE 'YANGI'
            END
        END::VARCHAR AS status
    FROM branch b
    JOIN branch_transaction bt ON bt.branch_id = b.id
    JOIN product p ON p.id = bt.product_id
    GROUP BY b.name, p.name, bt.type;
END;
$$;
