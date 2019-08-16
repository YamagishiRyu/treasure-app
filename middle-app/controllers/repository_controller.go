package controller

import (
	"encoding/json"
	"net/http"

	"github.com/YamagishiRyu/treasure-app/middle-app/models"
	"github.com/jmoiron/sqlx"
)

type RepositoryController struct {
	dbx *sqlx.DB
}

func NewRepositoryController(dbx *sqlx.DB) *RepositoryController {
	return &RepositoryController{dbx: dbx}
}

func (rc *RepositoryController) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newRepository := &models.Repository{}
	if err := json.NewDecoder(r.Body).Decode(&newRepository); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, newRepository, nil
}
