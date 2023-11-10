CREATE TABLE "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "adress" varchar NOT NULL,
  "year" INT NOT NULL,
  "founded_at" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "branch_products" (
    "id" uuid PRIMARY KEY,
    "product_id" uuid NOT NULL,
    "branch_id" uuid NOT NULL REFERENCES "branches"("id"),
    "quantity" INT NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);
