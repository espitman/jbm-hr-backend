package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RequestMeta holds the schema definition for the RequestMeta entity.
type RequestMeta struct {
	ent.Schema
}

// Fields of the RequestMeta.
func (RequestMeta) Fields() []ent.Field {
	return []ent.Field{
		field.Int("request_id"),
		field.String("key").
			NotEmpty(),
		field.String("value").
			NotEmpty(),
	}
}

// Edges of the RequestMeta.
func (RequestMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request", Request.Type).
			Ref("meta").
			Field("request_id").
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
