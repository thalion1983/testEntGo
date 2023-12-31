// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"testEntGo/ent/clothe"
	"testEntGo/ent/people"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Clothe is the model entity for the Clothe schema.
type Clothe struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// BuyDate holds the value of the "buy_date" field.
	BuyDate time.Time `json:"buy_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClotheQuery when eager-loading is set.
	Edges          ClotheEdges `json:"edges"`
	people_clothes *int
	selectValues   sql.SelectValues
}

// ClotheEdges holds the relations/edges for other nodes in the graph.
type ClotheEdges struct {
	// Owner holds the value of the owner edge.
	Owner *People `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClotheEdges) OwnerOrErr() (*People, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: people.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Clothe) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case clothe.FieldID:
			values[i] = new(sql.NullInt64)
		case clothe.FieldType, clothe.FieldColor:
			values[i] = new(sql.NullString)
		case clothe.FieldBuyDate:
			values[i] = new(sql.NullTime)
		case clothe.ForeignKeys[0]: // people_clothes
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Clothe fields.
func (c *Clothe) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case clothe.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case clothe.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case clothe.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				c.Color = value.String
			}
		case clothe.FieldBuyDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field buy_date", values[i])
			} else if value.Valid {
				c.BuyDate = value.Time
			}
		case clothe.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field people_clothes", value)
			} else if value.Valid {
				c.people_clothes = new(int)
				*c.people_clothes = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Clothe.
// This includes values selected through modifiers, order, etc.
func (c *Clothe) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Clothe entity.
func (c *Clothe) QueryOwner() *PeopleQuery {
	return NewClotheClient(c.config).QueryOwner(c)
}

// Update returns a builder for updating this Clothe.
// Note that you need to call Clothe.Unwrap() before calling this method if this Clothe
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Clothe) Update() *ClotheUpdateOne {
	return NewClotheClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Clothe entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Clothe) Unwrap() *Clothe {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Clothe is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Clothe) String() string {
	var builder strings.Builder
	builder.WriteString("Clothe(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(c.Color)
	builder.WriteString(", ")
	builder.WriteString("buy_date=")
	builder.WriteString(c.BuyDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Clothes is a parsable slice of Clothe.
type Clothes []*Clothe
