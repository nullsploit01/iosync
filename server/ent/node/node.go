// Code generated by ent, DO NOT EDIT.

package node

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldIsOnline holds the string denoting the is_online field in the database.
	FieldIsOnline = "is_online"
	// FieldLastOnlineAt holds the string denoting the last_online_at field in the database.
	FieldLastOnlineAt = "last_online_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAPIKeys holds the string denoting the api_keys edge name in mutations.
	EdgeAPIKeys = "api_keys"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// APIKeysTable is the table that holds the api_keys relation/edge.
	APIKeysTable = "node_api_keys"
	// APIKeysInverseTable is the table name for the NodeApiKey entity.
	// It exists in this package in order to avoid circular dependency with the "nodeapikey" package.
	APIKeysInverseTable = "node_api_keys"
	// APIKeysColumn is the table column denoting the api_keys relation/edge.
	APIKeysColumn = "node_api_keys"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldIdentifier,
	FieldIsActive,
	FieldIsOnline,
	FieldLastOnlineAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIdentifier holds the default value on creation for the "identifier" field.
	DefaultIdentifier func() string
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultIsOnline holds the default value on creation for the "is_online" field.
	DefaultIsOnline bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Node queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIdentifier orders the results by the identifier field.
func ByIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentifier, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByIsOnline orders the results by the is_online field.
func ByIsOnline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnline, opts...).ToFunc()
}

// ByLastOnlineAt orders the results by the last_online_at field.
func ByLastOnlineAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastOnlineAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAPIKeysCount orders the results by api_keys count.
func ByAPIKeysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAPIKeysStep(), opts...)
	}
}

// ByAPIKeys orders the results by api_keys terms.
func ByAPIKeys(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAPIKeysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAPIKeysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(APIKeysInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, APIKeysTable, APIKeysColumn),
	)
}
