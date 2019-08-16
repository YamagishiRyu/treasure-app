-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE mdfiles (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  path VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  repository_id INT UNSIGNED NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(repository_id) REFERENCES repositories(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE mdfiles;
