package biz

import (
	"context"
	"strings"

	errors2 "github.com/alichz2001/gombrella/services/todo/errors"
	"github.com/alichz2001/gombrella/services/todo/model"

	"github.com/google/uuid"
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
)

// TodoRepo is a todo repo.
type TodoRepo interface {
	FindByID(context.Context, uuid.UUID) (*model.Todo, error)
	ListTodos(context.Context, ...ListOption) ([]*model.Todo, error)
	CreateTodo(context.Context, *model.Todo) (*model.Todo, error)
	UpdateTodo(context.Context, *model.Todo) (*model.Todo, error)
	DeleteTodo(context.Context, uuid.UUID) error
}

// ListOption configures todo list queries.
type ListOption func(*ListOptions)

// ListOptions are todo list query options.
type ListOptions struct {
	Filter  filtering.Filter
	OrderBy ordering.OrderBy
	Offset  int
	Limit   int
}

// ListFilter sets a standard AIP filter.
func ListFilter(filter filtering.Filter) ListOption {
	return func(o *ListOptions) {
		o.Filter = filter
	}
}

// ListOrderBy sets a standard AIP order_by value.
func ListOrderBy(orderBy ordering.OrderBy) ListOption {
	return func(o *ListOptions) {
		o.OrderBy = orderBy
	}
}

// ListOffset sets an offset.
func ListOffset(offset int) ListOption {
	return func(o *ListOptions) {
		o.Offset = offset
	}
}

// ListLimit sets a limit.
func ListLimit(limit int) ListOption {
	return func(o *ListOptions) {
		o.Limit = limit
	}
}

// TodoUsecase is a Todo usecase.
type TodoUsecase struct {
	repo TodoRepo
}

// NewTodoUsecase new a Todo usecase.
func NewTodoUsecase(repo TodoRepo) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

// GetTodo returns a todo by ID.
func (uc *TodoUsecase) GetTodo(ctx context.Context, id uuid.UUID) (*model.Todo, error) {
	if id == uuid.Nil {
		return nil, errors2.ErrTodoInvalidArgument
	}
	return uc.repo.FindByID(ctx, id)
}

// ListTodos lists todos.
func (uc *TodoUsecase) ListTodos(ctx context.Context, opts ...ListOption) ([]*model.Todo, error) {
	return uc.repo.ListTodos(ctx, opts...)
}

// CreateTodo creates a todo.
func (uc *TodoUsecase) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	if err := validateTodo(todo); err != nil {
		return nil, err
	}
	return uc.repo.CreateTodo(ctx, todo)
}

// UpdateTodo updates a todo.
func (uc *TodoUsecase) UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	if todo == nil || todo.ID == uuid.Nil {
		return nil, errors2.ErrTodoInvalidArgument
	}
	if err := validateTodo(todo); err != nil {
		return nil, err
	}
	return uc.repo.UpdateTodo(ctx, todo)
}

// DeleteTodo soft-deletes a todo.
func (uc *TodoUsecase) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors2.ErrTodoInvalidArgument
	}
	return uc.repo.DeleteTodo(ctx, id)
}

func validateTodo(todo *model.Todo) error {
	if todo == nil || strings.TrimSpace(todo.Title) == "" {
		return errors2.ErrTodoInvalidArgument
	}
	return nil
}
