package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Comment struct {
	dbx *sqlx.DB
}

func NewComment(dbx *sqlx.DB) *Comment {
	return &Comment{dbx: dbx}
}

func (c *Comment) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	comments, err := repository.AllComment(c.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, comments, nil
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	article_id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	aid, err := strconv.ParseInt(article_id, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	newComment := &model.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		return http.StatusBadRequest, nil, err
	}
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newComment.UserID = &user.ID
	newComment.ArticleID = &aid

	commentService := service.NewCommentService(c.dbx)
	id, err := commentService.Create(newComment)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newComment.ID = id

	return http.StatusCreated, newComment, nil
}
