package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Clothe holds the schema definition for the Clothe entity.
type Clothe struct {
	ent.Schema
}

// Fields of the Clothe.
func (Clothe) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("color"),
		field.Time("buy_date"),
	}
}

// Edges of the Clothe.
func (Clothe) Edges() []ent.Edge {
	return []ent.Edge{
		// Unique ensures the registry is linked to only registry in peoples
		edge.From("owner", People.Type).Ref("clothes").Unique(),
	}
}
