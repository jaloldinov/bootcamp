CREATE TYPE "tariff_type" AS ENUM (
  'fixed',
  'percent'
);

CREATE TYPE "staff_type" AS ENUM (
  'cashier',
  'shop_assistant'
);

CREATE TABLE "tariffs" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "type" tariff_type NOT NULL,
  "amount_for_cash" NUMERIC(12, 2),
  "amount_for_card" NUMERIC(12, 2),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "staffs" (
  "id" uuid PRIMARY KEY,
  "name" varchar not null,
  "branch_id" uuid NOT NULL REFERENCES "branches"("id"),
  "tariff_id" uuid NOT NULL REFERENCES "tariffs"("id"),
  "staff_type" staff_type NOT NULL,
  "balance" NUMERIC(12, 2) NOT NULL DEFAULT 0.0,
  "username" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);