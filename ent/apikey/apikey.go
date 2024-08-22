// Code generated by ent, DO NOT EDIT.

package apikey

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the apikey type in the database.
	Label = "api_key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldRevokedAt holds the string denoting the revoked_at field in the database.
	FieldRevokedAt = "revoked_at"
	// FieldLastUsed holds the string denoting the last_used field in the database.
	FieldLastUsed = "last_used"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// EdgeTopics holds the string denoting the topics edge name in mutations.
	EdgeTopics = "topics"
	// Table holds the table name of the apikey in the database.
	Table = "api_keys"
	// DeviceTable is the table that holds the device relation/edge.
	DeviceTable = "api_keys"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "device_api_key"
	// TopicsTable is the table that holds the topics relation/edge.
	TopicsTable = "topics"
	// TopicsInverseTable is the table name for the Topic entity.
	// It exists in this package in order to avoid circular dependency with the "topic" package.
	TopicsInverseTable = "topics"
	// TopicsColumn is the table column denoting the topics relation/edge.
	TopicsColumn = "api_key_topics"
)

// Columns holds all SQL columns for apikey fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldIsActive,
	FieldRevokedAt,
	FieldLastUsed,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldExpiresAt,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "api_keys"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_api_key",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultLastUsed holds the default value on creation for the "last_used" field.
	DefaultLastUsed func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the ApiKey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByRevokedAt orders the results by the revoked_at field.
func ByRevokedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevokedAt, opts...).ToFunc()
}

// ByLastUsed orders the results by the last_used field.
func ByLastUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUsed, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByDeviceField orders the results by device field.
func ByDeviceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceStep(), sql.OrderByField(field, opts...))
	}
}

// ByTopicsCount orders the results by topics count.
func ByTopicsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTopicsStep(), opts...)
	}
}

// ByTopics orders the results by topics terms.
func ByTopics(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTopicsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDeviceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, DeviceTable, DeviceColumn),
	)
}
func newTopicsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TopicsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TopicsTable, TopicsColumn),
	)
}
