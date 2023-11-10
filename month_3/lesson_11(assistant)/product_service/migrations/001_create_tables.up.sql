CREATE TABLE IF NOT EXISTS "category" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "products" (
    "id" uuid PRIMARY KEY,
    "category_id" uuid REFERENCES "cateogry"("id"),
    "name" varchar(100) NOT NULL,
    "description" varchar(1000),
    "price" float NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);
