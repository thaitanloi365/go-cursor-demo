package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostTodo(t *testing.T) {
	originalTodos := make([]Todo, len(todos))
	copy(originalTodos, todos)
	t.Cleanup(func() {
		todos = append([]Todo(nil), originalTodos...)
	})

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "invalid JSON",
			body:       "invalid",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid JSON\n",
		},
		{
			name:       "empty title",
			body:       `{"title":""}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "Title is required\n",
		},
		{
			name:       "missing title",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "Title is required\n",
		},
		{
			name:       "valid todo",
			body:       `{"title":"New task"}`,
			wantStatus: http.StatusCreated,
			wantBody:   `{"id":3,"title":"New task"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			postTodo(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if body := rec.Body.String(); body != tt.wantBody {
				t.Errorf("body = %q, want %q", body, tt.wantBody)
			}
		})
	}
}
