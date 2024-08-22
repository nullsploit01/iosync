// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APIKeysColumns holds the columns for the "api_keys" table.
	APIKeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "revoked_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_used", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "device_api_key", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// APIKeysTable holds the schema information for the "api_keys" table.
	APIKeysTable = &schema.Table{
		Name:       "api_keys",
		Columns:    APIKeysColumns,
		PrimaryKey: []*schema.Column{APIKeysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "api_keys_devices_api_key",
				Columns:    []*schema.Column{APIKeysColumns[9]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_devices", Type: field.TypeInt, Nullable: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_users_devices",
				Columns:    []*schema.Column{DevicesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "session_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "user_sessions", Type: field.TypeInt, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TopicsColumns holds the columns for the "topics" table.
	TopicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "qos", Type: field.TypeInt, Default: 0},
		{Name: "retain", Type: field.TypeBool, Default: false},
		{Name: "last_used", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "device_topics", Type: field.TypeInt, Nullable: true},
	}
	// TopicsTable holds the schema information for the "topics" table.
	TopicsTable = &schema.Table{
		Name:       "topics",
		Columns:    TopicsColumns,
		PrimaryKey: []*schema.Column{TopicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "topics_devices_topics",
				Columns:    []*schema.Column{TopicsColumns[8]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "last_login", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APIKeysTable,
		DevicesTable,
		SessionsTable,
		TopicsTable,
		UsersTable,
	}
)

func init() {
	APIKeysTable.ForeignKeys[0].RefTable = DevicesTable
	DevicesTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	TopicsTable.ForeignKeys[0].RefTable = DevicesTable
}
