// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/api/src/infra/db/model/notification"
	"github.com/Southclaws/storyden/api/src/infra/db/model/subscription"
	"github.com/Southclaws/storyden/api/src/infra/db/model/user"
	"github.com/google/uuid"
)

// SubscriptionCreate is the builder for creating a Subscription entity.
type SubscriptionCreate struct {
	config
	mutation *SubscriptionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRefersType sets the "refersType" field.
func (sc *SubscriptionCreate) SetRefersType(st subscription.RefersType) *SubscriptionCreate {
	sc.mutation.SetRefersType(st)
	return sc
}

// SetRefersTo sets the "refersTo" field.
func (sc *SubscriptionCreate) SetRefersTo(s string) *SubscriptionCreate {
	sc.mutation.SetRefersTo(s)
	return sc
}

// SetCreatedAt sets the "createdAt" field.
func (sc *SubscriptionCreate) SetCreatedAt(t time.Time) *SubscriptionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetUpdatedAt sets the "updatedAt" field.
func (sc *SubscriptionCreate) SetUpdatedAt(t time.Time) *SubscriptionCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetDeletedAt sets the "deletedAt" field.
func (sc *SubscriptionCreate) SetDeletedAt(t time.Time) *SubscriptionCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (sc *SubscriptionCreate) SetNillableDeletedAt(t *time.Time) *SubscriptionCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetUserId sets the "userId" field.
func (sc *SubscriptionCreate) SetUserId(s string) *SubscriptionCreate {
	sc.mutation.SetUserId(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SubscriptionCreate) SetID(s string) *SubscriptionCreate {
	sc.mutation.SetID(s)
	return sc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (sc *SubscriptionCreate) AddUserIDs(ids ...uuid.UUID) *SubscriptionCreate {
	sc.mutation.AddUserIDs(ids...)
	return sc
}

// AddUser adds the "user" edges to the User entity.
func (sc *SubscriptionCreate) AddUser(u ...*User) *SubscriptionCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddUserIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (sc *SubscriptionCreate) AddNotificationIDs(ids ...string) *SubscriptionCreate {
	sc.mutation.AddNotificationIDs(ids...)
	return sc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (sc *SubscriptionCreate) AddNotifications(n ...*Notification) *SubscriptionCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return sc.AddNotificationIDs(ids...)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (sc *SubscriptionCreate) Mutation() *SubscriptionMutation {
	return sc.mutation
}

// Save creates the Subscription in the database.
func (sc *SubscriptionCreate) Save(ctx context.Context) (*Subscription, error) {
	var (
		err  error
		node *Subscription
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubscriptionCreate) SaveX(ctx context.Context) *Subscription {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubscriptionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubscriptionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubscriptionCreate) check() error {
	if _, ok := sc.mutation.RefersType(); !ok {
		return &ValidationError{Name: "refersType", err: errors.New(`model: missing required field "Subscription.refersType"`)}
	}
	if v, ok := sc.mutation.RefersType(); ok {
		if err := subscription.RefersTypeValidator(v); err != nil {
			return &ValidationError{Name: "refersType", err: fmt.Errorf(`model: validator failed for field "Subscription.refersType": %w`, err)}
		}
	}
	if _, ok := sc.mutation.RefersTo(); !ok {
		return &ValidationError{Name: "refersTo", err: errors.New(`model: missing required field "Subscription.refersTo"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`model: missing required field "Subscription.createdAt"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`model: missing required field "Subscription.updatedAt"`)}
	}
	if _, ok := sc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`model: missing required field "Subscription.userId"`)}
	}
	return nil
}

func (sc *SubscriptionCreate) sqlSave(ctx context.Context) (*Subscription, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Subscription.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (sc *SubscriptionCreate) createSpec() (*Subscription, *sqlgraph.CreateSpec) {
	var (
		_node = &Subscription{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subscription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: subscription.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.RefersType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: subscription.FieldRefersType,
		})
		_node.RefersType = value
	}
	if value, ok := sc.mutation.RefersTo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldRefersTo,
		})
		_node.RefersTo = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscription.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscription.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.UserId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldUserId,
		})
		_node.UserId = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.NotificationsTable,
			Columns: []string{subscription.NotificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: notification.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subscription.Create().
//		SetRefersType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubscriptionUpsert) {
//			SetRefersType(v+v).
//		}).
//		Exec(ctx)
//
func (sc *SubscriptionCreate) OnConflict(opts ...sql.ConflictOption) *SubscriptionUpsertOne {
	sc.conflict = opts
	return &SubscriptionUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *SubscriptionCreate) OnConflictColumns(columns ...string) *SubscriptionUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SubscriptionUpsertOne{
		create: sc,
	}
}

type (
	// SubscriptionUpsertOne is the builder for "upsert"-ing
	//  one Subscription node.
	SubscriptionUpsertOne struct {
		create *SubscriptionCreate
	}

	// SubscriptionUpsert is the "OnConflict" setter.
	SubscriptionUpsert struct {
		*sql.UpdateSet
	}
)

// SetRefersType sets the "refersType" field.
func (u *SubscriptionUpsert) SetRefersType(v subscription.RefersType) *SubscriptionUpsert {
	u.Set(subscription.FieldRefersType, v)
	return u
}

// UpdateRefersType sets the "refersType" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateRefersType() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldRefersType)
	return u
}

// SetRefersTo sets the "refersTo" field.
func (u *SubscriptionUpsert) SetRefersTo(v string) *SubscriptionUpsert {
	u.Set(subscription.FieldRefersTo, v)
	return u
}

// UpdateRefersTo sets the "refersTo" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateRefersTo() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldRefersTo)
	return u
}

// SetCreatedAt sets the "createdAt" field.
func (u *SubscriptionUpsert) SetCreatedAt(v time.Time) *SubscriptionUpsert {
	u.Set(subscription.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateCreatedAt() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updatedAt" field.
func (u *SubscriptionUpsert) SetUpdatedAt(v time.Time) *SubscriptionUpsert {
	u.Set(subscription.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updatedAt" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateUpdatedAt() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deletedAt" field.
func (u *SubscriptionUpsert) SetDeletedAt(v time.Time) *SubscriptionUpsert {
	u.Set(subscription.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deletedAt" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateDeletedAt() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (u *SubscriptionUpsert) ClearDeletedAt() *SubscriptionUpsert {
	u.SetNull(subscription.FieldDeletedAt)
	return u
}

// SetUserId sets the "userId" field.
func (u *SubscriptionUpsert) SetUserId(v string) *SubscriptionUpsert {
	u.Set(subscription.FieldUserId, v)
	return u
}

// UpdateUserId sets the "userId" field to the value that was provided on create.
func (u *SubscriptionUpsert) UpdateUserId() *SubscriptionUpsert {
	u.SetExcluded(subscription.FieldUserId)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subscription.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *SubscriptionUpsertOne) UpdateNewValues() *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(subscription.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Subscription.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *SubscriptionUpsertOne) Ignore() *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubscriptionUpsertOne) DoNothing() *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubscriptionCreate.OnConflict
// documentation for more info.
func (u *SubscriptionUpsertOne) Update(set func(*SubscriptionUpsert)) *SubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubscriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetRefersType sets the "refersType" field.
func (u *SubscriptionUpsertOne) SetRefersType(v subscription.RefersType) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetRefersType(v)
	})
}

// UpdateRefersType sets the "refersType" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateRefersType() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateRefersType()
	})
}

// SetRefersTo sets the "refersTo" field.
func (u *SubscriptionUpsertOne) SetRefersTo(v string) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetRefersTo(v)
	})
}

// UpdateRefersTo sets the "refersTo" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateRefersTo() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateRefersTo()
	})
}

// SetCreatedAt sets the "createdAt" field.
func (u *SubscriptionUpsertOne) SetCreatedAt(v time.Time) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateCreatedAt() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updatedAt" field.
func (u *SubscriptionUpsertOne) SetUpdatedAt(v time.Time) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updatedAt" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateUpdatedAt() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deletedAt" field.
func (u *SubscriptionUpsertOne) SetDeletedAt(v time.Time) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deletedAt" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateDeletedAt() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (u *SubscriptionUpsertOne) ClearDeletedAt() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserId sets the "userId" field.
func (u *SubscriptionUpsertOne) SetUserId(v string) *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetUserId(v)
	})
}

// UpdateUserId sets the "userId" field to the value that was provided on create.
func (u *SubscriptionUpsertOne) UpdateUserId() *SubscriptionUpsertOne {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateUserId()
	})
}

// Exec executes the query.
func (u *SubscriptionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubscriptionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubscriptionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubscriptionUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: SubscriptionUpsertOne.ID is not supported by MySQL driver. Use SubscriptionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubscriptionUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubscriptionCreateBulk is the builder for creating many Subscription entities in bulk.
type SubscriptionCreateBulk struct {
	config
	builders []*SubscriptionCreate
	conflict []sql.ConflictOption
}

// Save creates the Subscription entities in the database.
func (scb *SubscriptionCreateBulk) Save(ctx context.Context) ([]*Subscription, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subscription, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriptionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubscriptionCreateBulk) SaveX(ctx context.Context) []*Subscription {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubscriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubscriptionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subscription.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubscriptionUpsert) {
//			SetRefersType(v+v).
//		}).
//		Exec(ctx)
//
func (scb *SubscriptionCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubscriptionUpsertBulk {
	scb.conflict = opts
	return &SubscriptionUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *SubscriptionCreateBulk) OnConflictColumns(columns ...string) *SubscriptionUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SubscriptionUpsertBulk{
		create: scb,
	}
}

// SubscriptionUpsertBulk is the builder for "upsert"-ing
// a bulk of Subscription nodes.
type SubscriptionUpsertBulk struct {
	create *SubscriptionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subscription.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *SubscriptionUpsertBulk) UpdateNewValues() *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(subscription.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Subscription.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *SubscriptionUpsertBulk) Ignore() *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubscriptionUpsertBulk) DoNothing() *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubscriptionCreateBulk.OnConflict
// documentation for more info.
func (u *SubscriptionUpsertBulk) Update(set func(*SubscriptionUpsert)) *SubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubscriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetRefersType sets the "refersType" field.
func (u *SubscriptionUpsertBulk) SetRefersType(v subscription.RefersType) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetRefersType(v)
	})
}

// UpdateRefersType sets the "refersType" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateRefersType() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateRefersType()
	})
}

// SetRefersTo sets the "refersTo" field.
func (u *SubscriptionUpsertBulk) SetRefersTo(v string) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetRefersTo(v)
	})
}

// UpdateRefersTo sets the "refersTo" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateRefersTo() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateRefersTo()
	})
}

// SetCreatedAt sets the "createdAt" field.
func (u *SubscriptionUpsertBulk) SetCreatedAt(v time.Time) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateCreatedAt() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updatedAt" field.
func (u *SubscriptionUpsertBulk) SetUpdatedAt(v time.Time) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updatedAt" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateUpdatedAt() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deletedAt" field.
func (u *SubscriptionUpsertBulk) SetDeletedAt(v time.Time) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deletedAt" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateDeletedAt() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (u *SubscriptionUpsertBulk) ClearDeletedAt() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserId sets the "userId" field.
func (u *SubscriptionUpsertBulk) SetUserId(v string) *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.SetUserId(v)
	})
}

// UpdateUserId sets the "userId" field to the value that was provided on create.
func (u *SubscriptionUpsertBulk) UpdateUserId() *SubscriptionUpsertBulk {
	return u.Update(func(s *SubscriptionUpsert) {
		s.UpdateUserId()
	})
}

// Exec executes the query.
func (u *SubscriptionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SubscriptionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubscriptionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubscriptionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
