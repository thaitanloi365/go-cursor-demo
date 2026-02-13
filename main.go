package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Cursor Context"},
	{ID: 2, Title: "Ship AI Feature"},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func main() {
	http.HandleFunc("/todos", getTodos)
	println("Server running on :8888")
	http.ListenAndServe(":8888", nil)
}
