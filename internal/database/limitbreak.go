// Code generated by entc, DO NOT EDIT.

package database

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/iamnande/cardmod/internal/database/limitbreak"
)

// LimitBreak is the model entity for the LimitBreak schema.
type LimitBreak struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LimitBreak) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case limitbreak.FieldID:
			values[i] = new(sql.NullInt64)
		case limitbreak.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LimitBreak", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LimitBreak fields.
func (lb *LimitBreak) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case limitbreak.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lb.ID = int(value.Int64)
		case limitbreak.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				lb.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this LimitBreak.
// Note that you need to call LimitBreak.Unwrap() before calling this method if this LimitBreak
// was returned from a transaction, and the transaction was committed or rolled back.
func (lb *LimitBreak) Update() *LimitBreakUpdateOne {
	return (&LimitBreakClient{config: lb.config}).UpdateOne(lb)
}

// Unwrap unwraps the LimitBreak entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lb *LimitBreak) Unwrap() *LimitBreak {
	tx, ok := lb.config.driver.(*txDriver)
	if !ok {
		panic("database: LimitBreak is not a transactional entity")
	}
	lb.config.driver = tx.drv
	return lb
}

// String implements the fmt.Stringer.
func (lb *LimitBreak) String() string {
	var builder strings.Builder
	builder.WriteString("LimitBreak(")
	builder.WriteString(fmt.Sprintf("id=%v", lb.ID))
	builder.WriteString(", name=")
	builder.WriteString(lb.Name)
	builder.WriteByte(')')
	return builder.String()
}

// LimitBreaks is a parsable slice of LimitBreak.
type LimitBreaks []*LimitBreak

func (lb LimitBreaks) config(cfg config) {
	for _i := range lb {
		lb[_i].config = cfg
	}
}
