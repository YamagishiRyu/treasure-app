-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE clones (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  repository_id INT UNSIGNED NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(repository_id) REFERENCES repositories(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE clones;
