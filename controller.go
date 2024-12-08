package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TodoRequest struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type TodoResponse struct {
}

func ListAuthorTodos(todoSvc *TodoService) http.HandlerFunc {
	// probably doing something here to pass the db manager
	return func(w http.ResponseWriter, r *http.Request) {
		author := chi.URLParam(r, "author")
		todos := todoSvc.ListTodosFrom(r.Context(), author)
		fmt.Fprintln(w, todos)
	}
}

func InsertTodo(todoSvc *TodoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tr TodoRequest

		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		t := NewTodo(tr.Author, tr.Text)
		todoSvc.InsertTodo(r.Context(), t)
	}
}
