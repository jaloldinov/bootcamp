CREATE TABLE IF NOT EXISTS branch (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(30),
    adress VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user" (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL DEFAULT 'username',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS category (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL DEFAULT 'category name',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL DEFAULT 'product name',
    price NUMERIC(20,2),
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS branch_products (
    branch_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (branch_id) REFERENCES branch(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);

CREATE TABLE IF NOT EXISTS branch_transaction (
    id INT NOT NULL PRIMARY KEY,
    branch_id INT NOT NULL,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    type VARCHAR(5) NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (branch_id) REFERENCES branch(id),
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (user_id) REFERENCES "user"(id)
);