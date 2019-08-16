package models

type Mdfile struct {
	ID           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Path         string `db:"path" json:"path"`
	URL          string `db:"url" json:"url"`
	RepositoryID int64  `db:"repository_id" json:"repository_id"`
}
