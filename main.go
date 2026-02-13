package main

import "net/http"

func main() {
	http.HandleFunc("/todos", todoHandler)
	println("Server running on :8888")
	http.ListenAndServe(":8888", nil)
}
