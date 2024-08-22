// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"iosync/ent/apikey"
	"iosync/ent/device"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ApiKey is the model entity for the ApiKey schema.
type ApiKey struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// RevokedAt holds the value of the "revoked_at" field.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// LastUsed holds the value of the "last_used" field.
	LastUsed time.Time `json:"last_used,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApiKeyQuery when eager-loading is set.
	Edges          ApiKeyEdges `json:"edges"`
	device_api_key *int
	selectValues   sql.SelectValues
}

// ApiKeyEdges holds the relations/edges for other nodes in the graph.
type ApiKeyEdges struct {
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// Topics holds the value of the topics edge.
	Topics []*Topic `json:"topics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApiKeyEdges) DeviceOrErr() (*Device, error) {
	if e.Device != nil {
		return e.Device, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: device.Label}
	}
	return nil, &NotLoadedError{edge: "device"}
}

// TopicsOrErr returns the Topics value or an error if the edge
// was not loaded in eager-loading.
func (e ApiKeyEdges) TopicsOrErr() ([]*Topic, error) {
	if e.loadedTypes[1] {
		return e.Topics, nil
	}
	return nil, &NotLoadedError{edge: "topics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApiKey) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apikey.FieldIsActive:
			values[i] = new(sql.NullBool)
		case apikey.FieldID:
			values[i] = new(sql.NullInt64)
		case apikey.FieldKey, apikey.FieldDescription:
			values[i] = new(sql.NullString)
		case apikey.FieldRevokedAt, apikey.FieldLastUsed, apikey.FieldCreatedAt, apikey.FieldUpdatedAt, apikey.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case apikey.ForeignKeys[0]: // device_api_key
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApiKey fields.
func (ak *ApiKey) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apikey.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ak.ID = int(value.Int64)
		case apikey.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				ak.Key = value.String
			}
		case apikey.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				ak.IsActive = value.Bool
			}
		case apikey.FieldRevokedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field revoked_at", values[i])
			} else if value.Valid {
				ak.RevokedAt = new(time.Time)
				*ak.RevokedAt = value.Time
			}
		case apikey.FieldLastUsed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used", values[i])
			} else if value.Valid {
				ak.LastUsed = value.Time
			}
		case apikey.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ak.CreatedAt = value.Time
			}
		case apikey.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ak.UpdatedAt = value.Time
			}
		case apikey.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				ak.ExpiresAt = new(time.Time)
				*ak.ExpiresAt = value.Time
			}
		case apikey.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ak.Description = new(string)
				*ak.Description = value.String
			}
		case apikey.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field device_api_key", value)
			} else if value.Valid {
				ak.device_api_key = new(int)
				*ak.device_api_key = int(value.Int64)
			}
		default:
			ak.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ApiKey.
// This includes values selected through modifiers, order, etc.
func (ak *ApiKey) Value(name string) (ent.Value, error) {
	return ak.selectValues.Get(name)
}

// QueryDevice queries the "device" edge of the ApiKey entity.
func (ak *ApiKey) QueryDevice() *DeviceQuery {
	return NewApiKeyClient(ak.config).QueryDevice(ak)
}

// QueryTopics queries the "topics" edge of the ApiKey entity.
func (ak *ApiKey) QueryTopics() *TopicQuery {
	return NewApiKeyClient(ak.config).QueryTopics(ak)
}

// Update returns a builder for updating this ApiKey.
// Note that you need to call ApiKey.Unwrap() before calling this method if this ApiKey
// was returned from a transaction, and the transaction was committed or rolled back.
func (ak *ApiKey) Update() *ApiKeyUpdateOne {
	return NewApiKeyClient(ak.config).UpdateOne(ak)
}

// Unwrap unwraps the ApiKey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ak *ApiKey) Unwrap() *ApiKey {
	_tx, ok := ak.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApiKey is not a transactional entity")
	}
	ak.config.driver = _tx.drv
	return ak
}

// String implements the fmt.Stringer.
func (ak *ApiKey) String() string {
	var builder strings.Builder
	builder.WriteString("ApiKey(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ak.ID))
	builder.WriteString("key=")
	builder.WriteString(ak.Key)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", ak.IsActive))
	builder.WriteString(", ")
	if v := ak.RevokedAt; v != nil {
		builder.WriteString("revoked_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("last_used=")
	builder.WriteString(ak.LastUsed.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ak.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ak.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ak.ExpiresAt; v != nil {
		builder.WriteString("expires_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ak.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// ApiKeys is a parsable slice of ApiKey.
type ApiKeys []*ApiKey
