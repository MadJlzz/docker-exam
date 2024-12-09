package main

import (
	"net/http"
)

func main() {
	cfg := NewAppConfiguration()

	sl := NewLogger(cfg.Logger)
	defer sl.Sync()

	db := NewDatabase(cfg.Database)
	sl.Infof("dsn for database is [%s]", cfg.Database.SafeDsn())
	if err := db.Ping(); err != nil {
		sl.Fatalw("unable to connect to database", "error", err)
	}

	ts := NewTodoService(db, sl)

	sl.Infof("preparing todos rest controller")
	tc := NewTodoController(ts)

	s := NewServer(sl, tc)

	sl.Infof("starting server on [%s]", cfg.Server.Hostname)
	if err := http.ListenAndServe(cfg.Server.Hostname, s); err != nil {
		sl.Fatalln(err)
	}
}
