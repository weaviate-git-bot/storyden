// Code generated by ent, DO NOT EDIT.

package notification

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the notification type in the database.
	Label = "notification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldRead holds the string denoting the read field in the database.
	FieldRead = "read"
	// Table holds the table name of the notification in the database.
	Table = "notifications"
)

// Columns holds all SQL columns for notification fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldTitle,
	FieldDescription,
	FieldLink,
	FieldRead,
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Notification queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByLink orders the results by the link field.
func ByLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLink, opts...).ToFunc()
}

// ByRead orders the results by the read field.
func ByRead(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRead, opts...).ToFunc()
}
