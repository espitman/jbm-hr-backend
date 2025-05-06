package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("full_name").
			NotEmpty(),
		field.Enum("kind").
			Values(
				"employment",           // اشتغال به کار
				"payroll_stamped",      // فیش حقوقی مهر شده
				"salary_deduction",     // کسر از حقوق
				"introduction_letter",  // معرفی نامه
				"good_conduct_letter",  // نامه حسن انجام کار
				"confirmation_letter",  // نامه تاییدیه
				"embassy_letter",       // نامه سفارت
				"development_learning", // توسعه و یادگیری
				"marriage_gift",        // هدیه ازدواج
				"childbirth_gift",      // هدیه تولد فرزند
			),
		field.String("description").
			Optional(),
		field.Enum("status").
			Values("pending", "doing", "done", "rejected").
			Default("pending"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("requests").
			Field("user_id").
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("meta", RequestMeta.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
