CREATE TABLE IF NOT EXISTS users (
    ID bigserial PRIMARY KEY,
    Name varchar NOT NULL,
    Age int8 NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS users_name_unique ON users(Name);