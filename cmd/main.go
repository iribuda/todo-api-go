package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"guthub.com/iribuda/todo-api-go/pkg/models"
)

var Tasks []models.Task

// method for getting all the tasks
func getAllTasks(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllTasks")
	// encode tasks array to JSON string, write as part of response	
	json.NewEncoder(w).Encode(Tasks)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the ToDo List!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
	// add tasks to route and map to getAllTasks()-function
	http.HandleFunc("/tasks", getAllTasks)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Tasks = []models.Task{
		models.Task{TaskID: 1, Title: "Create DB", Text: "MySQL DB to be created", Deadline: time.Now(), CategoryID: 1},
	}
    handleRequests()	
}