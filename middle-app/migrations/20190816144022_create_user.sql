-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  github_id VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
