package repositories

import (
	"database/sql"
	"fmt"
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
	tasks := make([]*models.Task, 0)
	rows, err := tr.db.Query("SELECT * FROM task")
	if (err != nil){
		return nil, fmt.Errorf("all tasks: %v", err)
	}
	defer rows.Close()
	for rows.Next(){
		t, err := scanRowsIntoTask(rows)
		if err != nil {
			return nil, err 
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
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

func scanRowsIntoTask(rows *sql.Rows)(*models.Task, error){
	task := new(models.Task)
	err := rows.Scan(
		&task.TaskID,
		&task.Title,
		&task.Text,
		&task.Deadline,
		&task.CategoryID,
		&task.Done,
	)
	if err != nil{
		return nil, err
	}
	return task, nil
}