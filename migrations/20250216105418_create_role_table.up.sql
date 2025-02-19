CREATE TABLE IF NOT EXISTS "_roles" (
    "id" BIGSERIAL PRIMARY KEY,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    "created_by_id" INT,
    "lock" BOOLEAN DEFAULT false,
    "permissions" TEXT [],
    "name" VARCHAR(255) UNIQUE NOT NULL
)