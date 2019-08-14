package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllArticle(db *sqlx.DB) ([]model.Article, error) {
	a := make([]model.Article, 0)
	if err := db.Select(&a, `SELECT id, title, body FROM article`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindArticle(db *sqlx.DB, id int64) (*model.ArticleWithTag, error) {
	a := model.ArticleWithTag{}
	if err := db.Get(&a, `
SELECT id, title, body FROM article WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	tags := make([]model.Tag, 0)
	if err := db.Select(&tags, `
SELECT tag.id, tag.name FROM tag join tag_article ON tag.id = tag_article.tag_id WHERE tag_article.article_id = ?
`, id); err != nil {
		return nil, err
	}
	a.Tags = tags
	return &a, nil
}

func CreateArticle(db *sqlx.Tx, a *model.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article (title, body, user_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, a.UserID)
}

func UpdateArticle(db *sqlx.Tx, id int64, a *model.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article SET title = ?, body = ?, user_id = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, a.UserID, id)
}

func DestroyArticle(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
