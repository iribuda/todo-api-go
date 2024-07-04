package task

import (
	"database/sql"
	"fmt"
	"guthub.com/iribuda/todo-api-go/pkg/models"
)

// Implementierung der TaskRepository-Schnittstelle
// In jeder Funktion wird in SQL-Query durch JOIN sichergestellt, 
// dass Benutzer nur mit seinen Aufgaben arbeitet
type TaskRepositoryImpl struct{
	db *sql.DB
}

// Konstruktor
func NewRepository(db *sql.DB) *TaskRepositoryImpl{
	return &TaskRepositoryImpl{db: db}
}

// Abrufen der Aufgaben des Benutzer
func (tr *TaskRepositoryImpl) GetTasksByUser(userID int) ([]*models.TaskDTO, error){
	// Slice von DTOs für die Aufgaben
	tasks := make([]*models.TaskDTO, 0)

	// SQL Query, die nur die dem Benutzer gehörigen Aufgaben liefert
	rows, err := tr.db.Query("SELECT t.* , c.name as categoryName FROM task t JOIN user_task u ON u.taskId = t.taskId JOIN category c ON t.categoryId = c.categoryId WHERE u.userId = ? ", userID)
	if (err != nil){
		return nil, fmt.Errorf("all tasks by user: %v", err)
	}
	defer rows.Close()

	// Abrufen aus dem ResultSet
	for rows.Next(){
		t, err := scanRowsIntoTask(rows)
		if err != nil {
			return nil, err 
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// Abrufen einer bestimmten Aufgabe des Benutzers
func (tr *TaskRepositoryImpl) GetTaskByIDAndByUser(taskID, userID int)(*models.TaskDTO, error){
	// SQL Query, die dem Benutzer gehörigen Aufgabe liefert
	rows, err := tr.db.Query("SELECT t.*, c.name FROM task t JOIN category c ON t.categoryId = c.categoryId JOIN user_task ut ON t.taskId = ut.taskId WHERE t.taskId = ? AND ut.userId = ? ", taskID, userID)
	if err != nil {
		return nil, err
	}

	// Aufgabe wird als DTO gegeben
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

// Bearbeiten der dem Benutzer gehörigen Aufgabe
func (tr *TaskRepositoryImpl) UpdateTask(task *models.Task, userID int) error{
	// SQL Query, die dem Benutzer gehörigen Aufgabe entspechend ändert
	_, err := tr.db.Exec("UPDATE task t JOIN user_task ut ON t.taskId = ut.taskId SET title = ?, text = ?, deadline = ?, categoryId = ?, done = ? WHERE t.taskId = ? && ut.userID = ? ", 
		task.Title, task.Text, task.Deadline, task.CategoryID, task.Done, task.TaskID, userID)
	if err != nil{
		return err
	}
	return nil
}

// Speichern der Aufgabe
func (tr *TaskRepositoryImpl) CreateTask(task *models.Task, userID int) error{
	// Zuerst wird die Aufgabe selbst gespeichert
	_, err := tr.db.Exec("INSERT INTO `task` (title, text, deadline, categoryId) VALUES (?, ?, ?, ?);",
		task.Title, task.Text, task.Deadline, task.CategoryID)
		if err != nil{
			return err
		}

	// Und danach wird sie in Many-To-Many User-Task-Tabelle gespeichert
	// d.h. wird sie dem aktuellen Benutzer zugewiesen
	_, err = tr.db.Exec("INSERT INTO user_task (userId, taskId) VALUES (?, (SELECT LAST_INSERT_ID())) ", userID)
	if err != nil{
		fmt.Printf("create task %v", err)
		return err
	}
	return nil
}

// Löschen der Aufgabe
func (tr *TaskRepositoryImpl) DeleteTask(taskID int, userID int) error{
	// SQL Query für das Löschen dem Benutzer gehörigen Aufgabe
	_, err := tr.db.Exec("DELETE t FROM task t JOIN user_task ut ON t.taskId = ut.taskId WHERE t.taskId = ? AND ut.userId = ?", taskID, userID)
	if err != nil{
		return err
	}
	return nil
}

// Teilen der Aufgabe
func (tr *TaskRepositoryImpl) ShareTask(taskID int, userID int, sharedUserID int) error{
	// Überprüfen, ob diese Aufgabe dem aktuellen Benutzer gehört
	// Falls nicht, wird das Teilen nicht geführt
	if !tr.taskBelongsToUser(taskID, userID) {
		return nil
	}

	// Sonst wird die Query für Zuweisen der Aufgabe dem anderen Benutzer geführt
	_, err := tr.db.Exec("INSERT INTO user_task (userId, taskId) VALUES (?, ?) ", sharedUserID, taskID)
	if err != nil{
		fmt.Printf("create task %v", err)
		return err
	}
	return nil
}

// Überprüfen, ob die Aufgabe dem gegebenen Benutzer gehört
// Alternativ könnte diese Funktion vor anderen aufgeruft werden, um sicher zu stellen, 
// dass die Benutzer nur mit ihren Aufgaben interagieren können
func (tr *TaskRepositoryImpl) taskBelongsToUser(taskID, userID int) bool{
	rows, err := tr.db.Query("SELECT * FROM user_task WHERE taskId = ? AND userId = ? ", taskID, userID)
	if err != nil {
		return false
	}
	if rows.Next() {
		return true
	}
	return false
}

// Hilf-Funktion für Aufrufen der Aufgaben aus SQL-ResultSet
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