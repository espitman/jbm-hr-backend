package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			NotEmpty(),
		field.String("caption").
			Optional(),
		field.Int("display_order").
			Default(0).
			Comment("Order in which the album should be displayed"),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return []ent.Edge{
		// Add edges here if needed
	}
}
