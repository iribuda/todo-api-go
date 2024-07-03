package models

import (
	"time"
)

type Task struct {
	TaskID     string		`json:"taskId"`
	Title      string	`json:"title"`
	Text       string	`json:"text"`
	Deadline   time.Time	`json:"deadline"`
	CategoryID string	`json:"categoryId"`
	Done		bool `json:"done"`
}

// type Store struct{
// 	db *sql.DB
// }

// func NewStore(db *sql.DB) *Store{
// 	return &Store{db: db}
// }

// func (s *Store) GetAllTasksFromUser(id int64)([]Task, error){
// 	var tasks []Task
// 	rows, err := s.db.Query("SELECT * FROM task_user WHERE userId = ?", id)
// 	if (err != nil){
// 		return nil, fmt.Errorf("tasks from user %d: %v", id, err)
// 	}
// 	defer rows.Close()
// 	for rows.Next(){
// 		var task Task
// 		if err := rows.Scan(
// 			&task.TaskID,
// 			&task.Title,
// 			&task.Text,
// 			&task.Deadline,
// 			&task.CategoryID,
// 		); err != nil {
// 			return nil, err 
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	if err := rows.Err(); err != nil {
//         return nil, fmt.Errorf("tasksByUser %q: %v", id, err)
//     }
// 	return tasks, nil
// }

// func scanRowIntoTask(row *sql.Rows) (Task, error){
// 	var task Task
// 	err := row.Scan(
// 		&task.TaskID,
// 		&task.Title,
// 		&task.Text,
// 		&task.Deadline,
// 		&task.CategoryID,
// 	)
// 	if err != nil{
// 		return nil, err
// 	}
// 	return task, nil
// }

type User struct {
	UserID   string		`json:"userId"`
	Username string		`json:"username"`
	Email    string		`json:"email"`
	Password string		`json:"password"`
}

type Category struct{
	CategoryID	string		`json:"categoryId"`
	Name		string	`json:"name"`
}

// Repository-Schnittstelle 
type TaskRepository interface{
	GetTasks() ([]*Task, error)
	GetTaskByID(id int)(*Task, error)
	GetTasksByUser(id int)([]*Task, error)
	UpdateTask(Task) error
	CreateTask(Task) error
	DeleteTask(Task) error
}