package models

import "time"

type Task struct {
	TaskID     int
	Title      string
	Text       string
	Deadline   time.Time
	CategoryID int
}

type User struct {
	UserID   int
	Username string
	Email    string
	Password string
}

type Category struct{
	CategoryID	int
	Name		string
}