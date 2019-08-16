package models

type User struct {
	ID       int64  `db:"id" json:"id"`
	GithubID string `db:"github_id" json:"github_id"`
	Email    string `db:"email" json:"email"`
}
