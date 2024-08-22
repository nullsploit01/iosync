// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"iosync/ent/migrate"

	"iosync/ent/apikey"
	"iosync/ent/device"
	"iosync/ent/session"
	"iosync/ent/topic"
	"iosync/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ApiKey is the client for interacting with the ApiKey builders.
	ApiKey *ApiKeyClient
	// Device is the client for interacting with the Device builders.
	Device *DeviceClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// Topic is the client for interacting with the Topic builders.
	Topic *TopicClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ApiKey = NewApiKeyClient(c.config)
	c.Device = NewDeviceClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.Topic = NewTopicClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		ApiKey:  NewApiKeyClient(cfg),
		Device:  NewDeviceClient(cfg),
		Session: NewSessionClient(cfg),
		Topic:   NewTopicClient(cfg),
		User:    NewUserClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		ApiKey:  NewApiKeyClient(cfg),
		Device:  NewDeviceClient(cfg),
		Session: NewSessionClient(cfg),
		Topic:   NewTopicClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ApiKey.
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
	c.ApiKey.Use(hooks...)
	c.Device.Use(hooks...)
	c.Session.Use(hooks...)
	c.Topic.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ApiKey.Intercept(interceptors...)
	c.Device.Intercept(interceptors...)
	c.Session.Intercept(interceptors...)
	c.Topic.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ApiKeyMutation:
		return c.ApiKey.mutate(ctx, m)
	case *DeviceMutation:
		return c.Device.mutate(ctx, m)
	case *SessionMutation:
		return c.Session.mutate(ctx, m)
	case *TopicMutation:
		return c.Topic.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ApiKeyClient is a client for the ApiKey schema.
type ApiKeyClient struct {
	config
}

// NewApiKeyClient returns a client for the ApiKey from the given config.
func NewApiKeyClient(c config) *ApiKeyClient {
	return &ApiKeyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `apikey.Hooks(f(g(h())))`.
func (c *ApiKeyClient) Use(hooks ...Hook) {
	c.hooks.ApiKey = append(c.hooks.ApiKey, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `apikey.Intercept(f(g(h())))`.
func (c *ApiKeyClient) Intercept(interceptors ...Interceptor) {
	c.inters.ApiKey = append(c.inters.ApiKey, interceptors...)
}

// Create returns a builder for creating a ApiKey entity.
func (c *ApiKeyClient) Create() *ApiKeyCreate {
	mutation := newApiKeyMutation(c.config, OpCreate)
	return &ApiKeyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ApiKey entities.
func (c *ApiKeyClient) CreateBulk(builders ...*ApiKeyCreate) *ApiKeyCreateBulk {
	return &ApiKeyCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ApiKeyClient) MapCreateBulk(slice any, setFunc func(*ApiKeyCreate, int)) *ApiKeyCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ApiKeyCreateBulk{err: fmt.Errorf("calling to ApiKeyClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ApiKeyCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ApiKeyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ApiKey.
func (c *ApiKeyClient) Update() *ApiKeyUpdate {
	mutation := newApiKeyMutation(c.config, OpUpdate)
	return &ApiKeyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApiKeyClient) UpdateOne(ak *ApiKey) *ApiKeyUpdateOne {
	mutation := newApiKeyMutation(c.config, OpUpdateOne, withApiKey(ak))
	return &ApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApiKeyClient) UpdateOneID(id int) *ApiKeyUpdateOne {
	mutation := newApiKeyMutation(c.config, OpUpdateOne, withApiKeyID(id))
	return &ApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ApiKey.
func (c *ApiKeyClient) Delete() *ApiKeyDelete {
	mutation := newApiKeyMutation(c.config, OpDelete)
	return &ApiKeyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ApiKeyClient) DeleteOne(ak *ApiKey) *ApiKeyDeleteOne {
	return c.DeleteOneID(ak.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ApiKeyClient) DeleteOneID(id int) *ApiKeyDeleteOne {
	builder := c.Delete().Where(apikey.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApiKeyDeleteOne{builder}
}

// Query returns a query builder for ApiKey.
func (c *ApiKeyClient) Query() *ApiKeyQuery {
	return &ApiKeyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeApiKey},
		inters: c.Interceptors(),
	}
}

// Get returns a ApiKey entity by its id.
func (c *ApiKeyClient) Get(ctx context.Context, id int) (*ApiKey, error) {
	return c.Query().Where(apikey.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApiKeyClient) GetX(ctx context.Context, id int) *ApiKey {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDevice queries the device edge of a ApiKey.
func (c *ApiKeyClient) QueryDevice(ak *ApiKey) *DeviceQuery {
	query := (&DeviceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ak.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, id),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, apikey.DeviceTable, apikey.DeviceColumn),
		)
		fromV = sqlgraph.Neighbors(ak.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTopics queries the topics edge of a ApiKey.
func (c *ApiKeyClient) QueryTopics(ak *ApiKey) *TopicQuery {
	query := (&TopicClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ak.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, apikey.TopicsTable, apikey.TopicsColumn),
		)
		fromV = sqlgraph.Neighbors(ak.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ApiKeyClient) Hooks() []Hook {
	return c.hooks.ApiKey
}

// Interceptors returns the client interceptors.
func (c *ApiKeyClient) Interceptors() []Interceptor {
	return c.inters.ApiKey
}

func (c *ApiKeyClient) mutate(ctx context.Context, m *ApiKeyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ApiKeyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ApiKeyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ApiKeyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ApiKey mutation op: %q", m.Op())
	}
}

// DeviceClient is a client for the Device schema.
type DeviceClient struct {
	config
}

// NewDeviceClient returns a client for the Device from the given config.
func NewDeviceClient(c config) *DeviceClient {
	return &DeviceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `device.Hooks(f(g(h())))`.
func (c *DeviceClient) Use(hooks ...Hook) {
	c.hooks.Device = append(c.hooks.Device, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `device.Intercept(f(g(h())))`.
func (c *DeviceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Device = append(c.inters.Device, interceptors...)
}

// Create returns a builder for creating a Device entity.
func (c *DeviceClient) Create() *DeviceCreate {
	mutation := newDeviceMutation(c.config, OpCreate)
	return &DeviceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Device entities.
func (c *DeviceClient) CreateBulk(builders ...*DeviceCreate) *DeviceCreateBulk {
	return &DeviceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DeviceClient) MapCreateBulk(slice any, setFunc func(*DeviceCreate, int)) *DeviceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DeviceCreateBulk{err: fmt.Errorf("calling to DeviceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DeviceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DeviceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Device.
func (c *DeviceClient) Update() *DeviceUpdate {
	mutation := newDeviceMutation(c.config, OpUpdate)
	return &DeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeviceClient) UpdateOne(d *Device) *DeviceUpdateOne {
	mutation := newDeviceMutation(c.config, OpUpdateOne, withDevice(d))
	return &DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeviceClient) UpdateOneID(id int) *DeviceUpdateOne {
	mutation := newDeviceMutation(c.config, OpUpdateOne, withDeviceID(id))
	return &DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Device.
func (c *DeviceClient) Delete() *DeviceDelete {
	mutation := newDeviceMutation(c.config, OpDelete)
	return &DeviceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DeviceClient) DeleteOne(d *Device) *DeviceDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DeviceClient) DeleteOneID(id int) *DeviceDeleteOne {
	builder := c.Delete().Where(device.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeviceDeleteOne{builder}
}

// Query returns a query builder for Device.
func (c *DeviceClient) Query() *DeviceQuery {
	return &DeviceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDevice},
		inters: c.Interceptors(),
	}
}

// Get returns a Device entity by its id.
func (c *DeviceClient) Get(ctx context.Context, id int) (*Device, error) {
	return c.Query().Where(device.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeviceClient) GetX(ctx context.Context, id int) *Device {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Device.
func (c *DeviceClient) QueryUser(d *Device) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, device.UserTable, device.UserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAPIKey queries the api_key edge of a Device.
func (c *DeviceClient) QueryAPIKey(d *Device) *ApiKeyQuery {
	query := (&ApiKeyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, id),
			sqlgraph.To(apikey.Table, apikey.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, device.APIKeyTable, device.APIKeyColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTopics queries the topics edge of a Device.
func (c *DeviceClient) QueryTopics(d *Device) *TopicQuery {
	query := (&TopicClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, device.TopicsTable, device.TopicsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeviceClient) Hooks() []Hook {
	return c.hooks.Device
}

// Interceptors returns the client interceptors.
func (c *DeviceClient) Interceptors() []Interceptor {
	return c.inters.Device
}

func (c *DeviceClient) mutate(ctx context.Context, m *DeviceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DeviceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DeviceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Device mutation op: %q", m.Op())
	}
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `session.Intercept(f(g(h())))`.
func (c *SessionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Session = append(c.inters.Session, interceptors...)
}

// Create returns a builder for creating a Session entity.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SessionClient) MapCreateBulk(slice any, setFunc func(*SessionCreate, int)) *SessionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SessionCreateBulk{err: fmt.Errorf("calling to SessionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SessionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id int) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SessionClient) DeleteOneID(id int) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSession},
		inters: c.Interceptors(),
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id int) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id int) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Session.
func (c *SessionClient) QueryUser(s *Session) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, session.UserTable, session.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// Interceptors returns the client interceptors.
func (c *SessionClient) Interceptors() []Interceptor {
	return c.inters.Session
}

func (c *SessionClient) mutate(ctx context.Context, m *SessionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SessionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SessionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Session mutation op: %q", m.Op())
	}
}

// TopicClient is a client for the Topic schema.
type TopicClient struct {
	config
}

// NewTopicClient returns a client for the Topic from the given config.
func NewTopicClient(c config) *TopicClient {
	return &TopicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `topic.Hooks(f(g(h())))`.
func (c *TopicClient) Use(hooks ...Hook) {
	c.hooks.Topic = append(c.hooks.Topic, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `topic.Intercept(f(g(h())))`.
func (c *TopicClient) Intercept(interceptors ...Interceptor) {
	c.inters.Topic = append(c.inters.Topic, interceptors...)
}

// Create returns a builder for creating a Topic entity.
func (c *TopicClient) Create() *TopicCreate {
	mutation := newTopicMutation(c.config, OpCreate)
	return &TopicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Topic entities.
func (c *TopicClient) CreateBulk(builders ...*TopicCreate) *TopicCreateBulk {
	return &TopicCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TopicClient) MapCreateBulk(slice any, setFunc func(*TopicCreate, int)) *TopicCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TopicCreateBulk{err: fmt.Errorf("calling to TopicClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TopicCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TopicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Topic.
func (c *TopicClient) Update() *TopicUpdate {
	mutation := newTopicMutation(c.config, OpUpdate)
	return &TopicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TopicClient) UpdateOne(t *Topic) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopic(t))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TopicClient) UpdateOneID(id int) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopicID(id))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Topic.
func (c *TopicClient) Delete() *TopicDelete {
	mutation := newTopicMutation(c.config, OpDelete)
	return &TopicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TopicClient) DeleteOne(t *Topic) *TopicDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TopicClient) DeleteOneID(id int) *TopicDeleteOne {
	builder := c.Delete().Where(topic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TopicDeleteOne{builder}
}

// Query returns a query builder for Topic.
func (c *TopicClient) Query() *TopicQuery {
	return &TopicQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTopic},
		inters: c.Interceptors(),
	}
}

// Get returns a Topic entity by its id.
func (c *TopicClient) Get(ctx context.Context, id int) (*Topic, error) {
	return c.Query().Where(topic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TopicClient) GetX(ctx context.Context, id int) *Topic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDevice queries the device edge of a Topic.
func (c *TopicClient) QueryDevice(t *Topic) *DeviceQuery {
	query := (&DeviceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topic.Table, topic.FieldID, id),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, topic.DeviceTable, topic.DeviceColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAPIKey queries the api_key edge of a Topic.
func (c *TopicClient) QueryAPIKey(t *Topic) *ApiKeyQuery {
	query := (&ApiKeyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topic.Table, topic.FieldID, id),
			sqlgraph.To(apikey.Table, apikey.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, topic.APIKeyTable, topic.APIKeyColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TopicClient) Hooks() []Hook {
	return c.hooks.Topic
}

// Interceptors returns the client interceptors.
func (c *TopicClient) Interceptors() []Interceptor {
	return c.inters.Topic
}

func (c *TopicClient) mutate(ctx context.Context, m *TopicMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TopicCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TopicUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TopicDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Topic mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDevices queries the devices edge of a User.
func (c *UserClient) QueryDevices(u *User) *DeviceQuery {
	query := (&DeviceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DevicesTable, user.DevicesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySessions queries the sessions edge of a User.
func (c *UserClient) QuerySessions(u *User) *SessionQuery {
	query := (&SessionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SessionsTable, user.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ApiKey, Device, Session, Topic, User []ent.Hook
	}
	inters struct {
		ApiKey, Device, Session, Topic, User []ent.Interceptor
	}
)
