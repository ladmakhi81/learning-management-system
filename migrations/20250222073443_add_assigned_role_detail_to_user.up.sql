ALTER TABLE
    IF EXISTS "_users"
ADD
    "assigned_role_date" TIMESTAMP DEFAULT NULL;

ALTER TABLE
    IF EXISTS "_users"
ADD
    "assigned_role_by_id" INT DEFAULT NULL;