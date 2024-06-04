package repos

import (
	"context"
	"github.com/cherrycutter/todo_app/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type TodoRepository interface {
	GetAllTodos(ctx context.Context) ([]models.TodoModel, error)
	GetTodoById(ctx context.Context, id int) (models.TodoModel, error)
	CreateTodo(ctx context.Context, todo models.TodoModel) (models.TodoModel, error)
	UpdateTodo(ctx context.Context, id int, todo models.TodoModel) (models.TodoModel, error)
	DeleteTodoById(ctx context.Context, id int) error
}

type PgxConnIface interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}
