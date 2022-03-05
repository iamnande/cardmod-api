// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "level", Type: field.TypeInt32},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
	}
	// LimitBreaksColumns holds the columns for the "limit_breaks" table.
	LimitBreaksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// LimitBreaksTable holds the schema information for the "limit_breaks" table.
	LimitBreaksTable = &schema.Table{
		Name:       "limit_breaks",
		Columns:    LimitBreaksColumns,
		PrimaryKey: []*schema.Column{LimitBreaksColumns[0]},
	}
	// MagicsColumns holds the columns for the "magics" table.
	MagicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "purpose", Type: field.TypeEnum, Enums: []string{"Offensive", "Restorative", "Indirect"}},
	}
	// MagicsTable holds the schema information for the "magics" table.
	MagicsTable = &schema.Table{
		Name:       "magics",
		Columns:    MagicsColumns,
		PrimaryKey: []*schema.Column{MagicsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		ItemsTable,
		LimitBreaksTable,
		MagicsTable,
	}
)

func init() {
}
