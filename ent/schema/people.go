package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// People holds the schema definition for the People entity.
type People struct {
	ent.Schema
}

// Fields of the People.
func (People) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("last_name"),
		field.Int("age").Positive(),
	}
}

// Edges of the People.
func (People) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clothes", Clothe.Type),
		edge.From("kind", Group.Type).Ref("peoples"),
	}
}
