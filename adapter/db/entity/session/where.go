// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lstrihic/webapp/adapter/db/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IsValid applies equality check predicate on the "is_valid" field. It's identical to IsValidEQ.
func IsValid(v bool) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsValid), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPIsNil applies the IsNil predicate on the "ip" field.
func IPIsNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIP)))
	})
}

// IPNotNil applies the NotNil predicate on the "ip" field.
func IPNotNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIP)))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// IsValidEQ applies the EQ predicate on the "is_valid" field.
func IsValidEQ(v bool) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsValid), v))
	})
}

// IsValidNEQ applies the NEQ predicate on the "is_valid" field.
func IsValidNEQ(v bool) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsValid), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
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
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		p(s.Not())
	})
}
