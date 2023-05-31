// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flookybooky/ent/booking"
	"flookybooky/ent/ticket"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookingCreate is the builder for creating a Booking entity.
type BookingCreate struct {
	config
	mutation *BookingMutation
	hooks    []Hook
}

// SetCustomerID sets the "customer_id" field.
func (bc *BookingCreate) SetCustomerID(u uuid.UUID) *BookingCreate {
	bc.mutation.SetCustomerID(u)
	return bc
}

// SetGoingFlightID sets the "going_flight_id" field.
func (bc *BookingCreate) SetGoingFlightID(u uuid.UUID) *BookingCreate {
	bc.mutation.SetGoingFlightID(u)
	return bc
}

// SetReturnFlightID sets the "return_flight_id" field.
func (bc *BookingCreate) SetReturnFlightID(u uuid.UUID) *BookingCreate {
	bc.mutation.SetReturnFlightID(u)
	return bc
}

// SetNillableReturnFlightID sets the "return_flight_id" field if the given value is not nil.
func (bc *BookingCreate) SetNillableReturnFlightID(u *uuid.UUID) *BookingCreate {
	if u != nil {
		bc.SetReturnFlightID(*u)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BookingCreate) SetStatus(b booking.Status) *BookingCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BookingCreate) SetCreatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BookingCreate) SetNillableCreatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BookingCreate) SetID(u uuid.UUID) *BookingCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BookingCreate) SetNillableID(u *uuid.UUID) *BookingCreate {
	if u != nil {
		bc.SetID(*u)
	}
	return bc
}

// AddTicketIDs adds the "ticket" edge to the Ticket entity by IDs.
func (bc *BookingCreate) AddTicketIDs(ids ...uuid.UUID) *BookingCreate {
	bc.mutation.AddTicketIDs(ids...)
	return bc
}

// AddTicket adds the "ticket" edges to the Ticket entity.
func (bc *BookingCreate) AddTicket(t ...*Ticket) *BookingCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTicketIDs(ids...)
}

// Mutation returns the BookingMutation object of the builder.
func (bc *BookingCreate) Mutation() *BookingMutation {
	return bc.mutation
}

// Save creates the Booking in the database.
func (bc *BookingCreate) Save(ctx context.Context) (*Booking, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookingCreate) SaveX(ctx context.Context) *Booking {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookingCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookingCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookingCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := booking.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := booking.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookingCreate) check() error {
	if _, ok := bc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`ent: missing required field "Booking.customer_id"`)}
	}
	if _, ok := bc.mutation.GoingFlightID(); !ok {
		return &ValidationError{Name: "going_flight_id", err: errors.New(`ent: missing required field "Booking.going_flight_id"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Booking.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Booking.created_at"`)}
	}
	return nil
}

func (bc *BookingCreate) sqlSave(ctx context.Context) (*Booking, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookingCreate) createSpec() (*Booking, *sqlgraph.CreateSpec) {
	var (
		_node = &Booking{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(booking.Table, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeUUID))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.CustomerID(); ok {
		_spec.SetField(booking.FieldCustomerID, field.TypeUUID, value)
		_node.CustomerID = value
	}
	if value, ok := bc.mutation.GoingFlightID(); ok {
		_spec.SetField(booking.FieldGoingFlightID, field.TypeUUID, value)
		_node.GoingFlightID = value
	}
	if value, ok := bc.mutation.ReturnFlightID(); ok {
		_spec.SetField(booking.FieldReturnFlightID, field.TypeUUID, value)
		_node.ReturnFlightID = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := bc.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.TicketTable,
			Columns: []string{booking.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookingCreateBulk is the builder for creating many Booking entities in bulk.
type BookingCreateBulk struct {
	config
	builders []*BookingCreate
}

// Save creates the Booking entities in the database.
func (bcb *BookingCreateBulk) Save(ctx context.Context) ([]*Booking, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Booking, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookingCreateBulk) SaveX(ctx context.Context) []*Booking {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookingCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookingCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
