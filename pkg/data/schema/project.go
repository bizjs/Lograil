package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Project name"),
		field.String("description").
			Optional().
			Comment("Project description"),
		field.String("status").
			Default("active").
			Comment("Project status: active, inactive, archived"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("When the project was created"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("When the project was last updated"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("projects").
			Unique().
			Required().
			Comment("User who owns this project"),
		edge.To("api_keys", APIKey.Type).
			Comment("API keys associated with this project"),
		edge.To("retention_policies", RetentionPolicy.Type).
			Comment("Retention policies for this project"),
	}
}
