// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/api/src/infra/db/model/predicate"
	"github.com/Southclaws/storyden/api/src/infra/db/model/rule"
	"github.com/Southclaws/storyden/api/src/infra/db/model/server"
	"github.com/Southclaws/storyden/api/src/infra/db/model/user"
)

// ServerQuery is the builder for querying Server entities.
type ServerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Server
	// eager-loading edges.
	withRu   *RuleQuery
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServerQuery builder.
func (sq *ServerQuery) Where(ps ...predicate.Server) *ServerQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ServerQuery) Limit(limit int) *ServerQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ServerQuery) Offset(offset int) *ServerQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ServerQuery) Unique(unique bool) *ServerQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *ServerQuery) Order(o ...OrderFunc) *ServerQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryRu chains the current query on the "ru" edge.
func (sq *ServerQuery) QueryRu() *RuleQuery {
	query := &RuleQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, selector),
			sqlgraph.To(rule.Table, rule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, server.RuTable, server.RuColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "User" edge.
func (sq *ServerQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, server.UserTable, server.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Server entity from the query.
// Returns a *NotFoundError when no Server was found.
func (sq *ServerQuery) First(ctx context.Context) (*Server, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{server.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ServerQuery) FirstX(ctx context.Context) *Server {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Server ID from the query.
// Returns a *NotFoundError when no Server ID was found.
func (sq *ServerQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{server.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ServerQuery) FirstIDX(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Server entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Server entity is found.
// Returns a *NotFoundError when no Server entities are found.
func (sq *ServerQuery) Only(ctx context.Context) (*Server, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{server.Label}
	default:
		return nil, &NotSingularError{server.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ServerQuery) OnlyX(ctx context.Context) *Server {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Server ID in the query.
// Returns a *NotSingularError when more than one Server ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ServerQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = &NotSingularError{server.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ServerQuery) OnlyIDX(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Servers.
func (sq *ServerQuery) All(ctx context.Context) ([]*Server, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ServerQuery) AllX(ctx context.Context) []*Server {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Server IDs.
func (sq *ServerQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := sq.Select(server.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ServerQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ServerQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ServerQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ServerQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ServerQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ServerQuery) Clone() *ServerQuery {
	if sq == nil {
		return nil
	}
	return &ServerQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		predicates: append([]predicate.Server{}, sq.predicates...),
		withRu:     sq.withRu.Clone(),
		withUser:   sq.withUser.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithRu tells the query-builder to eager-load the nodes that are connected to
// the "ru" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ServerQuery) WithRu(opts ...func(*RuleQuery)) *ServerQuery {
	query := &RuleQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withRu = query
	return sq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "User" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ServerQuery) WithUser(opts ...func(*UserQuery)) *ServerQuery {
	query := &UserQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withUser = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IP string `json:"ip,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Server.Query().
//		GroupBy(server.FieldIP).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (sq *ServerQuery) GroupBy(field string, fields ...string) *ServerGroupBy {
	group := &ServerGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IP string `json:"ip,omitempty"`
//	}
//
//	client.Server.Query().
//		Select(server.FieldIP).
//		Scan(ctx, &v)
//
func (sq *ServerQuery) Select(fields ...string) *ServerSelect {
	sq.fields = append(sq.fields, fields...)
	return &ServerSelect{ServerQuery: sq}
}

func (sq *ServerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !server.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ServerQuery) sqlAll(ctx context.Context) ([]*Server, error) {
	var (
		nodes       = []*Server{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withRu != nil,
			sq.withUser != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, server.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Server{config: sq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withRu; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Server)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Ru = []*Rule{}
		}
		query.withFKs = true
		query.Where(predicate.Rule(func(s *sql.Selector) {
			s.Where(sql.InValues(server.RuColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.server_ru
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "server_ru" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "server_ru" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Ru = append(node.Edges.Ru, n)
		}
	}

	if query := sq.withUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Server)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.User = []*User{}
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(server.UserColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.server_user
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "server_user" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "server_user" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.User = append(node.Edges.User, n)
		}
	}

	return nodes, nil
}

func (sq *ServerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ServerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (sq *ServerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   server.Table,
			Columns: server.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: server.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, server.FieldID)
		for i := range fields {
			if fields[i] != server.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ServerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(server.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = server.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServerGroupBy is the group-by builder for Server entities.
type ServerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ServerGroupBy) Aggregate(fns ...AggregateFunc) *ServerGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *ServerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *ServerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("model: ServerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *ServerGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgb *ServerGroupBy) StringX(ctx context.Context) string {
	v, err := sgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("model: ServerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *ServerGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgb *ServerGroupBy) IntX(ctx context.Context) int {
	v, err := sgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("model: ServerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *ServerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgb *ServerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("model: ServerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *ServerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *ServerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgb *ServerGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *ServerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgb.fields {
		if !server.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ServerGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// ServerSelect is the builder for selecting fields of Server entities.
type ServerSelect struct {
	*ServerQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ServerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.ServerQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *ServerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("model: ServerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *ServerSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ss *ServerSelect) StringX(ctx context.Context) string {
	v, err := ss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("model: ServerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *ServerSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ss *ServerSelect) IntX(ctx context.Context) int {
	v, err := ss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("model: ServerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *ServerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ss *ServerSelect) Float64X(ctx context.Context) float64 {
	v, err := ss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("model: ServerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *ServerSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ss *ServerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{server.Label}
	default:
		err = fmt.Errorf("model: ServerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ss *ServerSelect) BoolX(ctx context.Context) bool {
	v, err := ss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *ServerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
