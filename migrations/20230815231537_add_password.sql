-- +goose Up
ALTER TABLE users
ADD COLUMN password TEXT
DEFAULT encode('test_password'::bytea, 'base64');

-- +goose Down
ALTER TABLE users
DROP COLUMN password;