-- +migrate Up
CREATE TABLE IF NOT EXISTS users_quotes (
    id  SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    quotes_id varchar(100) NOT NULL,
    quotes TEXT NOT NULL,
    author varchar(50),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


-- +migrate Down
DROP TABLE users_quotes;
