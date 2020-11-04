-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id          SERIAL          PRIMARY KEY,
    name        VARCHAR(255)    NOT NULL,
    email       VARCHAR (50)    UNIQUE NOT NULL,
    password    VARCHAR (100)   NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +migrate Down
DROP TABLE users;
DROP EXTENSION pgcrypto;