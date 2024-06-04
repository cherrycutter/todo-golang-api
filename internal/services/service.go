package services

import (
	"context"
	"fmt"
	"github.com/cherrycutter/todo_app/internal/models"
	"github.com/cherrycutter/todo_app/internal/repos"
)

type TodoServiceImpl struct {
	repo repos.TodoRepository
}

func NewTodoService(repo repos.TodoRepository) TodoService {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) GetTodos(ctx context.Context) ([]models.TodoModel, error) {
	return s.repo.GetAllTodos(ctx)
}

func (s *TodoServiceImpl) GetTodo(ctx context.Context, id int) (models.TodoModel, error) {
	return s.repo.GetTodoById(ctx, id)
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, todo models.TodoModel) (models.TodoModel, error) {
	if err := s.validateTodoInput(todo); err != nil {
		return todo, err
	}
	return s.repo.CreateTodo(ctx, todo)
}

func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, id int, todo models.TodoModel) (models.TodoModel, error) {
	if err := s.validateTodoInput(todo); err != nil {
		return todo, err
	}
	return s.repo.UpdateTodo(ctx, id, todo)
}

func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.DeleteTodoById(ctx, id)
}

func (s *TodoServiceImpl) validateTodoInput(todo models.TodoModel) error {
	if todo.Title == "" {
		return fmt.Errorf("field title cannot be empty")
	}
	if len(todo.Title) > 255 {
		return fmt.Errorf("title cannot be longer than 255 characters")
	}
	return nil
}
