package controllers

import (
	// "encoding/json"
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"guthub.com/iribuda/todo-api-go/pkg/models"
	"guthub.com/iribuda/todo-api-go/pkg/utils"
)

type TaskController struct{
	repository models.TaskRepository
}

func NewController(repository models.TaskRepository) *TaskController{
	return &TaskController{repository: repository}
}

func (tc *TaskController) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/tasks", tc.handleGetAllTasks)
	// router.HandleFunc("/task/{id}", tc.handleGetTaskById)
	// router.HandleFunc("/task", tc.handleCreateTask).Methods("POST")
	// router.HandleFunc("/task/{id}", tc.handleDeleteTask).Methods("DELETE")
}

// func (tc *TaskController) handleDeleteTask(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	for index, task := range Tasks{
// 		if task.TaskID == id{
// 			Tasks = append(Tasks[:index], Tasks[index+1]...)
// 		}
// 	}
// }

// func (tc *TaskController) handleCreateTask(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	var task models.Task
// 	json.Unmarshal([]byte(body), &task)
// 	Tasks = append(Tasks, task)
// 	json.NewEncoder(w).Encode(task)
// }

// func (tc *TaskController) handleGetTaskById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	for _, task := range Tasks {
// 		if task.TaskID == key {
// 			json.NewEncoder(w).Encode(task)
// 		}
// 	}
// }

func (tc *TaskController) handleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := tc.repository.GetTasks()
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, tasks)
}