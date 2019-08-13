-- +goose Up
DROP TABLE comments;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` INT UNSIGNED REFERENCES article(id),
  `user_id` INT UNSIGNED REFERENCES user(id),
  `body` TEXT NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='list of comments';

-- +goose Down
DROP TABLE comment;
