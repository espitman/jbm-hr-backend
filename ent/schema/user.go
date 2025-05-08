package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("phone").
			NotEmpty(),
		field.String("first_name").
			NotEmpty(),
		field.String("last_name").
			NotEmpty(),
		field.String("full_name").
			NotEmpty(),
		field.Enum("role").
			Values("admin", "employee").
			Default("employee"),
		field.String("avatar").
			Optional(),
		field.String("password").
			Optional().
			Sensitive(),
		field.Time("birthdate").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Optional(),
		field.Time("cooperation_start_date").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("otps", OTP.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("resumes", Resume.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("requests", Request.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("department", Department.Type).
			Ref("users").
			Unique(),
	}
}
