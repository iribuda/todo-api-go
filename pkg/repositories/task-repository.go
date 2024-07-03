package repositories

import (
	"database/sql"
	"guthub.com/iribuda/todo-api-go/pkg/models"
)

// Implementierung der TaskRepository-Schnittstelle
type TaskRepositoryImpl struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *TaskRepositoryImpl{
	return &TaskRepositoryImpl{db: db}
}

func (tr *TaskRepositoryImpl) GetTasks() ([]*models.Task, error){
	return nil, nil
}
func (tr *TaskRepositoryImpl) GetTaskByID(id int)(*models.Task, error){
	return nil, nil
}
func (tr *TaskRepositoryImpl) GetTasksByUser(id int)([]*models.Task, error){
	return nil, nil
}
func (tr *TaskRepositoryImpl) UpdateTask(models.Task) error{
	return nil
}
func (tr *TaskRepositoryImpl) CreateTask(models.Task) error{
	return nil
}
func (tr *TaskRepositoryImpl) DeleteTask(models.Task) error{
	return nil
}