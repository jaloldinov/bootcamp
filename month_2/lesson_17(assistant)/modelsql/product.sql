CREATE TABLE "product" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price"  int, 
    "category_id" UUID REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);