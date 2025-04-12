package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// create object struct
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// create data file
const dataFile = "tasks.json"

// http get
func getTasks(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		// validate file opened succesfully
		http.Error(w, "Unable to read data", http.StatusInternalServerError)
		return
	}

	// set request headers
	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Read existing tasks
	file, err := os.ReadFile(dataFile)
	if err != nil {
		http.Error(w, "Unable to read data", http.StatusInternalServerError)
		return
	}

	var tasks []Task
	json.Unmarshal(file, &tasks)

	// generate new unique ID
	// Simply adding an extra int to length of data
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	// format updated data
	updatedData, _ := json.MarshalIndent(tasks, "", " ")
	ioutil.WriteFile(dataFile, updatedData, 0644)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// entry point
func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		// check end point type
		if r.Method == http.MethodGet {
			getTasks(w, r)
		} else if r.Method == http.MethodPost {
			postTask(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// logging
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
