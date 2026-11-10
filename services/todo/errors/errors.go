package errors

import (
	v1 "github.com/alichz2001/gombrella/services/todo/api/v1"
	"github.com/go-kratos/kratos/v3/errors"
)

var (
	// ErrTodoNotFound is returned when a todo does not exist.
	ErrTodoNotFound = errors.NotFound(v1.ErrorReason_TODO_NOT_FOUND.String(), "todo not found")
	// ErrTodoInvalidArgument is returned when a todo request is invalid.
	ErrTodoInvalidArgument = errors.BadRequest(v1.ErrorReason_TODO_INVALID_ARGUMENT.String(), "invalid todo argument")
)
