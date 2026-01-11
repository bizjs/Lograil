package schema

import (
	"time"

	"entgo.io/ent"
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
		field.String("username").
			Unique().
			NotEmpty().
			Comment("Unique username for the user"),
		field.String("email").
			Unique().
			NotEmpty().
			Comment("Email address of the user"),
		field.String("password_hash").
			NotEmpty().
			Sensitive().
			Comment("Hashed password for authentication"),
		field.String("role").
			Default("user").
			Comment("User role: admin, user, etc."),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("When the user was created"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("When the user was last updated"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type).
			Comment("Projects owned by this user"),
		edge.To("api_keys", APIKey.Type).
			Comment("API keys created by this user"),
	}
}
