package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OTP holds the schema definition for the OTP entity.
type OTP struct {
	ent.Schema
}

// Fields of the OTP.
func (OTP) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			NotEmpty(),
		field.Time("expires_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("used").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
		field.Time("used_at").
			Optional().
			Nillable(),
	}
}

// Edges of the OTP.
func (OTP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("otps").
			Unique().
			Required(),
	}
}
