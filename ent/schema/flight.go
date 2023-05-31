package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.UUID("origin_id", uuid.UUID{}),
		field.UUID("destinartion_id", uuid.UUID{}),
		field.Time("departure_time"),
		field.Time("arrival_time"),
		field.Int("total_slots"),
		field.Int("available_slots"),
		field.Enum("status").Values("Cancelled", "Departed", "Landed", "Scheduled", "Delayed"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("origin", Airport.Type).Ref("origin").Unique().Field("origin_id").Required(),
		edge.From("destination", Airport.Type).Ref("destination").Unique().Field("destinartion_id").Required(),
	}
}
