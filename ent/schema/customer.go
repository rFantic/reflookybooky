package schema

import (
	"errors"
	"flookybooky/internal/validate"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("address"),
		field.String("license_id").
			Validate(validate.IsNumeric).
			Validate(func(s string) error {
				if len(s) > 12 {
					return errors.New("license id cannot be longer than 12")
				}
				return nil
			}),
		field.String("phone_number").
			Validate(validate.IsNumeric),
		field.String("email"),
		field.Time("timestamp").Immutable().Default(time.Now()),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
