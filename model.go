package main

import (
	"context"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:td"`

	ID     int64  `bun:",pk,autoincrement"`
	Author string `bun:"author"`
	//CreatedAt time.Time
	Text     string `bun:"text"`
	Complete bool   `bun:"complete"`
}

func NewTodo(author string, text string) *Todo {
	return &Todo{
		Author: author,
		//CreatedAt: time.Now(),
		Text:     text,
		Complete: false,
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
	var todos []Todo
	err := ts.db.NewSelect().Model(&todos).Where("author = ?", author).Scan(ctx)
	if err != nil {
		panic(err)
	}
	return todos
}

func (ts *TodoService) InsertTodo(ctx context.Context, todo *Todo) {
	_, err := ts.db.NewInsert().Model(todo).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
