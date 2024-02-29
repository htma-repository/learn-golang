package repository

import (
	"context"
	"learn-golang-todolist/model"
)

type TodoRepository interface {
	InsertTodo(ctx context.Context, todo model.Todo) (model.Todo, error)
	FindTodoById(ctx context.Context, id int32) (model.Todo, error)
	FindAllTodo(ctx context.Context) ([]model.Todo, error)
	UpdateTodo(ctx context.Context, todo model.Todo) (model.Todo, error)
	DeleteTodo(ctx context.Context, id int32) error
}
