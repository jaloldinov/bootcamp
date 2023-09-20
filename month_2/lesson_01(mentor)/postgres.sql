--products(id,name,description,price,created_at,updated_at)
-- 2. products table yaratish:id primary key bo'lishi kerak
CREATE TABLE
    products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(20),
        description TEXT,
        price NUMERIC(10, 2),
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

-- 3. 1.name columnga unique va not null qo'shish
ALTER TABLE products ALTER COLUMN name
SET
    NOT NULL;

ALTER TABLE products
ADD CONSTRAINT name_unique UNIQUE (name);

-- 3. 2. uniqueni olib tashlash
ALTER TABLE products DROP CONSTRAINT name_unique;

-- 4. nameda "a" harfi qatnashgan va narxi 5000dan katta bo'lgan productlarni select qilish
SELECT
    *
FROM
    products
WHERE
    name ilike '%a%'
    AND price > 5000;

-- 5.  oxirgi 5kunda kiritilgan productlarni faqat name,created_atlarini, kiritilgan vaqtini kamayish tartibida chiqarish(between ishlatilsin)
SELECT
    name,
    created_at
from
    products
WHERE
    created_at BETWEEN CURRENT_DATE - INTERVAL '5 days' AND CURRENT_DATE
ORDER BY
    created_at DESC;

-- 6. name yoki descriptionda "burger" so'zi qatnashgan va yaratilgandan keyin update qilingan productlarni chiqarish
SELECT
    *
FROM
    products
WHERE
    name ilike '%burger%'
    or description ilike '%burger%'
    AND created_at IS NOT NULL;