// Code generated by entc, DO NOT EDIT.

package category

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/api/src/infra/db/model/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Colour applies equality check predicate on the "colour" field. It's identical to ColourEQ.
func Colour(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColour), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// Admin applies equality check predicate on the "admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ColourEQ applies the EQ predicate on the "colour" field.
func ColourEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColour), v))
	})
}

// ColourNEQ applies the NEQ predicate on the "colour" field.
func ColourNEQ(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldColour), v))
	})
}

// ColourIn applies the In predicate on the "colour" field.
func ColourIn(vs ...string) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldColour), v...))
	})
}

// ColourNotIn applies the NotIn predicate on the "colour" field.
func ColourNotIn(vs ...string) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldColour), v...))
	})
}

// ColourGT applies the GT predicate on the "colour" field.
func ColourGT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldColour), v))
	})
}

// ColourGTE applies the GTE predicate on the "colour" field.
func ColourGTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldColour), v))
	})
}

// ColourLT applies the LT predicate on the "colour" field.
func ColourLT(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldColour), v))
	})
}

// ColourLTE applies the LTE predicate on the "colour" field.
func ColourLTE(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldColour), v))
	})
}

// ColourContains applies the Contains predicate on the "colour" field.
func ColourContains(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldColour), v))
	})
}

// ColourHasPrefix applies the HasPrefix predicate on the "colour" field.
func ColourHasPrefix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldColour), v))
	})
}

// ColourHasSuffix applies the HasSuffix predicate on the "colour" field.
func ColourHasSuffix(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldColour), v))
	})
}

// ColourEqualFold applies the EqualFold predicate on the "colour" field.
func ColourEqualFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldColour), v))
	})
}

// ColourContainsFold applies the ContainsFold predicate on the "colour" field.
func ColourContainsFold(v string) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldColour), v))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int) predicate.Category {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Category(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// AdminEQ applies the EQ predicate on the "admin" field.
func AdminEQ(v bool) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// AdminNEQ applies the NEQ predicate on the "admin" field.
func AdminNEQ(v bool) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdmin), v))
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		p(s.Not())
	})
}
