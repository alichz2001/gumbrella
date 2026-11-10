package ent_schema

import (
	"slices"
	"testing"

	"entgo.io/ent"
)

func TestBaseModelInheritedMixins(t *testing.T) {
	var model ent.Interface = struct{ BaseModel }{}
	mixins := model.Mixin()
	if len(mixins) != 2 {
		t.Fatalf("Mixin() returned %d mixins, want 2", len(mixins))
	}
	if _, ok := mixins[0].(IDMixin); !ok {
		t.Fatalf("first mixin = %T, want IDMixin", mixins[0])
	}
	if _, ok := mixins[1].(TimeMixin); !ok {
		t.Fatalf("second mixin = %T, want TimeMixin", mixins[1])
	}

	var fields, indexes []string
	for _, mixin := range mixins {
		for _, field := range mixin.Fields() {
			desc := field.Descriptor()
			if desc.Err != nil {
				t.Fatalf("field %q: %v", desc.Name, desc.Err)
			}
			fields = append(fields, desc.Name)
		}
		for _, index := range mixin.Indexes() {
			indexes = append(indexes, index.Descriptor().Fields...)
		}
	}
	if want := []string{"id", "created_at", "updated_at", "deleted_at"}; !slices.Equal(fields, want) {
		t.Errorf("inherited fields = %v, want %v", fields, want)
	}
	if want := []string{"updated_at", "deleted_at"}; !slices.Equal(indexes, want) {
		t.Errorf("inherited indexes = %v, want %v", indexes, want)
	}
}
