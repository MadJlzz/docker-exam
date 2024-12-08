package main

import (
	"net/http"
)

func main() {
	cfg := NewAppConfiguration()

	sl := NewLogger(cfg.Logger)
	defer sl.Sync()

	db := NewDatabase(cfg.Database)
	svc := NewTodoService(db, sl)

	if err := http.ListenAndServe(cfg.Server.Hostname, RegisterRoutes(svc)); err != nil {
		sl.Fatalln(err)
	}
}
