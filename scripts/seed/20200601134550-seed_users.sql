-- +migrate Up
INSERT INTO users (name, email, password) VALUES (
    'Test',
    'test@email.com',
    crypt('12345678',gen_salt('bf'))
);

-- +migrate Down
DELETE FROM users WHERE email = 'test@email.com';