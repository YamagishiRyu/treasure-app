package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type RepositoryController struct {
	dbx *sqlx.DB
}

func NewRepositoryController(dbx *sqlx.DB) *RepositoryController {
	return &RepositoryController{dbx: dbx}
}

func (rc *RepositoryController) Search(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	host := "https://api.github.com/"
	q, ok := r.URL.Query()["q"]
	if !ok || len(q[0]) < 1 {
		log.Error("Url Param 'key' is missing")
		return http.StatusBadRequest, nil, errors.New("a")
	}

	resp, err := http.Get(host + "search/repositories?sort=stars&q=" + q[0])
	if err != nil {
		log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	defer resp.Body.Close()

	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	fmt.Println("aaa")
	fmt.Printf("%v", reader)
	fmt.Println("bbb")
	v, err := fmt.Println(reader)
	if err != nil {
		fmt.Println(err)
		log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	fmt.Print(v)

	return http.StatusOK, string(reader), nil
}

func (rc *RepositoryController) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	return http.StatusCreated, "aaa", nil
}
