package repositories

import (
	"github.com/YamagishiRyu/treasure-app/middle-app/models"
	"github.com/jmoiron/sqlx"
)

func AllRepositories(db *sqlx.DB, user_id int64) ([]models.Repository, error) {
	repos := make([]models.Repository, 0)

	query := `
	SELECT repo.id, repo.name, repo.path, repo.url FROM clones as c INNER JOIN repositories as repo ON c.repository_id = repo.id WHERE c.user_id = ? 
	`
	if err := db.Select(&repos, query, user_id); err != nil {
		return nil, err
	}
	return repos, nil
}
