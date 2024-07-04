package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"guthub.com/iribuda/todo-api-go/pkg/services/task"
	"guthub.com/iribuda/todo-api-go/pkg/services/user"
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

	// Erstellen von Repositories und Controllers
	// Dependency Injection
	userRepository := user.NewRepository(s.db)
	userController := user.NewController(userRepository)
	userController.RegisterRoutes(router)

	taskRepository := task.NewRepository(s.db)
	taskController := task.NewController(taskRepository, userRepository)
	taskController.RegisterRoutes(router)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}