// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"flookybooky/ent/booking"
	"flookybooky/ent/predicate"
	"flookybooky/ent/ticket"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookingQuery is the builder for querying Booking entities.
type BookingQuery struct {
	config
	ctx        *QueryContext
	order      []booking.OrderOption
	inters     []Interceptor
	predicates []predicate.Booking
	withTicket *TicketQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookingQuery builder.
func (bq *BookingQuery) Where(ps ...predicate.Booking) *BookingQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BookingQuery) Limit(limit int) *BookingQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BookingQuery) Offset(offset int) *BookingQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BookingQuery) Unique(unique bool) *BookingQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BookingQuery) Order(o ...booking.OrderOption) *BookingQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryTicket chains the current query on the "ticket" edge.
func (bq *BookingQuery) QueryTicket() *TicketQuery {
	query := (&TicketClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, selector),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, booking.TicketTable, booking.TicketColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Booking entity from the query.
// Returns a *NotFoundError when no Booking was found.
func (bq *BookingQuery) First(ctx context.Context) (*Booking, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{booking.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookingQuery) FirstX(ctx context.Context) *Booking {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Booking ID from the query.
// Returns a *NotFoundError when no Booking ID was found.
func (bq *BookingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{booking.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BookingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Booking entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Booking entity is found.
// Returns a *NotFoundError when no Booking entities are found.
func (bq *BookingQuery) Only(ctx context.Context) (*Booking, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{booking.Label}
	default:
		return nil, &NotSingularError{booking.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookingQuery) OnlyX(ctx context.Context) *Booking {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Booking ID in the query.
// Returns a *NotSingularError when more than one Booking ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BookingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{booking.Label}
	default:
		err = &NotSingularError{booking.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BookingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bookings.
func (bq *BookingQuery) All(ctx context.Context) ([]*Booking, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Booking, *BookingQuery]()
	return withInterceptors[[]*Booking](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookingQuery) AllX(ctx context.Context) []*Booking {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Booking IDs.
func (bq *BookingQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(booking.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BookingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BookingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BookingQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookingQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookingQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookingQuery) Clone() *BookingQuery {
	if bq == nil {
		return nil
	}
	return &BookingQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]booking.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Booking{}, bq.predicates...),
		withTicket: bq.withTicket.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithTicket tells the query-builder to eager-load the nodes that are connected to
// the "ticket" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BookingQuery) WithTicket(opts ...func(*TicketQuery)) *BookingQuery {
	query := (&TicketClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withTicket = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CustomerID uuid.UUID `json:"customer_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Booking.Query().
//		GroupBy(booking.FieldCustomerID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BookingQuery) GroupBy(field string, fields ...string) *BookingGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BookingGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = booking.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CustomerID uuid.UUID `json:"customer_id,omitempty"`
//	}
//
//	client.Booking.Query().
//		Select(booking.FieldCustomerID).
//		Scan(ctx, &v)
func (bq *BookingQuery) Select(fields ...string) *BookingSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BookingSelect{BookingQuery: bq}
	sbuild.label = booking.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookingSelect configured with the given aggregations.
func (bq *BookingQuery) Aggregate(fns ...AggregateFunc) *BookingSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BookingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !booking.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Booking, error) {
	var (
		nodes       = []*Booking{}
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withTicket != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Booking).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Booking{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withTicket; query != nil {
		if err := bq.loadTicket(ctx, query, nodes,
			func(n *Booking) { n.Edges.Ticket = []*Ticket{} },
			func(n *Booking, e *Ticket) { n.Edges.Ticket = append(n.Edges.Ticket, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BookingQuery) loadTicket(ctx context.Context, query *TicketQuery, nodes []*Booking, init func(*Booking), assign func(*Booking, *Ticket)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Booking)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(ticket.FieldBookingID)
	}
	query.Where(predicate.Ticket(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(booking.TicketColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BookingID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "booking_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BookingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeUUID))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for i := range fields {
			if fields[i] != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(booking.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = booking.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookingGroupBy is the group-by builder for Booking entities.
type BookingGroupBy struct {
	selector
	build *BookingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookingGroupBy) Aggregate(fns ...AggregateFunc) *BookingGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BookingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookingQuery, *BookingGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BookingGroupBy) sqlScan(ctx context.Context, root *BookingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookingSelect is the builder for selecting fields of Booking entities.
type BookingSelect struct {
	*BookingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BookingSelect) Aggregate(fns ...AggregateFunc) *BookingSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BookingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookingQuery, *BookingSelect](ctx, bs.BookingQuery, bs, bs.inters, v)
}

func (bs *BookingSelect) sqlScan(ctx context.Context, root *BookingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
