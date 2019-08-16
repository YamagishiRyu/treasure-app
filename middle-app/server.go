package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	controllers "github.com/YamagishiRyu/treasure-app/middle-app/controllers"
	db "github.com/YamagishiRyu/treasure-app/middle-app/db"
)

type Server struct {
	dbx    *sqlx.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	db := db.NewDB(datasource)
	dbx, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.dbx = dbx
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	r := mux.NewRouter()

	repositoryController := controllers.NewRepositoryController(s.dbx)
	repo := r.PathPrefix("/repositories").Subrouter()

	repo.Methods(http.MethodGet).Path("/").Handler(AppHandler{repositoryController.Create})
	repo.Methods(http.MethodGet).Path("/search").Handler(AppHandler{repositoryController.Search})

	return r
}
