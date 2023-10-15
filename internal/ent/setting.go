// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/ent/setting"
)

// Setting is the model entity for the Setting schema.
type Setting struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Setting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case setting.FieldID, setting.FieldValue:
			values[i] = new(sql.NullString)
		case setting.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Setting fields.
func (s *Setting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case setting.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case setting.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = value.String
			}
		case setting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Setting.
// This includes values selected through modifiers, order, etc.
func (s *Setting) GetValue(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Setting.
// Note that you need to call Setting.Unwrap() before calling this method if this Setting
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Setting) Update() *SettingUpdateOne {
	return NewSettingClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Setting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Setting) Unwrap() *Setting {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Setting is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Setting) String() string {
	var builder strings.Builder
	builder.WriteString("Setting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("value=")
	builder.WriteString(s.Value)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Settings is a parsable slice of Setting.
type Settings []*Setting
