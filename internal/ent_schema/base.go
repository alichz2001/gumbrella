package ent_schema

import "entgo.io/ent"

// BaseModel supplies the standard ID and audit timestamps to Ent schemas.
// Embed it instead of ent.Schema to inherit both mixins.
type BaseModel struct {
	ent.Schema
}

// Mixin returns the shared ID and time mixins.
func (BaseModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}
