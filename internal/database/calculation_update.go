// Code generated by entc, DO NOT EDIT.

package database

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/iamnande/cardmod/internal/database/calculation"
	"github.com/iamnande/cardmod/internal/database/predicate"
)

// CalculationUpdate is the builder for updating Calculation entities.
type CalculationUpdate struct {
	config
	hooks    []Hook
	mutation *CalculationMutation
}

// Where appends a list predicates to the CalculationUpdate builder.
func (cu *CalculationUpdate) Where(ps ...predicate.Calculation) *CalculationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCardID sets the "card_id" field.
func (cu *CalculationUpdate) SetCardID(u uuid.UUID) *CalculationUpdate {
	cu.mutation.SetCardID(u)
	return cu
}

// SetMagicID sets the "magic_id" field.
func (cu *CalculationUpdate) SetMagicID(u uuid.UUID) *CalculationUpdate {
	cu.mutation.SetMagicID(u)
	return cu
}

// SetCardRatio sets the "card_ratio" field.
func (cu *CalculationUpdate) SetCardRatio(i int32) *CalculationUpdate {
	cu.mutation.ResetCardRatio()
	cu.mutation.SetCardRatio(i)
	return cu
}

// AddCardRatio adds i to the "card_ratio" field.
func (cu *CalculationUpdate) AddCardRatio(i int32) *CalculationUpdate {
	cu.mutation.AddCardRatio(i)
	return cu
}

// SetMagicRatio sets the "magic_ratio" field.
func (cu *CalculationUpdate) SetMagicRatio(i int32) *CalculationUpdate {
	cu.mutation.ResetMagicRatio()
	cu.mutation.SetMagicRatio(i)
	return cu
}

// AddMagicRatio adds i to the "magic_ratio" field.
func (cu *CalculationUpdate) AddMagicRatio(i int32) *CalculationUpdate {
	cu.mutation.AddMagicRatio(i)
	return cu
}

// Mutation returns the CalculationMutation object of the builder.
func (cu *CalculationUpdate) Mutation() *CalculationMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CalculationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CalculationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("database: uninitialized hook (forgotten import database/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CalculationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CalculationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CalculationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CalculationUpdate) check() error {
	if v, ok := cu.mutation.CardRatio(); ok {
		if err := calculation.CardRatioValidator(v); err != nil {
			return &ValidationError{Name: "card_ratio", err: fmt.Errorf("database: validator failed for field \"card_ratio\": %w", err)}
		}
	}
	if v, ok := cu.mutation.MagicRatio(); ok {
		if err := calculation.MagicRatioValidator(v); err != nil {
			return &ValidationError{Name: "magic_ratio", err: fmt.Errorf("database: validator failed for field \"magic_ratio\": %w", err)}
		}
	}
	return nil
}

func (cu *CalculationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   calculation.Table,
			Columns: calculation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: calculation.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CardID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: calculation.FieldCardID,
		})
	}
	if value, ok := cu.mutation.MagicID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: calculation.FieldMagicID,
		})
	}
	if value, ok := cu.mutation.CardRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldCardRatio,
		})
	}
	if value, ok := cu.mutation.AddedCardRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldCardRatio,
		})
	}
	if value, ok := cu.mutation.MagicRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldMagicRatio,
		})
	}
	if value, ok := cu.mutation.AddedMagicRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldMagicRatio,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{calculation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CalculationUpdateOne is the builder for updating a single Calculation entity.
type CalculationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CalculationMutation
}

// SetCardID sets the "card_id" field.
func (cuo *CalculationUpdateOne) SetCardID(u uuid.UUID) *CalculationUpdateOne {
	cuo.mutation.SetCardID(u)
	return cuo
}

// SetMagicID sets the "magic_id" field.
func (cuo *CalculationUpdateOne) SetMagicID(u uuid.UUID) *CalculationUpdateOne {
	cuo.mutation.SetMagicID(u)
	return cuo
}

// SetCardRatio sets the "card_ratio" field.
func (cuo *CalculationUpdateOne) SetCardRatio(i int32) *CalculationUpdateOne {
	cuo.mutation.ResetCardRatio()
	cuo.mutation.SetCardRatio(i)
	return cuo
}

// AddCardRatio adds i to the "card_ratio" field.
func (cuo *CalculationUpdateOne) AddCardRatio(i int32) *CalculationUpdateOne {
	cuo.mutation.AddCardRatio(i)
	return cuo
}

// SetMagicRatio sets the "magic_ratio" field.
func (cuo *CalculationUpdateOne) SetMagicRatio(i int32) *CalculationUpdateOne {
	cuo.mutation.ResetMagicRatio()
	cuo.mutation.SetMagicRatio(i)
	return cuo
}

// AddMagicRatio adds i to the "magic_ratio" field.
func (cuo *CalculationUpdateOne) AddMagicRatio(i int32) *CalculationUpdateOne {
	cuo.mutation.AddMagicRatio(i)
	return cuo
}

// Mutation returns the CalculationMutation object of the builder.
func (cuo *CalculationUpdateOne) Mutation() *CalculationMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CalculationUpdateOne) Select(field string, fields ...string) *CalculationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Calculation entity.
func (cuo *CalculationUpdateOne) Save(ctx context.Context) (*Calculation, error) {
	var (
		err  error
		node *Calculation
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CalculationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("database: uninitialized hook (forgotten import database/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CalculationUpdateOne) SaveX(ctx context.Context) *Calculation {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CalculationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CalculationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CalculationUpdateOne) check() error {
	if v, ok := cuo.mutation.CardRatio(); ok {
		if err := calculation.CardRatioValidator(v); err != nil {
			return &ValidationError{Name: "card_ratio", err: fmt.Errorf("database: validator failed for field \"card_ratio\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.MagicRatio(); ok {
		if err := calculation.MagicRatioValidator(v); err != nil {
			return &ValidationError{Name: "magic_ratio", err: fmt.Errorf("database: validator failed for field \"magic_ratio\": %w", err)}
		}
	}
	return nil
}

func (cuo *CalculationUpdateOne) sqlSave(ctx context.Context) (_node *Calculation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   calculation.Table,
			Columns: calculation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: calculation.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Calculation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, calculation.FieldID)
		for _, f := range fields {
			if !calculation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("database: invalid field %q for query", f)}
			}
			if f != calculation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CardID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: calculation.FieldCardID,
		})
	}
	if value, ok := cuo.mutation.MagicID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: calculation.FieldMagicID,
		})
	}
	if value, ok := cuo.mutation.CardRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldCardRatio,
		})
	}
	if value, ok := cuo.mutation.AddedCardRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldCardRatio,
		})
	}
	if value, ok := cuo.mutation.MagicRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldMagicRatio,
		})
	}
	if value, ok := cuo.mutation.AddedMagicRatio(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: calculation.FieldMagicRatio,
		})
	}
	_node = &Calculation{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{calculation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
