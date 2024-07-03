package api

import (
	"database/sql"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"guthub.com/iribuda/todo-api-go/pkg/controllers"
	"guthub.com/iribuda/todo-api-go/pkg/repositories"
)

// Typ für Kapselung der Serveradresse und des Zeigers auf eine Datenbankverbindung
type APIServer struct{
	addr string
	db *sql.DB
}

// Dependency Injection durch Konstruktor
func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

// Einrichting des Servers
func (s *APIServer) Run() error{
	router := mux.NewRouter()

	taskRepository := repositories.NewRepository(s.db)
	taskController := controllers.NewController(taskRepository)
	taskController.RegisterRoutes(router)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}