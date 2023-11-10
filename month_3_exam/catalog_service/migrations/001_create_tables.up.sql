CREATE TYPE product_type AS ENUM (
  'modifier',
  'product'
);

CREATE TABLE categories (
  id serial PRIMARY KEY,
  title varchar(100) NOT NULL,
  image varchar NOT NULL,
  active bool DEFAULT true,
  parent_id int REFERENCES categories(id),
  order_number serial,
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE products (
  id serial PRIMARY KEY,
  title varchar(100) NOT NULL,
  description text NOT NULL,
  photo varchar NOT NULL,
  order_number serial,
  active bool DEFAULT true,
  type product_type NOT NULL,
  price numeric NOT NULL,
  category_id int NOT NULL REFERENCES categories(id),
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp,
  deleted_at timestamp
);