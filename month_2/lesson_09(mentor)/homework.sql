-- HOMEWORK
-- miniProject
-- 1. Sale create/cancel qilinganda stafflarni balansiga pul 
-- hisoblanishini function/block ichida transaction bilan qiling

-- 2. agar staff bir kunda 10tadan ko'p va umumiy summasi 1 500 000 dan ko'p savdo qilsa 
-- 50 000 bonus berilishi kerak: balancega qo'shilishi va staff_transaction create qilinishi trigger bilan qiling
CREATE OR REPLACE FUNCTION give_bonus_trigger()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS $$
BEGIN
    DECLARE
        total_sales INTEGER;
        total_sum INTEGER;
    BEGIN
        -- Calculate the total sales count and sum for the staff member
        SELECT COUNT(*), SUM(sale_amount)
        INTO total_sales, total_sum
        FROM staff_sales
        WHERE staff_id = NEW.staff_id;

        -- Check if the conditions are met for giving the bonus
        IF total_sales > 10 AND total_sum > 1500000 THEN
            -- Add the bonus amount to the staff's balance
            UPDATE staff
            SET balance = balance + 50000
            WHERE id = NEW.staff_id;

            -- Insert a new transaction record in the staff_transaction table
            INSERT INTO staff_transaction (staff_id, bonus_amount)
            VALUES (NEW.staff_id, 50000);
        END IF;

        RETURN NEW;
    END;
END;
$$;

-- 3. pg_stat_activitydan username,query, exec_time ni exec time kamayish tartibida chiqaring
SELECT usename, query, clock_timestamp() - query_start AS exec_time
FROM pg_stat_activity
WHERE state = 'active'
ORDER BY exec_time DESC;

-- 4. pg_stat_user_indexesdan tablename, index_name, qancha ishlatilgani o'sish tartibida chiqaring
SELECT relname, indexrelname, idx_scan
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;
