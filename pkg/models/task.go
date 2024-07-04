package models

import (
	"strconv"
	"time"
)

// DTO ist notwendig, die Daten zu formatieren und vollständige Information für Frontend zu liefern
type TaskDTO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Deadline   string `json:"deadline"`
	CategoryID string `json:"categoryId"`
	Category   string `json:"category"`
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

// Umwandlung von TaskDTO zu Task-Model
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
