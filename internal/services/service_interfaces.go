package services

import (
	"context"
	"github.com/cherrycutter/todo_app/internal/models"
)

//go:generate mockgen -source=service_interfaces.go -destination=mocks/mock.go

type TodoService interface {
	GetTodos(ctx context.Context) ([]models.TodoModel, error)
	GetTodo(ctx context.Context, id int) (models.TodoModel, error)
	CreateTodo(ctx context.Context, todo models.TodoModel) (models.TodoModel, error)
	UpdateTodo(ctx context.Context, id int, todo models.TodoModel) (models.TodoModel, error)
	DeleteTodo(ctx context.Context, id int) error
}
