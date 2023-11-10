CREATE TYPE "product_type" AS ENUM (
  'modifier',
  'product'
);

CREATE TYPE "discount_type" AS ENUM (
  'sum',
  'percent'
);

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

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "title" varchar(100) NOT NULL,
  "image" varchar NOT NULL,
  "active" bool DEFAULT true,
  "parent_id" int,
  "order_number" serial,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "title" varchar(100) NOT NULL,
  "description" text NOT NULL,
  "photo" varchar NOT NULL,
  "order_number" serial,
  "active" bool DEFAULT true,
  "type" product_type NOT NULL,
  "price" numeric NOT NULL,
  "category_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "clients" (
  "id" serial PRIMARY KEY,
  "first_name" varchar(20) NOT NULL,
  "last_name" varchar(20),
  "phone" varchar(13) UNIQUE NOT NULL,
  "photo" varchar NOT NULL,
  "birth_date" date NOT NULL,
  "last_ordered_date" timestamp,
  "total_orders_sum" numeric DEFAULT 0,
  "total_orders_count" int DEFAULT 0,
  "discount_type" discount_type,
  "discount_amount" numeric,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "couriers" (
  "id" serial PRIMARY KEY,
  "first_name" varchar(20) NOT NULL,
  "last_name" varchar(20),
  "phone" varchar(13) UNIQUE NOT NULL,
  "active" bool DEFAULT true,
  "login" varchar(20) UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "max_order_count" int NOT NULL,
  "branch_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "first_name" varchar(20) NOT NULL,
  "last_name" varchar(20),
  "branch_id" int NOT NULL,
  "phone" varchar(13) UNIQUE NOT NULL,
  "active" bool DEFAULT true,
  "login" varchar(20) UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "branches" (
  "id" serial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "phone" varchar(13) NOT NULL,
  "photo" varchar(255) NOT NULL,
  "delivery_tarif_id" int NOT NULL,
  "work_hour_start" time NOT NULL,
  "work_hout_end" time NOT NULL,
  "address" text NOT NULL,
  "destination" text NOT NULL,
  "active" bool DEFAULT true,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
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
  "order_id" int NOT NULL,
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
  "delivery_tarif_id" int NOT NULL,
  "from_price" numeric,
  "to_price" numeric,
  "price" numeric,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "categories" ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "couriers" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "order_products" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "delivery_tarif_values" ADD FOREIGN KEY ("delivery_tarif_id") REFERENCES "delivery_tarif" ("id");

