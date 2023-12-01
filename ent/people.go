// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"testEntGo/ent/people"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// People is the model entity for the People schema.
type People struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Age holds the value of the "age" field.
	Age          int `json:"age,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*People) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case people.FieldID, people.FieldAge:
			values[i] = new(sql.NullInt64)
		case people.FieldName, people.FieldLastName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the People fields.
func (pe *People) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case people.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case people.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case people.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				pe.LastName = value.String
			}
		case people.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pe.Age = int(value.Int64)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the People.
// This includes values selected through modifiers, order, etc.
func (pe *People) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// Update returns a builder for updating this People.
// Note that you need to call People.Unwrap() before calling this method if this People
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *People) Update() *PeopleUpdateOne {
	return NewPeopleClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the People entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *People) Unwrap() *People {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: People is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *People) String() string {
	var builder strings.Builder
	builder.WriteString("People(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(pe.LastName)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", pe.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Peoples is a parsable slice of People.
type Peoples []*People