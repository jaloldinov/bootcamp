-- jsonData:
-- 1. 'minus' branch taransaction create bo'lishida, 
-- sklad(branch_products)da quantity yetarliligini tekshiradigan trigger yarating. 
-- quantity yetmasa error(exception) qaytarsin
-- 1 ============================================================================================
CREATE OR REPLACE FUNCTION check_quantity()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
BEGIN
    IF NEW.type = 'minus' AND NEW.quantity > (
        SELECT quantity
        FROM branch_products
        WHERE branch_id = NEW.branch_id
          AND product_id = NEW.product_id
        LIMIT 1
    ) THEN
        RAISE EXCEPTION 'Product quantity not available';
    END IF;
    RETURN NEW;
END;
$$;

CREATE TRIGGER check_quantity_trigger
BEFORE INSERT ON branch_transaction
FOR EACH ROW
EXECUTE FUNCTION check_quantity();
-- update branch_transaction set type='minus', quantity = 200 where id = 18 and product_id = 2 and branch_id =1;
-- DROP TRIGGER check_quantity ON branch_transaction

-- 2 ============================================================================================
-- 2. transaction create bo'lganda skladdagi productlarni quantitysini o'zgartiradigan trigger yozing. 
-- agar skladda yo'q product bo'yicha transaction yaratilsa: plus bo'lsa skladga yangi qo'shiladi,
--  minus bo'lsa error qaytariladi
CREATE OR REPLACE FUNCTION update_branch_quantity()
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
    IF NEW.type = 'plus' THEN
        -- add quantity
        UPDATE branch_products
        SET quantity = quantity + NEW.quantity
        WHERE branch_id = NEW.branch_id
            AND product_id = NEW.product_id;
        -- if not found, insert it
        IF NOT FOUND THEN
            INSERT INTO branch_products (branch_id, product_id, quantity)
            VALUES (NEW.branch_id, NEW.product_id, NEW.quantity);
        END IF;
    ELSIF NEW.type = 'minus' THEN 
        -- check if enough quantity exists
        IF NEW.quantity > (
            SELECT quantity
            FROM branch_products
            WHERE branch_id = NEW.branch_id
            AND product_id = NEW.product_id
        ) THEN
            RAISE EXCEPTION 'not enough quantity in branch_products';
        ELSE 
            -- subtract quantity
            UPDATE branch_products
            SET quantity = quantity - NEW.quantity
            WHERE branch_id = NEW.branch_id AND product_id = NEW.product_id;
        END IF;
    END IF;
    
    RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER update_branch_quantity
BEFORE INSERT OR UPDATE ON branch_transaction
FOR EACH ROW
EXECUTE FUNCTION update_branch_quantity();

update branch_transaction set type='minus', quantity = 200 where id = 18 and product_id = 2 and branch_id =1;
