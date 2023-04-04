// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package node

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/o2orecur/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldID, id))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldValue, v))
}

// PrevID applies equality check predicate on the "prev_id" field. It's identical to PrevIDEQ.
func PrevID(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldPrevID, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldValue, v))
}

// PrevIDEQ applies the EQ predicate on the "prev_id" field.
func PrevIDEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldPrevID, v))
}

// PrevIDNEQ applies the NEQ predicate on the "prev_id" field.
func PrevIDNEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldPrevID, v))
}

// PrevIDIn applies the In predicate on the "prev_id" field.
func PrevIDIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldPrevID, vs...))
}

// PrevIDNotIn applies the NotIn predicate on the "prev_id" field.
func PrevIDNotIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldPrevID, vs...))
}

// PrevIDIsNil applies the IsNil predicate on the "prev_id" field.
func PrevIDIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldPrevID))
}

// PrevIDNotNil applies the NotNil predicate on the "prev_id" field.
func PrevIDNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldPrevID))
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
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
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		p(s.Not())
	})
}
