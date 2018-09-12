package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/apsdehal/go-logger"
	"github.com/gorilla/mux"
)

// Task task entity
type Task struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	Assignee    *Assignee `json:"assignee,omitempty"`
}

// Assignee assignee entity
type Assignee struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var tasks []Task

//GetTasks retrieves current tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

//GetTask retrieves a task by id
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, task := range tasks {
		if task.ID == params["id"] {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

//CreateTask creates a task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = params["id"]
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(tasks)
}

//DeleteTask deletes a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(tasks)
	}
}

// Entrypoint
func main() {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	tasks = append(tasks, Task{ID: "1", Title: "Task 1", Description: "First task", Status: "New", Assignee: &Assignee{Name: "Fabian Rivera", Email: "fabian.riveram@gmail.com"}})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Description: "Second task", Status: "New", Assignee: &Assignee{Name: "Fabian Rivera", Email: "fabian.riveram@gmail.com"}})

	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	http.ListenAndServe(":8000", router)
	log.Info("Listening at port 8000")
}
