// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flookybooky/ent/airport"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Airport is the model entity for the Airport schema.
type Airport struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AirportQuery when eager-loading is set.
	Edges        AirportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AirportEdges holds the relations/edges for other nodes in the graph.
type AirportEdges struct {
	// Origin holds the value of the origin edge.
	Origin []*Flight `json:"origin,omitempty"`
	// Destination holds the value of the destination edge.
	Destination []*Flight `json:"destination,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OriginOrErr returns the Origin value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) OriginOrErr() ([]*Flight, error) {
	if e.loadedTypes[0] {
		return e.Origin, nil
	}
	return nil, &NotLoadedError{edge: "origin"}
}

// DestinationOrErr returns the Destination value or an error if the edge
// was not loaded in eager-loading.
func (e AirportEdges) DestinationOrErr() ([]*Flight, error) {
	if e.loadedTypes[1] {
		return e.Destination, nil
	}
	return nil, &NotLoadedError{edge: "destination"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Airport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case airport.FieldName, airport.FieldAddress:
			values[i] = new(sql.NullString)
		case airport.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Airport fields.
func (a *Airport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case airport.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case airport.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case airport.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				a.Address = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Airport.
// This includes values selected through modifiers, order, etc.
func (a *Airport) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryOrigin queries the "origin" edge of the Airport entity.
func (a *Airport) QueryOrigin() *FlightQuery {
	return NewAirportClient(a.config).QueryOrigin(a)
}

// QueryDestination queries the "destination" edge of the Airport entity.
func (a *Airport) QueryDestination() *FlightQuery {
	return NewAirportClient(a.config).QueryDestination(a)
}

// Update returns a builder for updating this Airport.
// Note that you need to call Airport.Unwrap() before calling this method if this Airport
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Airport) Update() *AirportUpdateOne {
	return NewAirportClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Airport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Airport) Unwrap() *Airport {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Airport is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Airport) String() string {
	var builder strings.Builder
	builder.WriteString("Airport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(a.Address)
	builder.WriteByte(')')
	return builder.String()
}

// Airports is a parsable slice of Airport.
type Airports []*Airport
