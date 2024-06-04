package repos

import (
	"context"
	"errors"
	"github.com/cherrycutter/todo_app/internal/models"
	"github.com/jackc/pgx/v5"
)

type TodoRepositoryImpl struct {
	db PgxConnIface
}

func NewTodoRepo(db PgxConnIface) TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

var (
	ErrTodoNotFound = errors.New("todo not found")
)

func (r *TodoRepositoryImpl) GetAllTodos(ctx context.Context) ([]models.TodoModel, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.TodoModel
	for rows.Next() {
		var todo models.TodoModel
		if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepositoryImpl) GetTodoById(ctx context.Context, id int) (models.TodoModel, error) {
	var todo models.TodoModel
	err := r.db.QueryRow(ctx, "SELECT * FROM todo WHERE id = $1", id).
		Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.TodoModel{}, ErrTodoNotFound
		}
		return models.TodoModel{}, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) CreateTodo(ctx context.Context, todo models.TodoModel) (models.TodoModel, error) {
	query := `
			INSERT INTO todo (title, description, completed, created_at)
			VALUES ($1, $2, $3, NOW())
			RETURNING id, created_at
		`
	err := r.db.QueryRow(ctx, query, todo.Title, todo.Description, todo.Completed).Scan(&todo.Id, &todo.CreatedAt)
	if err != nil {
		return models.TodoModel{}, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) UpdateTodo(ctx context.Context, id int, todo models.TodoModel) (models.TodoModel, error) {
	query := `
		UPDATE todo
		SET title = $1, description = $2, completed = $3
		WHERE id = $4
		RETURNING id, title, description, completed, created_at
	`
	var updatedTodo models.TodoModel
	err := r.db.QueryRow(
		ctx,
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		id,
	).Scan(
		&updatedTodo.Id,
		&updatedTodo.Title,
		&updatedTodo.Description,
		&updatedTodo.Completed,
		&updatedTodo.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.TodoModel{}, ErrTodoNotFound
		}
		return models.TodoModel{}, err
	}
	return updatedTodo, nil
}

func (r *TodoRepositoryImpl) DeleteTodoById(ctx context.Context, id int) error {
	cmdTag, err := r.db.Exec(ctx, "DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrTodoNotFound
	}
	return nil
}
