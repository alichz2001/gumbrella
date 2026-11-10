package model

import (
	"time"

	"github.com/google/uuid"
)

// TodoStatus is the lifecycle state of a todo. Values are the canonical enum
// names shared with protobuf and PostgreSQL; numeric proto values stay at the
// transport boundary.
type TodoStatus string

const (
	// TodoStatusUnspecified represents an unset API status; it is never persisted.
	TodoStatusUnspecified TodoStatus = "TODO_STATUS_UNSPECIFIED"
	// TodoStatusActive marks a live todo.
	TodoStatusActive TodoStatus = "TODO_STATUS_ACTIVE"
	// TodoStatusDeleted marks a soft-deleted todo. The record is retained but
	// hidden from reads.
	TodoStatusDeleted TodoStatus = "TODO_STATUS_DELETED"
)

// String returns the canonical enum name.
func (s TodoStatus) String() string {
	return string(s)
}

// Values returns the allowed persisted values. The unspecified sentinel and
// the empty Go zero value are deliberately excluded.
func (TodoStatus) Values() []string {
	return []string{
		TodoStatusActive.String(),
		TodoStatusDeleted.String(),
	}
}

// Todo is a Todo model.
type Todo struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Completed bool
	Status    TodoStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
