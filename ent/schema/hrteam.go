package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HRTeam holds the schema definition for the HRTeam entity.
type HRTeam struct {
	ent.Schema
}

// Fields of the HRTeam.
func (HRTeam) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").
			NotEmpty().
			Comment("Full name of the HR team member"),
		field.String("role").
			NotEmpty().
			Comment("Role of the HR team member"),
		field.String("email").
			NotEmpty().
			Comment("Email address of the HR team member"),
		field.String("phone").
			NotEmpty().
			Comment("Phone number of the HR team member"),
	}
}

// Edges of the HRTeam.
func (HRTeam) Edges() []ent.Edge {
	return nil
}
