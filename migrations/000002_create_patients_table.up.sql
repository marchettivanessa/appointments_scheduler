CREATE TABLE IF NOT EXISTS scheduler.patients(
    id SERIAL PRIMARY KEY,
    date_of_birth VARCHAR,
    address VARCHAR,
    email VARCHAR NOT NULL,
    cause_of_treatment TEXT NOT NULL,
    phone_number STRING NOT NULL,
);