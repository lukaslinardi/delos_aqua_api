CREATE TABLE "pond"
(
  "id" SERIAL PRIMARY KEY,
  "pond_name" VARCHAR(255) NOT NULL,
  "is_deleted" BOOLEAN NOT NULL,
  "farm_id" INT NOT NULL,
  "created_at" timestamptz DEFAULT current_timestamp,
  "updated_at" timestamptz DEFAULT current_timestamp
);
