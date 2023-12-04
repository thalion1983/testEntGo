// Code generated by ent, DO NOT EDIT.

package group

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgePeoples holds the string denoting the peoples edge name in mutations.
	EdgePeoples = "peoples"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// PeoplesTable is the table that holds the peoples relation/edge. The primary key declared below.
	PeoplesTable = "group_peoples"
	// PeoplesInverseTable is the table name for the People entity.
	// It exists in this package in order to avoid circular dependency with the "people" package.
	PeoplesInverseTable = "peoples"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// PeoplesPrimaryKey and PeoplesColumn2 are the table columns denoting the
	// primary key for the peoples relation (M2M).
	PeoplesPrimaryKey = []string{"group_id", "people_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Group queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPeoplesCount orders the results by peoples count.
func ByPeoplesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPeoplesStep(), opts...)
	}
}

// ByPeoples orders the results by peoples terms.
func ByPeoples(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeoplesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPeoplesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PeoplesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PeoplesTable, PeoplesPrimaryKey...),
	)
}