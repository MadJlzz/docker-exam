package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TodoRequest struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type TodoResponse struct {
}

type TodoController struct {
	ts *TodoService
}

func NewTodoController(ts *TodoService) *TodoController {
	return &TodoController{ts: ts}
}

func (tc *TodoController) ListTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := tc.ts.ListTodos(r.Context())
		fmt.Fprintln(w, todos)
	}
}

func (tc *TodoController) InsertTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tr TodoRequest

		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		t := NewTodo(tr.Author, tr.Text)
		tc.ts.InsertTodo(r.Context(), t)
	}
}
