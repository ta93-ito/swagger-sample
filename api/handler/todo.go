package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ta93-ito/golang-swagger-sample/api/db"
)

type TODO struct {
	ID          int64     `json:"id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type POSTTODORequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type TODOResponse struct {
	*TODO
}

type TODOs []TODO
type TODOsResponse TODOs

const (
	selectAll  = `SELECT * FROM todos`
	selectByID = `SELECT * FROM todos WHERE id = ?`
	insert     = `INSERT INTO todos(subject, description) VALUES(?,?)`
)

func POSTTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method != "POST" {
		http.Error(w, fmt.Errorf("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	var req POSTTODORequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Subject == "" || req.Description == "" {
		http.Error(w, fmt.Errorf("missing parameter subject or description").Error(), http.StatusBadRequest)
		return
	}
	result, err := db.DB.Exec(insert, req.Subject, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row := db.DB.QueryRow(selectByID, i)
	todo := new(TODO)
	if err := row.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := &TODOResponse{TODO: todo}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GETAllTODOs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method != "GET" {
		http.Error(w, fmt.Errorf("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.DB.Query(selectAll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var todos []TODO
	for rows.Next() {
		todo := new(TODO)
		if err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, *todo)
	}
	var res TODOsResponse = todos
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
