package controllers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	router.HandleFunc("/task/{id}", tc.handleGetTaskById)
	router.HandleFunc("/task", tc.handleCreateTask).Methods("POST")
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

func (tc *TaskController) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO models.TaskDTO
	if err := utils.ParseJSON(r, &taskDTO); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	task := taskDTO.ToModel()
	err := tc.repository.CreateTask(task)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, task)
}

func (tc *TaskController) handleGetTaskById(w http.ResponseWriter, r *http.Request) {
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

	task, err := tc.repository.GetTaskByID(taskId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (tc *TaskController) handleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := tc.repository.GetTasks()
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, tasks)
}