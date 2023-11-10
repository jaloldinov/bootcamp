CREATE TYPE "order_type" AS ENUM (
  'delivery',
  'pick_up'
);

CREATE TYPE "order_status" AS ENUM (
  'accepted',
  'courier_accepted',
  'ready_in_branch',
  'on_way',
  'finished',
  'canceled'
);

CREATE TYPE "order_payment_type" AS ENUM (
  'cash',
  'card'
);

CREATE TYPE "delivery_tarif_type" AS ENUM (
  'fixed',
  'alternative'
);

CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "order_id" varchar(6) UNIQUE NOT NULL,
  "client_id" int NOT NULL,
  "branch_id" int NOT NULL,
  "type" order_type NOT NULL,
  "address" text,
  "courier_id" int,
  "price" numeric NOT NULL,
  "delivery_price" numeric,
  "discount" numeric,
  "status" order_status,
  "payment_type" order_payment_type,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "order_products" (
  "order_id" varchar NOT NULL REFERENCES "orders"("order_id"),
  "product_id" int,
  "quantity" int,
  "price" numeric,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "delivery_tarif" (
  "id" serial PRIMARY KEY,
  "name" varchar(255),
  "type" delivery_tarif_type,
  "base_price" numeric,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "delivery_tarif_values" (
  "delivery_tarif_id" int NOT NULL REFERENCES "delivery_tarif"("id"),
  "from_price" numeric,
  "to_price" numeric,
  "price" numeric,
);
