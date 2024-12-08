package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ListAuthorTodos(todoSvc *TodoService) http.HandlerFunc {
	// probably doing something here to pass the db manager
	return func(w http.ResponseWriter, r *http.Request) {
		author := chi.URLParam(r, "author")
		todos := todoSvc.ListTodosFrom(r.Context(), author)
		fmt.Fprintln(w, todos)
	}
}
