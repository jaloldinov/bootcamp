CREATE TABLE users (
    "id" uuid PRIMARY KEY,
    "login" VARCHAR UNIQUE NOT NULL,
    "password" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    "age" NUMERIC NOT NULL,
    "phone_number" VARCHAR UNIQUE NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);