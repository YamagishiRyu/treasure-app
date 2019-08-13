package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllComment(db *sqlx.DB) ([]model.Comment, error) {
	c := make([]model.Comment, 0)
	if err := db.Select(&c, `SELECT id, body, user_id, article_id FROM comment`); err != nil {
		return nil, err
	}
	return c, nil
}

func FindComment(db *sqlx.DB, id int64) (*model.Comment, error) {
	c := model.Comment{}
	if err := db.Get(&c, `
SELECT id, body, user_id, article_id FROM comment WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &c, nil
}

func CreateComment(db *sqlx.Tx, c *model.Comment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO comment (body, user_id, article_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.Body, c.UserID, c.ArticleID)
}
