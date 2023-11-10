CREATE TYPE "discount_type" AS ENUM (
  'sum',
  'percent'
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

ALTER TABLE "couriers" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");
