-- 3.Masala: json data bo'yicha branchdan product_id lik productdan n quantity sotish kerak.
-- Shunday plpgsql procedure/function yozing agar branchda product yetarli bo'lsa, 
-- transaction create qilib branchdagi productlardan n quantity kamaytirsin, 
-- agar kam bo'lsa 'product yetarli emas' xabari chiqsin
CREATE OR REPLACE FUNCTION sell_product_from_branch(
    branch_id_param INT,
    product_id_param INT,
    quantity_param INT
) 
RETURNS branch_products 
LANGUAGE plpgsql 
AS $$
DECLARE
    available_quantity INT;
    sold_product branch_products%ROWTYPE; 
BEGIN
    SELECT quantity INTO available_quantity
    FROM branch_products 
    WHERE branch_id = branch_id_param AND product_id = product_id_param;
    
    IF available_quantity IS NOT NULL AND available_quantity >= quantity_param THEN
        UPDATE branch_products 
        SET quantity = quantity - quantity_param
        WHERE branch_id = branch_id_param AND product_id = product_id_param
        RETURNING * INTO sold_product; 
        
        RAISE NOTICE 'Product sold successfully';
        RETURN sold_product; 
    ELSE
        RAISE EXCEPTION 'Product is not available in sufficient quantity';
    END IF;
END;
$$;