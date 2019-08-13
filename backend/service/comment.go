package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Comment struct {
	db *sqlx.DB
}

func NewCommentService(db *sqlx.DB) *Comment {
	return &Comment{db}
}

func (c *Comment) Create(newComment *model.Comment) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(c.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateComment(tx, newComment)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed article insert transaction")
	}
	return createdId, nil
}
