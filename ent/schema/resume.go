package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Resume holds the schema definition for the Resume entity.
type Resume struct {
	ent.Schema
}

// Fields of the Resume.
func (Resume) Fields() []ent.Field {
	return []ent.Field{
		field.String("introduced_name").
			NotEmpty().
			MaxLen(100),
		field.String("introduced_phone").
			NotEmpty().
			MaxLen(20),
		field.String("position").
			NotEmpty().
			MaxLen(100),
		field.String("file").
			NotEmpty(),
		field.Enum("status").
			Values("pending", "reviewed", "accepted", "rejected").
			Default("pending"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("user_id"),
	}
}

// Edges of the Resume.
func (Resume) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("resumes").
			Field("user_id").
			Required().
			Unique(),
	}
}
