package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type ServerConfiguration struct {
	Hostname string `yaml:"hostname"`
}

func RegisterRoutes(todoSvc *TodoService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/v1/todos/{author}", ListAuthorTodos(todoSvc))
	r.Post("/api/v1/todo", InsertTodo(todoSvc))

	return r
}
