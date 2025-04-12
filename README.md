# 🧠 Go REST API - Simple Task Manager

A minimal REST API built using Go (Golang) to manage tasks.

---

## 🚀 Features

- ✅ Create new tasks
- 📋 List all tasks
- 🗂 Mark tasks as done
- 💾 Save and load tasks from a local `data.json` file (no database!)
- ⚡ Fast and lightweight with native Go HTTP server

---

## 🧪 API Endpoints

### GET `/tasks`
Returns the list of all tasks.

```bash
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"name": "Learn REST API", "done": false}'
```

✅ Go Basics
How to create and run a Go project with go mod init

How to define structs (Task) and manage JSON data

✅ HTTP in Go
Setting up a basic HTTP server using http.ListenAndServe

Using http.HandleFunc to define endpoints

Parsing request bodies using json.Unmarshal

Sending responses with json.NewEncoder

✅ Working with Files
Reading from and writing to a local .json file

Using os.OpenFile, ioutil.ReadFile, and json.MarshalIndent
