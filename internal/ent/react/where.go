// Code generated by ent, DO NOT EDIT.

package react

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.React {
	return predicate.React(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.React {
	return predicate.React(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.React {
	return predicate.React(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.React {
	return predicate.React(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.React {
	return predicate.React(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.React {
	return predicate.React(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.React {
	return predicate.React(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.React {
	return predicate.React(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.React {
	return predicate.React(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.React {
	return predicate.React(sql.FieldEQ(FieldCreatedAt, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v xid.ID) predicate.React {
	return predicate.React(sql.FieldEQ(FieldAccountID, v))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v xid.ID) predicate.React {
	return predicate.React(sql.FieldEQ(FieldPostID, v))
}

// Emoji applies equality check predicate on the "emoji" field. It's identical to EmojiEQ.
func Emoji(v string) predicate.React {
	return predicate.React(sql.FieldEQ(FieldEmoji, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.React {
	return predicate.React(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.React {
	return predicate.React(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.React {
	return predicate.React(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.React {
	return predicate.React(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.React {
	return predicate.React(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.React {
	return predicate.React(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.React {
	return predicate.React(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.React {
	return predicate.React(sql.FieldLTE(FieldCreatedAt, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v xid.ID) predicate.React {
	return predicate.React(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v xid.ID) predicate.React {
	return predicate.React(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...xid.ID) predicate.React {
	return predicate.React(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...xid.ID) predicate.React {
	return predicate.React(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v xid.ID) predicate.React {
	return predicate.React(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v xid.ID) predicate.React {
	return predicate.React(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v xid.ID) predicate.React {
	return predicate.React(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v xid.ID) predicate.React {
	return predicate.React(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldContains(FieldAccountID, vc))
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldHasPrefix(FieldAccountID, vc))
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldHasSuffix(FieldAccountID, vc))
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldEqualFold(FieldAccountID, vc))
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldContainsFold(FieldAccountID, vc))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v xid.ID) predicate.React {
	return predicate.React(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v xid.ID) predicate.React {
	return predicate.React(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...xid.ID) predicate.React {
	return predicate.React(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...xid.ID) predicate.React {
	return predicate.React(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v xid.ID) predicate.React {
	return predicate.React(sql.FieldGT(FieldPostID, v))
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v xid.ID) predicate.React {
	return predicate.React(sql.FieldGTE(FieldPostID, v))
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v xid.ID) predicate.React {
	return predicate.React(sql.FieldLT(FieldPostID, v))
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v xid.ID) predicate.React {
	return predicate.React(sql.FieldLTE(FieldPostID, v))
}

// PostIDContains applies the Contains predicate on the "post_id" field.
func PostIDContains(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldContains(FieldPostID, vc))
}

// PostIDHasPrefix applies the HasPrefix predicate on the "post_id" field.
func PostIDHasPrefix(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldHasPrefix(FieldPostID, vc))
}

// PostIDHasSuffix applies the HasSuffix predicate on the "post_id" field.
func PostIDHasSuffix(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldHasSuffix(FieldPostID, vc))
}

// PostIDEqualFold applies the EqualFold predicate on the "post_id" field.
func PostIDEqualFold(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldEqualFold(FieldPostID, vc))
}

// PostIDContainsFold applies the ContainsFold predicate on the "post_id" field.
func PostIDContainsFold(v xid.ID) predicate.React {
	vc := v.String()
	return predicate.React(sql.FieldContainsFold(FieldPostID, vc))
}

// EmojiEQ applies the EQ predicate on the "emoji" field.
func EmojiEQ(v string) predicate.React {
	return predicate.React(sql.FieldEQ(FieldEmoji, v))
}

// EmojiNEQ applies the NEQ predicate on the "emoji" field.
func EmojiNEQ(v string) predicate.React {
	return predicate.React(sql.FieldNEQ(FieldEmoji, v))
}

// EmojiIn applies the In predicate on the "emoji" field.
func EmojiIn(vs ...string) predicate.React {
	return predicate.React(sql.FieldIn(FieldEmoji, vs...))
}

// EmojiNotIn applies the NotIn predicate on the "emoji" field.
func EmojiNotIn(vs ...string) predicate.React {
	return predicate.React(sql.FieldNotIn(FieldEmoji, vs...))
}

// EmojiGT applies the GT predicate on the "emoji" field.
func EmojiGT(v string) predicate.React {
	return predicate.React(sql.FieldGT(FieldEmoji, v))
}

// EmojiGTE applies the GTE predicate on the "emoji" field.
func EmojiGTE(v string) predicate.React {
	return predicate.React(sql.FieldGTE(FieldEmoji, v))
}

// EmojiLT applies the LT predicate on the "emoji" field.
func EmojiLT(v string) predicate.React {
	return predicate.React(sql.FieldLT(FieldEmoji, v))
}

// EmojiLTE applies the LTE predicate on the "emoji" field.
func EmojiLTE(v string) predicate.React {
	return predicate.React(sql.FieldLTE(FieldEmoji, v))
}

// EmojiContains applies the Contains predicate on the "emoji" field.
func EmojiContains(v string) predicate.React {
	return predicate.React(sql.FieldContains(FieldEmoji, v))
}

// EmojiHasPrefix applies the HasPrefix predicate on the "emoji" field.
func EmojiHasPrefix(v string) predicate.React {
	return predicate.React(sql.FieldHasPrefix(FieldEmoji, v))
}

// EmojiHasSuffix applies the HasSuffix predicate on the "emoji" field.
func EmojiHasSuffix(v string) predicate.React {
	return predicate.React(sql.FieldHasSuffix(FieldEmoji, v))
}

// EmojiEqualFold applies the EqualFold predicate on the "emoji" field.
func EmojiEqualFold(v string) predicate.React {
	return predicate.React(sql.FieldEqualFold(FieldEmoji, v))
}

// EmojiContainsFold applies the ContainsFold predicate on the "emoji" field.
func EmojiContainsFold(v string) predicate.React {
	return predicate.React(sql.FieldContainsFold(FieldEmoji, v))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.React {
	return predicate.React(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.React {
	return predicate.React(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPost applies the HasEdge predicate on the "Post" edge.
func HasPost() predicate.React {
	return predicate.React(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "Post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.React {
	return predicate.React(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.React) predicate.React {
	return predicate.React(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.React) predicate.React {
	return predicate.React(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.React) predicate.React {
	return predicate.React(sql.NotPredicates(p))
}
