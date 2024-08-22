// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"iosync/ent/apikey"
	"iosync/ent/device"
	"iosync/ent/topic"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Topic is the model entity for the Topic schema.
type Topic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Qos holds the value of the "qos" field.
	Qos int `json:"qos,omitempty"`
	// Retain holds the value of the "retain" field.
	Retain bool `json:"retain,omitempty"`
	// LastUsed holds the value of the "last_used" field.
	LastUsed time.Time `json:"last_used,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopicQuery when eager-loading is set.
	Edges          TopicEdges `json:"edges"`
	api_key_topics *int
	device_topics  *int
	selectValues   sql.SelectValues
}

// TopicEdges holds the relations/edges for other nodes in the graph.
type TopicEdges struct {
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// APIKey holds the value of the api_key edge.
	APIKey *ApiKey `json:"api_key,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TopicEdges) DeviceOrErr() (*Device, error) {
	if e.Device != nil {
		return e.Device, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: device.Label}
	}
	return nil, &NotLoadedError{edge: "device"}
}

// APIKeyOrErr returns the APIKey value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TopicEdges) APIKeyOrErr() (*ApiKey, error) {
	if e.APIKey != nil {
		return e.APIKey, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: apikey.Label}
	}
	return nil, &NotLoadedError{edge: "api_key"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Topic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case topic.FieldIsActive, topic.FieldRetain:
			values[i] = new(sql.NullBool)
		case topic.FieldID, topic.FieldQos:
			values[i] = new(sql.NullInt64)
		case topic.FieldName:
			values[i] = new(sql.NullString)
		case topic.FieldLastUsed, topic.FieldCreatedAt, topic.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case topic.ForeignKeys[0]: // api_key_topics
			values[i] = new(sql.NullInt64)
		case topic.ForeignKeys[1]: // device_topics
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Topic fields.
func (t *Topic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case topic.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case topic.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				t.IsActive = value.Bool
			}
		case topic.FieldQos:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field qos", values[i])
			} else if value.Valid {
				t.Qos = int(value.Int64)
			}
		case topic.FieldRetain:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field retain", values[i])
			} else if value.Valid {
				t.Retain = value.Bool
			}
		case topic.FieldLastUsed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used", values[i])
			} else if value.Valid {
				t.LastUsed = value.Time
			}
		case topic.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case topic.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case topic.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field api_key_topics", value)
			} else if value.Valid {
				t.api_key_topics = new(int)
				*t.api_key_topics = int(value.Int64)
			}
		case topic.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field device_topics", value)
			} else if value.Valid {
				t.device_topics = new(int)
				*t.device_topics = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Topic.
// This includes values selected through modifiers, order, etc.
func (t *Topic) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryDevice queries the "device" edge of the Topic entity.
func (t *Topic) QueryDevice() *DeviceQuery {
	return NewTopicClient(t.config).QueryDevice(t)
}

// QueryAPIKey queries the "api_key" edge of the Topic entity.
func (t *Topic) QueryAPIKey() *ApiKeyQuery {
	return NewTopicClient(t.config).QueryAPIKey(t)
}

// Update returns a builder for updating this Topic.
// Note that you need to call Topic.Unwrap() before calling this method if this Topic
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Topic) Update() *TopicUpdateOne {
	return NewTopicClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Topic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Topic) Unwrap() *Topic {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Topic is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Topic) String() string {
	var builder strings.Builder
	builder.WriteString("Topic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", t.IsActive))
	builder.WriteString(", ")
	builder.WriteString("qos=")
	builder.WriteString(fmt.Sprintf("%v", t.Qos))
	builder.WriteString(", ")
	builder.WriteString("retain=")
	builder.WriteString(fmt.Sprintf("%v", t.Retain))
	builder.WriteString(", ")
	builder.WriteString("last_used=")
	builder.WriteString(t.LastUsed.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Topics is a parsable slice of Topic.
type Topics []*Topic
