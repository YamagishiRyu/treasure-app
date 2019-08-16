package models

type Repository struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Path string `db:"path" json:"path"`
	URL  string `db:"string" json:"string"`
}
