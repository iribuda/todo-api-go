package task

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

func (tr *TaskRepositoryImpl) GetTasksByUser(userID int) ([]*models.TaskDTO, error){
	tasks := make([]*models.TaskDTO, 0)
	rows, err := tr.db.Query("SELECT t.* , c.name as categoryName FROM task t JOIN user_task u ON u.taskId = t.taskId JOIN category c ON t.categoryId = c.categoryId WHERE u.userId = ? ", userID)
	if (err != nil){
		return nil, fmt.Errorf("all tasks by user: %v", err)
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

func (tr *TaskRepositoryImpl) GetTaskByIDAndByUser(taskID, userID int)(*models.TaskDTO, error){
	rows, err := tr.db.Query("SELECT t.*, c.name FROM task t JOIN category c ON t.categoryId = c.categoryId JOIN user_task ut ON t.taskId = ut.taskId WHERE t.taskId = ? AND ut.userId = ? ", taskID, userID)
	if err != nil {
		return nil, err
	}

	t := new(models.TaskDTO)
	defer rows.Close()

	for rows.Next(){
		t, err = scanRowsIntoTask(rows)
		if err != nil {
			return nil, err 
		}
	}

	if err == sql.ErrNoRows {
		return t, fmt.Errorf("taskById: no such task %v by user %v", taskID, userID)
	} else if err != nil{
        return t, fmt.Errorf("taskById %d: %v", taskID, err)
	}

	return t, nil
}

func (tr *TaskRepositoryImpl) UpdateTask(task *models.Task, userID int) error{
	// _, err := tr.db.Exec("UPDATE task SET title = ?, text = ?, deadline = ?, categoryId = ?, done = ? WHERE taskId = ?", 
	_, err := tr.db.Exec("UPDATE task t JOIN user_task ut ON t.taskId = ut.taskId SET title = ?, text = ?, deadline = ?, categoryId = ?, done = ? WHERE t.taskId = ? && ut.userID = ? ", 
		task.Title, task.Text, task.Deadline, task.CategoryID, task.Done, task.TaskID, userID)
	if err != nil{
		return err
	}
	return nil
}

func (tr *TaskRepositoryImpl) CreateTask(task *models.Task, userID int) error{
	_, err := tr.db.Exec("INSERT INTO `task` (title, text, deadline, categoryId) VALUES (?, ?, ?, ?);",
		task.Title, task.Text, task.Deadline, task.CategoryID)
		if err != nil{
			return err
		}
	_, err = tr.db.Exec("INSERT INTO user_task (userId, taskId) VALUES (?, (SELECT LAST_INSERT_ID())) ", userID)
	if err != nil{
		fmt.Printf("create task %v", err)
		return err
	}
	return nil
}

func (tr *TaskRepositoryImpl) DeleteTask(taskID int, userID int) error{
	_, err := tr.db.Exec("DELETE t FROM task t JOIN user_task ut ON t.taskId = ut.taskId WHERE t.taskId = ? AND ut.userId = ?", taskID, userID)
	if err != nil{
		return err
	}
	return nil
}

// Hilf-Funktion für Aufrufen der Aufgaben aus sql-ResultSet
func scanRowsIntoTask(rows *sql.Rows)(*models.TaskDTO, error){
	task := new(models.TaskDTO)
	err := rows.Scan(
		&task.ID,
		&task.Title,
		&task.Text,
		&task.Deadline,
		&task.CategoryID,
		&task.Done,
		&task.Category,
	)
	if err != nil{
		return nil, err
	}
	return task, nil
}