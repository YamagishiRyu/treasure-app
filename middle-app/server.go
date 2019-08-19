package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	controllers "github.com/YamagishiRyu/treasure-app/middle-app/controllers"
	db "github.com/YamagishiRyu/treasure-app/middle-app/db"
	"github.com/YamagishiRyu/treasure-app/middle-app/middleware"
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
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	repositoryController := controllers.NewRepositoryController(s.dbx)
	repo := r.PathPrefix("/repositories").Subrouter()

	repo.Methods(http.MethodGet).Path("/").Handler(commonChain.Then(AppHandler{repositoryController.Index}))
	repo.Methods(http.MethodGet).Path("/show/{id}").Handler(commonChain.Then(AppHandler{repositoryController.Show}))
	repo.Methods(http.MethodGet).Path("/clone").Handler(commonChain.Then(AppHandler{repositoryController.Create}))
	repo.Methods(http.MethodGet).Path("/search").Handler(commonChain.Then(AppHandler{repositoryController.Search}))

	return r
}
