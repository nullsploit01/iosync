// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "is_online", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
	}
	// NodeAPIKeysColumns holds the columns for the "node_api_keys" table.
	NodeAPIKeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "api_key", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_revoked", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "node_api_keys", Type: field.TypeInt, Nullable: true},
	}
	// NodeAPIKeysTable holds the schema information for the "node_api_keys" table.
	NodeAPIKeysTable = &schema.Table{
		Name:       "node_api_keys",
		Columns:    NodeAPIKeysColumns,
		PrimaryKey: []*schema.Column{NodeAPIKeysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_api_keys_nodes_api_keys",
				Columns:    []*schema.Column{NodeAPIKeysColumns[6]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NodeValuesColumns holds the columns for the "node_values" table.
	NodeValuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "node_api_key_values", Type: field.TypeInt, Nullable: true},
	}
	// NodeValuesTable holds the schema information for the "node_values" table.
	NodeValuesTable = &schema.Table{
		Name:       "node_values",
		Columns:    NodeValuesColumns,
		PrimaryKey: []*schema.Column{NodeValuesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_values_node_api_keys_values",
				Columns:    []*schema.Column{NodeValuesColumns[4]},
				RefColumns: []*schema.Column{NodeAPIKeysColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NodesTable,
		NodeAPIKeysTable,
		NodeValuesTable,
	}
)

func init() {
	NodeAPIKeysTable.ForeignKeys[0].RefTable = NodesTable
	NodeValuesTable.ForeignKeys[0].RefTable = NodeAPIKeysTable
}
