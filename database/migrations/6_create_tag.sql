-- +goose Up
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL UNIQUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tag';

CREATE TABLE `tag_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` INT UNSIGNED,
  `tag_id` INT, 
  PRIMARY KEY (`id`),
  FOREIGN KEY (article_id) REFERENCES article(id),
  FOREIGN KEY (tag_id) REFERENCES tag(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tag article middle table';

-- +goose Down
DROP TABLE tag_article;
DROP TABLE tag;
