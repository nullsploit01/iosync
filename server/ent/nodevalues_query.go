// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nullsploit01/iosync/ent/nodeapikey"
	"github.com/nullsploit01/iosync/ent/nodevalues"
	"github.com/nullsploit01/iosync/ent/predicate"
)

// NodeValuesQuery is the builder for querying NodeValues entities.
type NodeValuesQuery struct {
	config
	ctx            *QueryContext
	order          []nodevalues.OrderOption
	inters         []Interceptor
	predicates     []predicate.NodeValues
	withNodeAPIKey *NodeApiKeyQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NodeValuesQuery builder.
func (nvq *NodeValuesQuery) Where(ps ...predicate.NodeValues) *NodeValuesQuery {
	nvq.predicates = append(nvq.predicates, ps...)
	return nvq
}

// Limit the number of records to be returned by this query.
func (nvq *NodeValuesQuery) Limit(limit int) *NodeValuesQuery {
	nvq.ctx.Limit = &limit
	return nvq
}

// Offset to start from.
func (nvq *NodeValuesQuery) Offset(offset int) *NodeValuesQuery {
	nvq.ctx.Offset = &offset
	return nvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nvq *NodeValuesQuery) Unique(unique bool) *NodeValuesQuery {
	nvq.ctx.Unique = &unique
	return nvq
}

// Order specifies how the records should be ordered.
func (nvq *NodeValuesQuery) Order(o ...nodevalues.OrderOption) *NodeValuesQuery {
	nvq.order = append(nvq.order, o...)
	return nvq
}

// QueryNodeAPIKey chains the current query on the "node_api_key" edge.
func (nvq *NodeValuesQuery) QueryNodeAPIKey() *NodeApiKeyQuery {
	query := (&NodeApiKeyClient{config: nvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nodevalues.Table, nodevalues.FieldID, selector),
			sqlgraph.To(nodeapikey.Table, nodeapikey.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, nodevalues.NodeAPIKeyTable, nodevalues.NodeAPIKeyColumn),
		)
		fromU = sqlgraph.SetNeighbors(nvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NodeValues entity from the query.
// Returns a *NotFoundError when no NodeValues was found.
func (nvq *NodeValuesQuery) First(ctx context.Context) (*NodeValues, error) {
	nodes, err := nvq.Limit(1).All(setContextOp(ctx, nvq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nodevalues.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nvq *NodeValuesQuery) FirstX(ctx context.Context) *NodeValues {
	node, err := nvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NodeValues ID from the query.
// Returns a *NotFoundError when no NodeValues ID was found.
func (nvq *NodeValuesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nvq.Limit(1).IDs(setContextOp(ctx, nvq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nodevalues.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nvq *NodeValuesQuery) FirstIDX(ctx context.Context) int {
	id, err := nvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NodeValues entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NodeValues entity is found.
// Returns a *NotFoundError when no NodeValues entities are found.
func (nvq *NodeValuesQuery) Only(ctx context.Context) (*NodeValues, error) {
	nodes, err := nvq.Limit(2).All(setContextOp(ctx, nvq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nodevalues.Label}
	default:
		return nil, &NotSingularError{nodevalues.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nvq *NodeValuesQuery) OnlyX(ctx context.Context) *NodeValues {
	node, err := nvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NodeValues ID in the query.
// Returns a *NotSingularError when more than one NodeValues ID is found.
// Returns a *NotFoundError when no entities are found.
func (nvq *NodeValuesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nvq.Limit(2).IDs(setContextOp(ctx, nvq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nodevalues.Label}
	default:
		err = &NotSingularError{nodevalues.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nvq *NodeValuesQuery) OnlyIDX(ctx context.Context) int {
	id, err := nvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NodeValuesSlice.
func (nvq *NodeValuesQuery) All(ctx context.Context) ([]*NodeValues, error) {
	ctx = setContextOp(ctx, nvq.ctx, ent.OpQueryAll)
	if err := nvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NodeValues, *NodeValuesQuery]()
	return withInterceptors[[]*NodeValues](ctx, nvq, qr, nvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nvq *NodeValuesQuery) AllX(ctx context.Context) []*NodeValues {
	nodes, err := nvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NodeValues IDs.
func (nvq *NodeValuesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nvq.ctx.Unique == nil && nvq.path != nil {
		nvq.Unique(true)
	}
	ctx = setContextOp(ctx, nvq.ctx, ent.OpQueryIDs)
	if err = nvq.Select(nodevalues.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nvq *NodeValuesQuery) IDsX(ctx context.Context) []int {
	ids, err := nvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nvq *NodeValuesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nvq.ctx, ent.OpQueryCount)
	if err := nvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nvq, querierCount[*NodeValuesQuery](), nvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nvq *NodeValuesQuery) CountX(ctx context.Context) int {
	count, err := nvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nvq *NodeValuesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nvq.ctx, ent.OpQueryExist)
	switch _, err := nvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nvq *NodeValuesQuery) ExistX(ctx context.Context) bool {
	exist, err := nvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeValuesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nvq *NodeValuesQuery) Clone() *NodeValuesQuery {
	if nvq == nil {
		return nil
	}
	return &NodeValuesQuery{
		config:         nvq.config,
		ctx:            nvq.ctx.Clone(),
		order:          append([]nodevalues.OrderOption{}, nvq.order...),
		inters:         append([]Interceptor{}, nvq.inters...),
		predicates:     append([]predicate.NodeValues{}, nvq.predicates...),
		withNodeAPIKey: nvq.withNodeAPIKey.Clone(),
		// clone intermediate query.
		sql:  nvq.sql.Clone(),
		path: nvq.path,
	}
}

// WithNodeAPIKey tells the query-builder to eager-load the nodes that are connected to
// the "node_api_key" edge. The optional arguments are used to configure the query builder of the edge.
func (nvq *NodeValuesQuery) WithNodeAPIKey(opts ...func(*NodeApiKeyQuery)) *NodeValuesQuery {
	query := (&NodeApiKeyClient{config: nvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nvq.withNodeAPIKey = query
	return nvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NodeValues.Query().
//		GroupBy(nodevalues.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nvq *NodeValuesQuery) GroupBy(field string, fields ...string) *NodeValuesGroupBy {
	nvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NodeValuesGroupBy{build: nvq}
	grbuild.flds = &nvq.ctx.Fields
	grbuild.label = nodevalues.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//	}
//
//	client.NodeValues.Query().
//		Select(nodevalues.FieldValue).
//		Scan(ctx, &v)
func (nvq *NodeValuesQuery) Select(fields ...string) *NodeValuesSelect {
	nvq.ctx.Fields = append(nvq.ctx.Fields, fields...)
	sbuild := &NodeValuesSelect{NodeValuesQuery: nvq}
	sbuild.label = nodevalues.Label
	sbuild.flds, sbuild.scan = &nvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NodeValuesSelect configured with the given aggregations.
func (nvq *NodeValuesQuery) Aggregate(fns ...AggregateFunc) *NodeValuesSelect {
	return nvq.Select().Aggregate(fns...)
}

func (nvq *NodeValuesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nvq); err != nil {
				return err
			}
		}
	}
	for _, f := range nvq.ctx.Fields {
		if !nodevalues.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nvq.path != nil {
		prev, err := nvq.path(ctx)
		if err != nil {
			return err
		}
		nvq.sql = prev
	}
	return nil
}

func (nvq *NodeValuesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NodeValues, error) {
	var (
		nodes       = []*NodeValues{}
		withFKs     = nvq.withFKs
		_spec       = nvq.querySpec()
		loadedTypes = [1]bool{
			nvq.withNodeAPIKey != nil,
		}
	)
	if nvq.withNodeAPIKey != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, nodevalues.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NodeValues).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NodeValues{config: nvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nvq.withNodeAPIKey; query != nil {
		if err := nvq.loadNodeAPIKey(ctx, query, nodes, nil,
			func(n *NodeValues, e *NodeApiKey) { n.Edges.NodeAPIKey = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nvq *NodeValuesQuery) loadNodeAPIKey(ctx context.Context, query *NodeApiKeyQuery, nodes []*NodeValues, init func(*NodeValues), assign func(*NodeValues, *NodeApiKey)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*NodeValues)
	for i := range nodes {
		if nodes[i].node_api_key_values == nil {
			continue
		}
		fk := *nodes[i].node_api_key_values
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(nodeapikey.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "node_api_key_values" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (nvq *NodeValuesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nvq.querySpec()
	_spec.Node.Columns = nvq.ctx.Fields
	if len(nvq.ctx.Fields) > 0 {
		_spec.Unique = nvq.ctx.Unique != nil && *nvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nvq.driver, _spec)
}

func (nvq *NodeValuesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nodevalues.Table, nodevalues.Columns, sqlgraph.NewFieldSpec(nodevalues.FieldID, field.TypeInt))
	_spec.From = nvq.sql
	if unique := nvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nvq.path != nil {
		_spec.Unique = true
	}
	if fields := nvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nodevalues.FieldID)
		for i := range fields {
			if fields[i] != nodevalues.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nvq *NodeValuesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nvq.driver.Dialect())
	t1 := builder.Table(nodevalues.Table)
	columns := nvq.ctx.Fields
	if len(columns) == 0 {
		columns = nodevalues.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nvq.sql != nil {
		selector = nvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nvq.ctx.Unique != nil && *nvq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nvq.predicates {
		p(selector)
	}
	for _, p := range nvq.order {
		p(selector)
	}
	if offset := nvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NodeValuesGroupBy is the group-by builder for NodeValues entities.
type NodeValuesGroupBy struct {
	selector
	build *NodeValuesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nvgb *NodeValuesGroupBy) Aggregate(fns ...AggregateFunc) *NodeValuesGroupBy {
	nvgb.fns = append(nvgb.fns, fns...)
	return nvgb
}

// Scan applies the selector query and scans the result into the given value.
func (nvgb *NodeValuesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nvgb.build.ctx, ent.OpQueryGroupBy)
	if err := nvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeValuesQuery, *NodeValuesGroupBy](ctx, nvgb.build, nvgb, nvgb.build.inters, v)
}

func (nvgb *NodeValuesGroupBy) sqlScan(ctx context.Context, root *NodeValuesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nvgb.fns))
	for _, fn := range nvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nvgb.flds)+len(nvgb.fns))
		for _, f := range *nvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NodeValuesSelect is the builder for selecting fields of NodeValues entities.
type NodeValuesSelect struct {
	*NodeValuesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nvs *NodeValuesSelect) Aggregate(fns ...AggregateFunc) *NodeValuesSelect {
	nvs.fns = append(nvs.fns, fns...)
	return nvs
}

// Scan applies the selector query and scans the result into the given value.
func (nvs *NodeValuesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nvs.ctx, ent.OpQuerySelect)
	if err := nvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeValuesQuery, *NodeValuesSelect](ctx, nvs.NodeValuesQuery, nvs, nvs.inters, v)
}

func (nvs *NodeValuesSelect) sqlScan(ctx context.Context, root *NodeValuesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nvs.fns))
	for _, fn := range nvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
