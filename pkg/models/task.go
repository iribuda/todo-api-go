package models

import (
	"strconv"
	"time"
)

// DTO ist notwendig, die Daten zu formatieren
type TaskDTO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Deadline   string `json:"deadline"`
	CategoryID string `json:"categoryId"`
	Category 	string	`json:"category"`
	Done       bool   `json:"done"`
}

type Task struct {
	TaskID     int
	Title      string
	Text       string    
	Deadline   time.Time 
	CategoryID int       
	Done       bool     
}

// func (task *Task) ToDto() TaskDTO {
// 	return TaskDTO{
// 		// ID:         strconv.Itoa(task.TaskID),
// 		Title:      task.Title,
// 		Text:       task.Text,
// 		Deadline:   task.Deadline.Format("2006-01-02"),
// 		CategoryID: strconv.Itoa(task.CategoryID),
// 		Category: task.c,
// 		Done:       task.Done,
// 	}
// }

func (t *TaskDTO) ToModel() *Task {
	deadline, _ := time.Parse("2006-01-02", t.Deadline)
	categoryId, _ := strconv.Atoi(t.CategoryID)

	return &Task{
		Title:      t.Title,
		Text:       t.Text,
		Deadline:   deadline,
		CategoryID: categoryId,
		Done:       t.Done,
	}
}

// Repository-Schnittstelle
type TaskRepository interface {
	GetTaskByIDAndByUser(taskId, userId int) (*TaskDTO, error)
	GetTasksByUser(id int) ([]*TaskDTO, error)
	UpdateTask(*Task, int) error
	CreateTask(*Task, int) error
	ShareTask(int, int, int) error
	DeleteTask(taskId int, userId int) error
}
