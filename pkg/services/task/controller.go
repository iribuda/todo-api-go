package task

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"guthub.com/iribuda/todo-api-go/pkg/models"
	"guthub.com/iribuda/todo-api-go/pkg/services/auth"
	"guthub.com/iribuda/todo-api-go/pkg/utils"
)

// Haupt Controller, der sowohl Task- als auch User-Repositories braucht
type TaskController struct{
	taskRepository models.TaskRepository
	userRepository models.UserRepository
}

// Konstruktor
func NewController(taskRepository models.TaskRepository, userRepository models.UserRepository) *TaskController{
	return &TaskController{taskRepository: taskRepository, userRepository: userRepository}
}

//  Deklaration der Routen
// Vor jedem Handler wird die auth.WithJWTAuth() aufgerufen, damit nur angemeldete Benutzer mit laufendem Token 
// mit nur seinen Aufgaben arbeiten können. 
func (tc *TaskController) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/tasks", auth.WithJWTAuth(tc.handleGetAllTasks, tc.userRepository))
	router.HandleFunc("/tasks", auth.WithJWTAuth(tc.handleCreateTask, tc.userRepository)).Methods("POST")
	router.HandleFunc("/tasks/{id}", auth.WithJWTAuth(tc.handleGetTaskById, tc.userRepository)).Methods("GET")
	router.HandleFunc("/tasks/{id}", auth.WithJWTAuth(tc.handleDeleteTask, tc.userRepository)).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", auth.WithJWTAuth(tc.handleUpdateTask, tc.userRepository)).Methods("PUT")
	router.HandleFunc("/tasks/{id}/complete", auth.WithJWTAuth(tc.handleCompleteTask, tc.userRepository)).Methods("PUT")
	router.HandleFunc("/tasks/{id}/share", auth.WithJWTAuth(tc.handleShareTask, tc.userRepository)).Methods("POST")
}

// Handler für Löschen der Aufgabe
func (tc *TaskController) handleDeleteTask(w http.ResponseWriter, r *http.Request){
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())

	// Abrufen der ID der zu löschenden Aufgabe
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing task with ID %v", str))
		return
	}

	// Umwandeln zu int
	taskId, err := strconv.Atoi(str)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid taskId"))
		return
	}

	// Abruden der Löschen-Funktion im Repository
	err = tc.taskRepository.DeleteTask(taskId, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully."})
}

// Handler für Bearbeitung der Aufgabe
func (tc *TaskController) handleUpdateTask(w http.ResponseWriter, r *http.Request){
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())

	// Abrufen der ID der zu löschenden Aufgabe
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing task with ID %v", str))
		return
	}

	// Umwandeln zu int
	taskId, err := strconv.Atoi(str)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid taskId"))
		return
	}

	// Marshalling des JSON zu DTO der Aufgabe
	var taskDTO models.TaskDTO
	if err := utils.ParseJSON(r, &taskDTO); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	task := taskDTO.ToModel()
	task.TaskID = taskId

	// Abrufen der Löschen-Funktion im Repository
	err = tc.taskRepository.UpdateTask(task, userID)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

// Handler für Erledigen der Aufgabe
func (tc *TaskController) handleCompleteTask (w http.ResponseWriter, r *http.Request){
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())

	// Abrufen der ID der zu löschenden Aufgabe
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing task with ID %v", str))
		return
	}

	// Umwandeln zu int
	taskId, err := strconv.Atoi(str)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid taskId"))
		return
	}

	// Abrufen der Löschen-Funktion im Repository
	err = tc.taskRepository.CompleteTask(taskId, userID)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task completed successfully."})
}

// Handler für Erstellen der Aufgabe
func (tc *TaskController) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())

	// Marshalling des JSON zu DTO der Aufgabe
	var taskDTO models.TaskDTO
	if err := utils.ParseJSON(r, &taskDTO); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	task := taskDTO.ToModel()

	// Abrufen der Speichern-Funktion im Repository
	err := tc.taskRepository.CreateTask(task, userID)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

// Handler für das Lesen einer bestimmten Aufgabe
func (tc *TaskController) handleGetTaskById(w http.ResponseWriter, r *http.Request) {
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())

	// Abrufen der ID der zu lesenden Aufgabe
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing task with ID %v", str))
		return
	}

	taskId, err := strconv.Atoi(str)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid taskId"))
		return
	}

	// Abrufen der Lesen-Funktion im Repository
	task, err := tc.taskRepository.GetTaskByIDAndByUser(taskId, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

// Handler für das Lesen aller Aufgabe, die dem Benutzer gehören
func (tc *TaskController) handleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())
	fmt.Printf("ID: %v", userID)

	// Abrufen der Lesen-Funktion im Repository
	tasks, err := tc.taskRepository.GetTasksByUser(userID)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, tasks)
}

// Handler für das Teilen der Aufgabe mit anderen Benutzer
func (tc* TaskController) handleShareTask(w http.ResponseWriter, r *http.Request){
	// Abrufen aus dem Context
	userID := auth.GetUserIDFromContext(r.Context())

	// Abrufen der ID der zu teilenden Aufgabe
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing task with ID %v", str))
		return
	}

	taskId, err := strconv.Atoi(str)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid taskId"))
		return
	}

	// Abrufen der Benutzer-ID, mit dem die Aufgabe geteilt wird
	// Sie wird in JSON von der Benutzer.Seite abgegeben
	var requestBody struct {
        SharedUserID string `json:"sharedUserID"`
    }
	
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

	sharesUserID, _ := strconv.Atoi(requestBody.SharedUserID)

	// Abrufen der Teilen-Funktion im Repository
	err = tc.taskRepository.ShareTask(taskId, userID, sharesUserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task shared successfully."})
}