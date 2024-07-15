CREATE TABLE "farm"
(
  "id" SERIAL PRIMARY KEY,
  "farm_name" VARCHAR(255) NOT NULL,
  "is_deleted" BOOLEAN NOT NULL,
  "created_at" timestamptz DEFAULT current_timestamp,
  "updated_at" timestamptz DEFAULT current_timestamp
);
