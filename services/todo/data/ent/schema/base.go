package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/mixin"

	"github.com/alichz2001/gombrella/internal/ent_schema"
)

const postgresSchema = "todo"

// baseModel scopes all Todo service models to the service's PostgreSQL schema
// while retaining the shared ID and time mixins. Keep it unexported so Ent does
// not discover it as a separate entity.
type baseModel struct {
	ent_schema.BaseModel
}

func (b baseModel) Mixin() []ent.Mixin {
	return append(b.BaseModel.Mixin(), namespaceMixin{})
}

// Use a mixin so per-model table annotations retain the inherited namespace.
type namespaceMixin struct {
	mixin.Schema
}

func (namespaceMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Schema(postgresSchema)}
}
