package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/YamagishiRyu/treasure-app/middle-app/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
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

func FindRepository(db *sqlx.DB, name string) (*models.Repository, error) {
	repo := models.Repository{}

	query := `SELECT id, name, path, url FROM repositories WHERE name = ?`
	if err := db.Get(&repo, query, name); err != nil {
		log.Error(err)
		return nil, err
	}

	return &repo, nil
}

func CreateRepository(db *sqlx.DB, repo *models.Repository) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO repositories (name, path, url) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(repo.Name, repo.Path, repo.URL)
}

func CreateClone(db *sqlx.DB, user_id int64, repo_id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO clones (user_id, repository_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(user_id, repo_id)
}

func CreateMdfiles(db *sqlx.DB, mdfiles []models.Mdfile) (sql.Result, error) {
	valueString := make([]string, 0, len(mdfiles))
	valueArgs := make([]interface{}, 0, len(mdfiles)*3)
	for _, mdfile := range mdfiles {
		valueString = append(valueString, "(?, ?, ?)")
		valueArgs = append(valueArgs, mdfile.Name)
		valueArgs = append(valueArgs, mdfile.Path)
		valueArgs = append(valueArgs, mdfile.RepositoryID)
	}
	stmt := fmt.Sprintf("INSERT INTO mdfiles (name, path, repository_id) VALUES %s", strings.Join(valueString, ","))
	return db.Exec(stmt, valueArgs...)
}

func SelectMdfilesFromRepository(db *sqlx.DB, repository_id int64) ([]models.Mdfile, error) {
	mdfiles := make([]models.Mdfile, 0)
	query := `
	SELECT id, name, path, url, repository_id FROM mdfiles WHERE repository_id = ? 
	`
	if err := db.Select(&mdfiles, query, repository_id); err != nil {
		return nil, err
	}
	return mdfiles, nil
}
