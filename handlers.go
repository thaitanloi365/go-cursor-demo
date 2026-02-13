package main

import (
	"encoding/json"
	"net/http"
)

var todos = []Todo{
	{ID: 1, Title: "Learn Cursor Context"},
	{ID: 2, Title: "Ship AI Feature"},
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		postTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var req Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	nextID := 1
	for _, t := range todos {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
	newTodo := Todo{ID: nextID, Title: req.Title}
	todos = append(todos, newTodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}
