ALTER TABLE
    IF EXISTS "_users"
ADD
    "password_change_by" INT DEFAULT NULL;

ALTER TABLE
    IF EXISTS "_users"
ADD
    "password_change_date" TIMESTAMP DEFAULT NULL;