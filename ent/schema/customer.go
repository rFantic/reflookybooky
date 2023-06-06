package schema

import (
	"errors"
	"flookybooky/internal/util"
	"flookybooky/internal/validate"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		util.TimeMixin{},
	}
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
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
