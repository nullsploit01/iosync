// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/nullsploit01/iosync/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nullsploit01/iosync/ent/node"
	"github.com/nullsploit01/iosync/ent/nodeapikey"
	"github.com/nullsploit01/iosync/ent/nodevalues"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Node is the client for interacting with the Node builders.
	Node *NodeClient
	// NodeApiKey is the client for interacting with the NodeApiKey builders.
	NodeApiKey *NodeApiKeyClient
	// NodeValues is the client for interacting with the NodeValues builders.
	NodeValues *NodeValuesClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Node = NewNodeClient(c.config)
	c.NodeApiKey = NewNodeApiKeyClient(c.config)
	c.NodeValues = NewNodeValuesClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Node:       NewNodeClient(cfg),
		NodeApiKey: NewNodeApiKeyClient(cfg),
		NodeValues: NewNodeValuesClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Node:       NewNodeClient(cfg),
		NodeApiKey: NewNodeApiKeyClient(cfg),
		NodeValues: NewNodeValuesClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Node.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Node.Use(hooks...)
	c.NodeApiKey.Use(hooks...)
	c.NodeValues.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Node.Intercept(interceptors...)
	c.NodeApiKey.Intercept(interceptors...)
	c.NodeValues.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *NodeMutation:
		return c.Node.mutate(ctx, m)
	case *NodeApiKeyMutation:
		return c.NodeApiKey.mutate(ctx, m)
	case *NodeValuesMutation:
		return c.NodeValues.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// NodeClient is a client for the Node schema.
type NodeClient struct {
	config
}

// NewNodeClient returns a client for the Node from the given config.
func NewNodeClient(c config) *NodeClient {
	return &NodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `node.Hooks(f(g(h())))`.
func (c *NodeClient) Use(hooks ...Hook) {
	c.hooks.Node = append(c.hooks.Node, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `node.Intercept(f(g(h())))`.
func (c *NodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Node = append(c.inters.Node, interceptors...)
}

// Create returns a builder for creating a Node entity.
func (c *NodeClient) Create() *NodeCreate {
	mutation := newNodeMutation(c.config, OpCreate)
	return &NodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Node entities.
func (c *NodeClient) CreateBulk(builders ...*NodeCreate) *NodeCreateBulk {
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *NodeClient) MapCreateBulk(slice any, setFunc func(*NodeCreate, int)) *NodeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &NodeCreateBulk{err: fmt.Errorf("calling to NodeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*NodeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Node.
func (c *NodeClient) Update() *NodeUpdate {
	mutation := newNodeMutation(c.config, OpUpdate)
	return &NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeClient) UpdateOne(n *Node) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNode(n))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeClient) UpdateOneID(id int) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNodeID(id))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Node.
func (c *NodeClient) Delete() *NodeDelete {
	mutation := newNodeMutation(c.config, OpDelete)
	return &NodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NodeClient) DeleteOne(n *Node) *NodeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NodeClient) DeleteOneID(id int) *NodeDeleteOne {
	builder := c.Delete().Where(node.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeDeleteOne{builder}
}

// Query returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeNode},
		inters: c.Interceptors(),
	}
}

// Get returns a Node entity by its id.
func (c *NodeClient) Get(ctx context.Context, id int) (*Node, error) {
	return c.Query().Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeClient) GetX(ctx context.Context, id int) *Node {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAPIKeys queries the api_keys edge of a Node.
func (c *NodeClient) QueryAPIKeys(n *Node) *NodeApiKeyQuery {
	query := (&NodeApiKeyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(nodeapikey.Table, nodeapikey.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, node.APIKeysTable, node.APIKeysColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeClient) Hooks() []Hook {
	return c.hooks.Node
}

// Interceptors returns the client interceptors.
func (c *NodeClient) Interceptors() []Interceptor {
	return c.inters.Node
}

func (c *NodeClient) mutate(ctx context.Context, m *NodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&NodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&NodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Node mutation op: %q", m.Op())
	}
}

// NodeApiKeyClient is a client for the NodeApiKey schema.
type NodeApiKeyClient struct {
	config
}

// NewNodeApiKeyClient returns a client for the NodeApiKey from the given config.
func NewNodeApiKeyClient(c config) *NodeApiKeyClient {
	return &NodeApiKeyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `nodeapikey.Hooks(f(g(h())))`.
func (c *NodeApiKeyClient) Use(hooks ...Hook) {
	c.hooks.NodeApiKey = append(c.hooks.NodeApiKey, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `nodeapikey.Intercept(f(g(h())))`.
func (c *NodeApiKeyClient) Intercept(interceptors ...Interceptor) {
	c.inters.NodeApiKey = append(c.inters.NodeApiKey, interceptors...)
}

// Create returns a builder for creating a NodeApiKey entity.
func (c *NodeApiKeyClient) Create() *NodeApiKeyCreate {
	mutation := newNodeApiKeyMutation(c.config, OpCreate)
	return &NodeApiKeyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NodeApiKey entities.
func (c *NodeApiKeyClient) CreateBulk(builders ...*NodeApiKeyCreate) *NodeApiKeyCreateBulk {
	return &NodeApiKeyCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *NodeApiKeyClient) MapCreateBulk(slice any, setFunc func(*NodeApiKeyCreate, int)) *NodeApiKeyCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &NodeApiKeyCreateBulk{err: fmt.Errorf("calling to NodeApiKeyClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*NodeApiKeyCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &NodeApiKeyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NodeApiKey.
func (c *NodeApiKeyClient) Update() *NodeApiKeyUpdate {
	mutation := newNodeApiKeyMutation(c.config, OpUpdate)
	return &NodeApiKeyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeApiKeyClient) UpdateOne(nak *NodeApiKey) *NodeApiKeyUpdateOne {
	mutation := newNodeApiKeyMutation(c.config, OpUpdateOne, withNodeApiKey(nak))
	return &NodeApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeApiKeyClient) UpdateOneID(id int) *NodeApiKeyUpdateOne {
	mutation := newNodeApiKeyMutation(c.config, OpUpdateOne, withNodeApiKeyID(id))
	return &NodeApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NodeApiKey.
func (c *NodeApiKeyClient) Delete() *NodeApiKeyDelete {
	mutation := newNodeApiKeyMutation(c.config, OpDelete)
	return &NodeApiKeyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NodeApiKeyClient) DeleteOne(nak *NodeApiKey) *NodeApiKeyDeleteOne {
	return c.DeleteOneID(nak.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NodeApiKeyClient) DeleteOneID(id int) *NodeApiKeyDeleteOne {
	builder := c.Delete().Where(nodeapikey.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeApiKeyDeleteOne{builder}
}

// Query returns a query builder for NodeApiKey.
func (c *NodeApiKeyClient) Query() *NodeApiKeyQuery {
	return &NodeApiKeyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeNodeApiKey},
		inters: c.Interceptors(),
	}
}

// Get returns a NodeApiKey entity by its id.
func (c *NodeApiKeyClient) Get(ctx context.Context, id int) (*NodeApiKey, error) {
	return c.Query().Where(nodeapikey.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeApiKeyClient) GetX(ctx context.Context, id int) *NodeApiKey {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNode queries the node edge of a NodeApiKey.
func (c *NodeApiKeyClient) QueryNode(nak *NodeApiKey) *NodeQuery {
	query := (&NodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := nak.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(nodeapikey.Table, nodeapikey.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, nodeapikey.NodeTable, nodeapikey.NodeColumn),
		)
		fromV = sqlgraph.Neighbors(nak.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryValues queries the values edge of a NodeApiKey.
func (c *NodeApiKeyClient) QueryValues(nak *NodeApiKey) *NodeValuesQuery {
	query := (&NodeValuesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := nak.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(nodeapikey.Table, nodeapikey.FieldID, id),
			sqlgraph.To(nodevalues.Table, nodevalues.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, nodeapikey.ValuesTable, nodeapikey.ValuesColumn),
		)
		fromV = sqlgraph.Neighbors(nak.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeApiKeyClient) Hooks() []Hook {
	return c.hooks.NodeApiKey
}

// Interceptors returns the client interceptors.
func (c *NodeApiKeyClient) Interceptors() []Interceptor {
	return c.inters.NodeApiKey
}

func (c *NodeApiKeyClient) mutate(ctx context.Context, m *NodeApiKeyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&NodeApiKeyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&NodeApiKeyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&NodeApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&NodeApiKeyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown NodeApiKey mutation op: %q", m.Op())
	}
}

// NodeValuesClient is a client for the NodeValues schema.
type NodeValuesClient struct {
	config
}

// NewNodeValuesClient returns a client for the NodeValues from the given config.
func NewNodeValuesClient(c config) *NodeValuesClient {
	return &NodeValuesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `nodevalues.Hooks(f(g(h())))`.
func (c *NodeValuesClient) Use(hooks ...Hook) {
	c.hooks.NodeValues = append(c.hooks.NodeValues, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `nodevalues.Intercept(f(g(h())))`.
func (c *NodeValuesClient) Intercept(interceptors ...Interceptor) {
	c.inters.NodeValues = append(c.inters.NodeValues, interceptors...)
}

// Create returns a builder for creating a NodeValues entity.
func (c *NodeValuesClient) Create() *NodeValuesCreate {
	mutation := newNodeValuesMutation(c.config, OpCreate)
	return &NodeValuesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NodeValues entities.
func (c *NodeValuesClient) CreateBulk(builders ...*NodeValuesCreate) *NodeValuesCreateBulk {
	return &NodeValuesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *NodeValuesClient) MapCreateBulk(slice any, setFunc func(*NodeValuesCreate, int)) *NodeValuesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &NodeValuesCreateBulk{err: fmt.Errorf("calling to NodeValuesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*NodeValuesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &NodeValuesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NodeValues.
func (c *NodeValuesClient) Update() *NodeValuesUpdate {
	mutation := newNodeValuesMutation(c.config, OpUpdate)
	return &NodeValuesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeValuesClient) UpdateOne(nv *NodeValues) *NodeValuesUpdateOne {
	mutation := newNodeValuesMutation(c.config, OpUpdateOne, withNodeValues(nv))
	return &NodeValuesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeValuesClient) UpdateOneID(id int) *NodeValuesUpdateOne {
	mutation := newNodeValuesMutation(c.config, OpUpdateOne, withNodeValuesID(id))
	return &NodeValuesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NodeValues.
func (c *NodeValuesClient) Delete() *NodeValuesDelete {
	mutation := newNodeValuesMutation(c.config, OpDelete)
	return &NodeValuesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NodeValuesClient) DeleteOne(nv *NodeValues) *NodeValuesDeleteOne {
	return c.DeleteOneID(nv.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NodeValuesClient) DeleteOneID(id int) *NodeValuesDeleteOne {
	builder := c.Delete().Where(nodevalues.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeValuesDeleteOne{builder}
}

// Query returns a query builder for NodeValues.
func (c *NodeValuesClient) Query() *NodeValuesQuery {
	return &NodeValuesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeNodeValues},
		inters: c.Interceptors(),
	}
}

// Get returns a NodeValues entity by its id.
func (c *NodeValuesClient) Get(ctx context.Context, id int) (*NodeValues, error) {
	return c.Query().Where(nodevalues.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeValuesClient) GetX(ctx context.Context, id int) *NodeValues {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNodeAPIKey queries the node_api_key edge of a NodeValues.
func (c *NodeValuesClient) QueryNodeAPIKey(nv *NodeValues) *NodeApiKeyQuery {
	query := (&NodeApiKeyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := nv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(nodevalues.Table, nodevalues.FieldID, id),
			sqlgraph.To(nodeapikey.Table, nodeapikey.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, nodevalues.NodeAPIKeyTable, nodevalues.NodeAPIKeyColumn),
		)
		fromV = sqlgraph.Neighbors(nv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeValuesClient) Hooks() []Hook {
	return c.hooks.NodeValues
}

// Interceptors returns the client interceptors.
func (c *NodeValuesClient) Interceptors() []Interceptor {
	return c.inters.NodeValues
}

func (c *NodeValuesClient) mutate(ctx context.Context, m *NodeValuesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&NodeValuesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&NodeValuesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&NodeValuesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&NodeValuesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown NodeValues mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Node, NodeApiKey, NodeValues []ent.Hook
	}
	inters struct {
		Node, NodeApiKey, NodeValues []ent.Interceptor
	}
)
