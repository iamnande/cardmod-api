// Code generated by entc, DO NOT EDIT.

package refinement

import (
	"entgo.io/ent/dialect/sql"
	"github.com/iamnande/cardmod/internal/database/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// Target applies equality check predicate on the "target" field. It's identical to TargetEQ.
func Target(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTarget), v))
	})
}

// Numerator applies equality check predicate on the "numerator" field. It's identical to NumeratorEQ.
func Numerator(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumerator), v))
	})
}

// Denominator applies equality check predicate on the "denominator" field. It's identical to DenominatorEQ.
func Denominator(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDenominator), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// TargetEQ applies the EQ predicate on the "target" field.
func TargetEQ(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTarget), v))
	})
}

// TargetNEQ applies the NEQ predicate on the "target" field.
func TargetNEQ(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTarget), v))
	})
}

// TargetIn applies the In predicate on the "target" field.
func TargetIn(vs ...string) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTarget), v...))
	})
}

// TargetNotIn applies the NotIn predicate on the "target" field.
func TargetNotIn(vs ...string) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTarget), v...))
	})
}

// TargetGT applies the GT predicate on the "target" field.
func TargetGT(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTarget), v))
	})
}

// TargetGTE applies the GTE predicate on the "target" field.
func TargetGTE(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTarget), v))
	})
}

// TargetLT applies the LT predicate on the "target" field.
func TargetLT(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTarget), v))
	})
}

// TargetLTE applies the LTE predicate on the "target" field.
func TargetLTE(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTarget), v))
	})
}

// TargetContains applies the Contains predicate on the "target" field.
func TargetContains(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTarget), v))
	})
}

// TargetHasPrefix applies the HasPrefix predicate on the "target" field.
func TargetHasPrefix(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTarget), v))
	})
}

// TargetHasSuffix applies the HasSuffix predicate on the "target" field.
func TargetHasSuffix(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTarget), v))
	})
}

// TargetEqualFold applies the EqualFold predicate on the "target" field.
func TargetEqualFold(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTarget), v))
	})
}

// TargetContainsFold applies the ContainsFold predicate on the "target" field.
func TargetContainsFold(v string) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTarget), v))
	})
}

// NumeratorEQ applies the EQ predicate on the "numerator" field.
func NumeratorEQ(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumerator), v))
	})
}

// NumeratorNEQ applies the NEQ predicate on the "numerator" field.
func NumeratorNEQ(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumerator), v))
	})
}

// NumeratorIn applies the In predicate on the "numerator" field.
func NumeratorIn(vs ...int32) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumerator), v...))
	})
}

// NumeratorNotIn applies the NotIn predicate on the "numerator" field.
func NumeratorNotIn(vs ...int32) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumerator), v...))
	})
}

// NumeratorGT applies the GT predicate on the "numerator" field.
func NumeratorGT(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumerator), v))
	})
}

// NumeratorGTE applies the GTE predicate on the "numerator" field.
func NumeratorGTE(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumerator), v))
	})
}

// NumeratorLT applies the LT predicate on the "numerator" field.
func NumeratorLT(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumerator), v))
	})
}

// NumeratorLTE applies the LTE predicate on the "numerator" field.
func NumeratorLTE(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumerator), v))
	})
}

// DenominatorEQ applies the EQ predicate on the "denominator" field.
func DenominatorEQ(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDenominator), v))
	})
}

// DenominatorNEQ applies the NEQ predicate on the "denominator" field.
func DenominatorNEQ(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDenominator), v))
	})
}

// DenominatorIn applies the In predicate on the "denominator" field.
func DenominatorIn(vs ...int32) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDenominator), v...))
	})
}

// DenominatorNotIn applies the NotIn predicate on the "denominator" field.
func DenominatorNotIn(vs ...int32) predicate.Refinement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Refinement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDenominator), v...))
	})
}

// DenominatorGT applies the GT predicate on the "denominator" field.
func DenominatorGT(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDenominator), v))
	})
}

// DenominatorGTE applies the GTE predicate on the "denominator" field.
func DenominatorGTE(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDenominator), v))
	})
}

// DenominatorLT applies the LT predicate on the "denominator" field.
func DenominatorLT(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDenominator), v))
	})
}

// DenominatorLTE applies the LTE predicate on the "denominator" field.
func DenominatorLTE(v int32) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDenominator), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Refinement) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Refinement) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
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
func Not(p predicate.Refinement) predicate.Refinement {
	return predicate.Refinement(func(s *sql.Selector) {
		p(s.Not())
	})
}
