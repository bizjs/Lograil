package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RetentionPolicy holds the schema definition for the RetentionPolicy entity.
type RetentionPolicy struct {
	ent.Schema
}

// Fields of the RetentionPolicy.
func (RetentionPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Retention policy name"),
		field.Int("duration_days").
			Positive().
			Default(30).
			Comment("Number of days to retain logs"),
		field.String("description").
			Optional().
			Comment("Description of the retention policy"),
		field.Bool("is_active").
			Default(true).
			Comment("Whether the retention policy is active"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("When the retention policy was created"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("When the retention policy was last updated"),
	}
}

// Edges of the RetentionPolicy.
func (RetentionPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("retention_policies").
			Unique().
			Required().
			Comment("Project this retention policy belongs to"),
	}
}
