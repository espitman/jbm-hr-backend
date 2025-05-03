package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Unique(),
		field.String("description").
			NotEmpty(),
		field.String("image").
			NotEmpty(),
		field.String("icon").
			NotEmpty(),
		field.String("color").
			NotEmpty(),
		field.String("shortName").
			NotEmpty().
			Unique(),
		field.Int("display_order").
			Default(0).
			Comment("Order in which the department should be displayed"),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return nil
}
