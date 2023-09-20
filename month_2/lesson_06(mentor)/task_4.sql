-- 4. array_agg() ishlatib branchda qanaqa categoriyadagi productlar borligini chiqaring
-- Example:

--  branch      |  categories
-- Chilonzor   | {meva, ichimlik, kiyim}
-- MGorkiy     | {meva}  

SELECT b.name as branch_name,  ARRAY_AGG(DISTINCT category.name) as category
FROM branch b
JOIN branch_products bp ON bp.branch_id = b.id
JOIN product p on p.id = bp.product_id
JOIN category ON category.id = p.category_id
GROUP BY b.name
