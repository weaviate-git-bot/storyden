// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/notification"
	"github.com/Southclaws/storyden/internal/ent/predicate"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks     []Hook
	mutation  *NotificationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetTitle sets the "title" field.
func (nu *NotificationUpdate) SetTitle(s string) *NotificationUpdate {
	nu.mutation.SetTitle(s)
	return nu
}

// SetDescription sets the "description" field.
func (nu *NotificationUpdate) SetDescription(s string) *NotificationUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetLink sets the "link" field.
func (nu *NotificationUpdate) SetLink(s string) *NotificationUpdate {
	nu.mutation.SetLink(s)
	return nu
}

// SetRead sets the "read" field.
func (nu *NotificationUpdate) SetRead(b bool) *NotificationUpdate {
	nu.mutation.SetRead(b)
	return nu
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nu *NotificationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotificationUpdate {
	nu.modifiers = append(nu.modifiers, modifiers...)
	return nu
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.SetField(notification.FieldDescription, field.TypeString, value)
	}
	if value, ok := nu.mutation.Link(); ok {
		_spec.SetField(notification.FieldLink, field.TypeString, value)
	}
	if value, ok := nu.mutation.Read(); ok {
		_spec.SetField(notification.FieldRead, field.TypeBool, value)
	}
	_spec.AddModifiers(nu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *NotificationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTitle sets the "title" field.
func (nuo *NotificationUpdateOne) SetTitle(s string) *NotificationUpdateOne {
	nuo.mutation.SetTitle(s)
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NotificationUpdateOne) SetDescription(s string) *NotificationUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetLink sets the "link" field.
func (nuo *NotificationUpdateOne) SetLink(s string) *NotificationUpdateOne {
	nuo.mutation.SetLink(s)
	return nuo
}

// SetRead sets the "read" field.
func (nuo *NotificationUpdateOne) SetRead(b bool) *NotificationUpdateOne {
	nuo.mutation.SetRead(b)
	return nuo
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nuo *NotificationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotificationUpdateOne {
	nuo.modifiers = append(nuo.modifiers, modifiers...)
	return nuo
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.SetField(notification.FieldDescription, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Link(); ok {
		_spec.SetField(notification.FieldLink, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Read(); ok {
		_spec.SetField(notification.FieldRead, field.TypeBool, value)
	}
	_spec.AddModifiers(nuo.modifiers...)
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
