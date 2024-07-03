package task

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"guthub.com/iribuda/todo-api-go/pkg/utils"
)

type TaskController struct{
	repository TaskRepository
}

func NewController(repository TaskRepository) *TaskController{
	return &TaskController{repository: repository}
}

func (tc *TaskController) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/tasks", tc.handleGetAllTasks)
	router.HandleFunc("/task", tc.handleCreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", tc.handleGetTaskById).Methods("GET")
	router.HandleFunc("/task/{id}", tc.handleDeleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", tc.handleUpdateTask).Methods("PUT")
}

func (tc *TaskController) handleDeleteTask(w http.ResponseWriter, r *http.Request){
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

	err = tc.repository.DeleteTask(taskId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (tc *TaskController) handleUpdateTask(w http.ResponseWriter, r *http.Request){
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

	var taskDTO TaskDTO
	if err := utils.ParseJSON(r, &taskDTO); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	task := taskDTO.ToModel()
	task.TaskID = taskId
	err = tc.repository.UpdateTask(task)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (tc *TaskController) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
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