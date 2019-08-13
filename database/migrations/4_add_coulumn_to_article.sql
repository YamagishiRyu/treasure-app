-- +goose Up
ALTER TABLE article ADD user_id INT UNSIGNED REFERENCES user(id);

-- +goose Down
ALTER TABLE article DROP COLUMN user_id;
