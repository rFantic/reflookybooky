// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"flookybooky/ent/predicate"
	"flookybooky/ent/ticket"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TicketDelete is the builder for deleting a Ticket entity.
type TicketDelete struct {
	config
	hooks    []Hook
	mutation *TicketMutation
}

// Where appends a list predicates to the TicketDelete builder.
func (td *TicketDelete) Where(ps ...predicate.Ticket) *TicketDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TicketDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TicketDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TicketDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ticket.Table, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TicketDeleteOne is the builder for deleting a single Ticket entity.
type TicketDeleteOne struct {
	td *TicketDelete
}

// Where appends a list predicates to the TicketDelete builder.
func (tdo *TicketDeleteOne) Where(ps ...predicate.Ticket) *TicketDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TicketDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ticket.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TicketDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}
