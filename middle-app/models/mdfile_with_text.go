package models

type MdfileWithText struct {
	ID           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Text         string `db:"text" json:"text"`
	URL          string `db:"url" json:"url"`
	RepositoryID int64  `db:"repository_id" json:"repository_id"`
}
