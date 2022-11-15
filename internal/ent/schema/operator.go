package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Operator holds the schema definition for the Operator entity.
type Operator struct {
	ent.Schema
}

// Fields of the Operator.
func (Operator) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("email").Optional().Unique(),
		field.String("password"),
		field.String("nickname").Optional(),
		field.String("phone").Unique(),
		field.String("avatar").Optional(),
		field.String("remember_token").Optional(),
	}
}

// Edges of the Operator.
func (Operator) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Operator) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "operators"},
	}
}

func (Operator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
