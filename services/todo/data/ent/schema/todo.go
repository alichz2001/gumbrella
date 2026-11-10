package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/alichz2001/gombrella/services/todo/biz"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	baseModel
}

// Annotations gives the resource a singular table name within the shared schema.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "todo"}}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default(""),
		field.String("content").Default(""),
		field.Bool("completed").Default(false),
		// status marks the row lifecycle: deletes flip it to deleted instead of
		// removing the row. PostgreSQL stores the canonical proto names in a
		// native enum type, created by the SQL migrations before this table.
		field.Enum("status").
			GoType(biz.TodoStatus("")).
			Default(biz.TodoStatusActive.String()).
			SchemaType(map[string]string{
				dialect.Postgres: postgresSchema + ".todo_status",
			}),
	}
}

// Indexes of the Todo.
func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
	}
}
