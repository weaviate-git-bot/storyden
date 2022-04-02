// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/api/src/infra/db/model/subscription"
	"github.com/google/uuid"
)

// Subscription is the model entity for the Subscription schema.
type Subscription struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// RefersType holds the value of the "refersType" field.
	RefersType subscription.RefersType `json:"refersType,omitempty"`
	// RefersTo holds the value of the "refersTo" field.
	RefersTo string `json:"refersTo,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId string `json:"userId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubscriptionQuery when eager-loading is set.
	Edges                     SubscriptionEdges `json:"edges"`
	notification_subscription *string
	user_subscriptions        *uuid.UUID
}

// SubscriptionEdges holds the relations/edges for other nodes in the graph.
type SubscriptionEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e SubscriptionEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e SubscriptionEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[1] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscription) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscription.FieldID, subscription.FieldRefersType, subscription.FieldRefersTo, subscription.FieldUserId:
			values[i] = new(sql.NullString)
		case subscription.FieldCreatedAt, subscription.FieldUpdatedAt, subscription.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case subscription.ForeignKeys[0]: // notification_subscription
			values[i] = new(sql.NullString)
		case subscription.ForeignKeys[1]: // user_subscriptions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subscription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscription fields.
func (s *Subscription) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscription.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case subscription.FieldRefersType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refersType", values[i])
			} else if value.Valid {
				s.RefersType = subscription.RefersType(value.String)
			}
		case subscription.FieldRefersTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refersTo", values[i])
			} else if value.Valid {
				s.RefersTo = value.String
			}
		case subscription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case subscription.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case subscription.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case subscription.FieldUserId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				s.UserId = value.String
			}
		case subscription.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notification_subscription", values[i])
			} else if value.Valid {
				s.notification_subscription = new(string)
				*s.notification_subscription = value.String
			}
		case subscription.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_subscriptions", values[i])
			} else if value.Valid {
				s.user_subscriptions = new(uuid.UUID)
				*s.user_subscriptions = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Subscription entity.
func (s *Subscription) QueryUser() *UserQuery {
	return (&SubscriptionClient{config: s.config}).QueryUser(s)
}

// QueryNotifications queries the "notifications" edge of the Subscription entity.
func (s *Subscription) QueryNotifications() *NotificationQuery {
	return (&SubscriptionClient{config: s.config}).QueryNotifications(s)
}

// Update returns a builder for updating this Subscription.
// Note that you need to call Subscription.Unwrap() before calling this method if this Subscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscription) Update() *SubscriptionUpdateOne {
	return (&SubscriptionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscription) Unwrap() *Subscription {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("model: Subscription is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscription) String() string {
	var builder strings.Builder
	builder.WriteString("Subscription(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", refersType=")
	builder.WriteString(fmt.Sprintf("%v", s.RefersType))
	builder.WriteString(", refersTo=")
	builder.WriteString(s.RefersTo)
	builder.WriteString(", createdAt=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deletedAt=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", userId=")
	builder.WriteString(s.UserId)
	builder.WriteByte(')')
	return builder.String()
}

// Subscriptions is a parsable slice of Subscription.
type Subscriptions []*Subscription

func (s Subscriptions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
