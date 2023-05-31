package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("customer_id", uuid.UUID{}),
		field.UUID("going_flight_id", uuid.UUID{}),
		field.UUID("return_flight_id", uuid.UUID{}).Optional(),
		field.Enum("status").Values("Cancelled", "Scheduled", "Departed"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ticket", Ticket.Type),
	}
}
