package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type ServerConfiguration struct {
	Hostname string `yaml:"hostname"`
}

type Server struct {
	logger *zap.SugaredLogger
	mux    *chi.Mux
}

func NewServer(l *zap.SugaredLogger, tc *TodoController) *Server {
	r := chi.NewRouter()

	// Base middlewares to use.
	//r.Use(NewLoggerMiddleware(nil).ZapLogger())

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(requestLoggerMiddleware(l))
		r.Route("/todos", func(r chi.Router) {
			r.Get("/", tc.ListTodos())
			r.Post("/", tc.InsertTodo())
		})
	})

	return &Server{l, r}
}

func requestLoggerMiddleware(l *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			l.Infof("%s %s", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
