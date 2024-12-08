package main

import (
	"context"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"time"
)

type Todo struct {
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	Complete  bool      `json:"complete"`
}

func NewTodo(author string, text string) *Todo {
	return &Todo{
		Author:    author,
		CreatedAt: time.Now(),
		Text:      text,
		Complete:  false,
	}
}

type TodoService struct {
	db     *bun.DB
	logger *zap.SugaredLogger
}

func NewTodoService(db *bun.DB, logger *zap.SugaredLogger) *TodoService {
	return &TodoService{
		db:     db,
		logger: logger,
	}
}

func (ts *TodoService) ListTodosFrom(ctx context.Context, author string) []Todo {
	//_, err := db.NewInsert().Model(&model).Exec(ctx)
	return nil
}
