package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"github.com/YamagishiRyu/treasure-app/middle-app/models"
	"github.com/YamagishiRyu/treasure-app/middle-app/repositories"
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
		log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var re models.GitResponse
	if err := json.Unmarshal(reader, &re); err != nil {
		log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, re.Items, nil
}

func (rc *RepositoryController) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	userId := int64(1)
	repos, err := repositories.AllRepositories(rc.dbx, userId)
	if err != nil {
		log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, repos, nil
}

func (rc *RepositoryController) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	repoNames, ok := r.URL.Query()["repo_name"]
	repoName := repoNames[0]
	if !ok || len(repoName) < 1 {
		log.Error("Url Param 'key' is missing")
		return http.StatusBadRequest, nil, errors.New("a")
	}

	userId := int64(1)
	response := map[string]string{}

	repo, err := repositories.FindRepository(rc.dbx, repoName)
	if err != nil {
		// new repository
		cmd := exec.Command("sh", "./clone_cmd.sh", repoName)
		_, err := cmd.Output()
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}

		fmt.Print("---------- clone done ------------\n")

		github_host := "https://github.com/"
		repo := models.Repository{Name: repoName, Path: "files/" + repoName, URL: github_host + repoName}
		result, err := repositories.CreateRepository(rc.dbx, &repo)
		repoId, err := result.LastInsertId()
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}

		fmt.Print("---------- url fetch done ------------\n")

		_, err = repositories.CreateClone(rc.dbx, userId, repoId)
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}

		fmt.Print("---------- create clone done ------------\n")

		findCmd := exec.Command("sh", "./find_cmd.sh", repoName)
		pathList, err := findCmd.Output()
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}

		fmt.Print("---------- find mdfile done ------------\n")

		pathSlice := strings.Split(string(pathList), "\n")
		mdfiles := []models.Mdfile{}
		for _, path := range pathSlice {
			str := []byte(path)
			assigned := regexp.MustCompile("files/" + repoName + "/(.*)")
			group := assigned.FindSubmatch(str)
			if len(group) > 1 {
				mdfiles = append(mdfiles, models.Mdfile{Name: string(group[1]), Path: path, RepositoryID: repoId})
			}
		}

		fmt.Print("---------- create mdfile models done ------------\n")

		mdfileResult, err := repositories.CreateMdfiles(rc.dbx, mdfiles)
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}

		count, err := mdfileResult.RowsAffected()
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}

		fmt.Print("---------- save mdfiles done ------------\n")
		response["message"] = fmt.Sprintf("%v mdfiles found", count)
	} else {
		_, err = repositories.CreateClone(rc.dbx, userId, repo.ID)
		if err != nil {
			log.Error(err)
			return http.StatusBadRequest, nil, err
		}
		response["message"] = "already cloned"
	}

	return http.StatusCreated, response, nil
}
