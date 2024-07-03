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

// Konstruktor
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
	rows, err := tr.db.Query("SELECT * FROM task WHERE taskId = ?", id)
	if err != nil {
		return nil, err
	}

	t := new(models.Task)
	defer rows.Close()

	for rows.Next(){
		t, err = scanRowsIntoTask(rows)
		if err != nil {
			return nil, err 
		}
	}

	if err == sql.ErrNoRows {
		return t, fmt.Errorf("taskById %d: no such task", id)
	} else if err != nil{
        return t, fmt.Errorf("taskById %d: %v", id, err)
	}

	return t, nil
}
func (tr *TaskRepositoryImpl) GetTasksByUser(id int)([]*models.Task, error){
	return nil, nil
}

func (tr *TaskRepositoryImpl) UpdateTask(task *models.Task) error{
	_, err := tr.db.Exec("UPDATE task SET title = ?, text = ?, deadline = ?, categoryId = ?, done = ? WHERE taskId = ?", 
		task.Title, task.Text, task.Deadline, task.CategoryID, task.Done, task.TaskID)
	if err != nil{
		return err
	}
	return nil
}

func (tr *TaskRepositoryImpl) CreateTask(task *models.Task) error{
	_, err := tr.db.Exec("INSERT INTO task (title, text, deadline, categoryId) VALUES (?, ?, ?, ?)",
		task.Title, task.Text, task.Deadline, task.CategoryID)
	if err != nil{
		return err
	}
	return nil
}

func (tr *TaskRepositoryImpl) DeleteTask(id int) error{
	_, err := tr.db.Exec("DELETE FROM task WHERE taskId = ?", id)
	if err != nil{
		return err
	}
	return nil
}

// Hilf-Funktion für Aufrufen der Aufgaben aus sql-ResultSet
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