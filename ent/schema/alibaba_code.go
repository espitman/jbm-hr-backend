package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AlibabaCode holds the schema definition for the AlibabaCode entity.
type AlibabaCode struct {
	ent.Schema
}

// Fields of the AlibabaCode.
func (AlibabaCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			Unique().
			NotEmpty().
			Comment("The unique code for Alibaba access"),
		field.Bool("used").
			Default(false).
			Comment("Whether the code has been used"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("When the code was created"),
		field.Int("assign_to_user_id").
			Optional().
			Nillable().
			Comment("The ID of the user this code is assigned to"),
		field.Time("assign_at").
			Optional().
			Nillable().
			Comment("When the code was assigned to a user"),
		field.Enum("type").
			Values("1m", "3m", "6m", "12m", "25m").
			Comment("The type of the code (price of the code)"),
	}
}

// Edges of the AlibabaCode.
func (AlibabaCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assigned_to_user", User.Type).
			Unique().
			Field("assign_to_user_id"),
	}
}
