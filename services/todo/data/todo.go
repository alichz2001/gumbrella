package data

import (
	"context"

	"github.com/alichz2001/gombrella/services/todo/biz"
	ent2 "github.com/alichz2001/gombrella/services/todo/data/ent"
	todo2 "github.com/alichz2001/gombrella/services/todo/data/ent/todo"
	errors2 "github.com/alichz2001/gombrella/services/todo/errors"
	"github.com/alichz2001/gombrella/services/todo/model"

	"github.com/go-kratos/aip-go/ents"
	"github.com/google/uuid"
)

// toBiz converts a persisted todo into its domain representation. The status
// column is bound to model.TodoStatus, so it carries over as-is.
func toBiz(po *ent2.Todo) *model.Todo {
	if po == nil {
		return nil
	}
	return &model.Todo{
		ID:        po.ID,
		Title:     po.Title,
		Content:   po.Content,
		Completed: po.Completed,
		Status:    po.Status,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}
}

type todoRepo struct {
	data *Data
}

// NewTodoRepo creates a new TodoRepo instance.
func NewTodoRepo(data *Data) biz.TodoRepo {
	return &todoRepo{data: data}
}

func (r *todoRepo) FindByID(ctx context.Context, id uuid.UUID) (*model.Todo, error) {
	po, err := r.data.db.Todo.Query().
		Where(todo2.IDEQ(id), todo2.StatusEQ(model.TodoStatusActive)).
		Only(ctx)
	if err != nil {
		if ent2.IsNotFound(err) {
			return nil, errors2.ErrTodoNotFound
		}
		return nil, err
	}
	return toBiz(po), nil
}

func (r *todoRepo) ListTodos(ctx context.Context, opts ...biz.ListOption) ([]*model.Todo, error) {
	options := biz.ListOptions{Limit: 20}
	for _, opt := range opts {
		opt(&options)
	}
	if options.Offset < 0 || options.Limit <= 0 {
		return nil, errors2.ErrTodoInvalidArgument
	}
	// Offset pagination needs a total order, so id is always appended as the
	// last sort key: without it rows that tie on the requested field may be
	// returned in a different order per query, duplicating or skipping records
	// across pages. UUIDv7 ids are time-ordered, so id alone is a sensible
	// default when the caller asks for nothing.
	pos, err := r.data.db.Todo.Query().
		Where(todo2.StatusEQ(model.TodoStatusActive)).
		Where(ents.ApplyFilter(options.Filter)).
		Order(ents.ApplyOrderBy(options.OrderBy), todo2.ByID()).
		Offset(options.Offset).
		Limit(options.Limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	todos := make([]*model.Todo, 0, len(pos))
	for _, po := range pos {
		todos = append(todos, toBiz(po))
	}
	return todos, nil
}

func (r *todoRepo) CreateTodo(ctx context.Context, t *model.Todo) (*model.Todo, error) {
	po, err := r.data.db.Todo.Create().
		SetTitle(t.Title).
		SetContent(t.Content).
		SetCompleted(t.Completed).
		SetStatus(model.TodoStatusActive).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toBiz(po), nil
}

func (r *todoRepo) UpdateTodo(ctx context.Context, t *model.Todo) (*model.Todo, error) {
	// status is not updatable here: it moves only through DeleteTodo.
	po, err := r.data.db.Todo.UpdateOneID(t.ID).
		Where(todo2.StatusEQ(model.TodoStatusActive)).
		SetTitle(t.Title).
		SetContent(t.Content).
		SetCompleted(t.Completed).
		Save(ctx)
	if err != nil {
		if ent2.IsNotFound(err) {
			return nil, errors2.ErrTodoNotFound
		}
		return nil, err
	}
	return toBiz(po), nil
}

// DeleteTodo soft-deletes a todo by flipping its status to deleted. The row is
// kept so it stays auditable, and every read path filters it out.
func (r *todoRepo) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	affected, err := r.data.db.Todo.Update().
		Where(todo2.IDEQ(id), todo2.StatusEQ(model.TodoStatusActive)).
		SetStatus(model.TodoStatusDeleted).
		Save(ctx)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors2.ErrTodoNotFound
	}
	return nil
}
