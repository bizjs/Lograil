package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// APIKey holds the schema definition for the APIKey entity.
type APIKey struct {
	ent.Schema
}

// Fields of the APIKey.
func (APIKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("API key name for identification"),
		field.String("hashed_key").
			Unique().
			NotEmpty().
			Sensitive().
			Comment("Hashed API key for authentication"),
		field.String("permissions").
			Default("read").
			Comment("API key permissions: read, write, admin"),
		field.Bool("is_active").
			Default(true).
			Comment("Whether the API key is active"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("When the API key was created"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("When the API key was last updated"),
		field.Time("last_used_at").
			Optional().
			Comment("When the API key was last used"),
		field.Time("expires_at").
			Optional().
			Comment("When the API key expires"),
	}
}

// Edges of the APIKey.
func (APIKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("api_keys").
			Unique().
			Required().
			Comment("Project this API key belongs to"),
		edge.From("created_by", User.Type).
			Ref("api_keys").
			Unique().
			Required().
			Comment("User who created this API key"),
	}
}
