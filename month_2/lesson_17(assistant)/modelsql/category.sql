CREATE TABLE "category" (
    "id" UUID PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "parent_id" UUID REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);