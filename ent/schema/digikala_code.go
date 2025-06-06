package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DigikalaCode holds the schema definition for the DigikalaCode entity.
type DigikalaCode struct {
	ent.Schema
}

// Fields of the DigikalaCode.
func (DigikalaCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			Unique().
			NotEmpty(),
		field.Bool("used").
			Default(false),
		field.Time("created_at"),
		field.Int("assign_to_user_id").
			Optional(),
		field.Time("assign_at").
			Optional(),
	}
}

// Edges of the DigikalaCode.
func (DigikalaCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assigned_to", User.Type).
			Ref("digikala_codes").
			Field("assign_to_user_id").
			Unique(),
	}
}
