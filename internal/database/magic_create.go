// Code generated by entc, DO NOT EDIT.

package database

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/iamnande/cardmod/internal/database/magic"
)

// MagicCreate is the builder for creating a Magic entity.
type MagicCreate struct {
	config
	mutation *MagicMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *MagicCreate) SetName(s string) *MagicCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MagicCreate) SetID(u uuid.UUID) *MagicCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MagicCreate) SetNillableID(u *uuid.UUID) *MagicCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// Mutation returns the MagicMutation object of the builder.
func (mc *MagicCreate) Mutation() *MagicMutation {
	return mc.mutation
}

// Save creates the Magic in the database.
func (mc *MagicCreate) Save(ctx context.Context) (*Magic, error) {
	var (
		err  error
		node *Magic
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MagicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("database: uninitialized hook (forgotten import database/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MagicCreate) SaveX(ctx context.Context) *Magic {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MagicCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MagicCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MagicCreate) defaults() {
	if _, ok := mc.mutation.ID(); !ok {
		v := magic.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MagicCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`database: missing required field "name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := magic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`database: validator failed for field "name": %w`, err)}
		}
	}
	return nil
}

func (mc *MagicCreate) sqlSave(ctx context.Context) (*Magic, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (mc *MagicCreate) createSpec() (*Magic, *sqlgraph.CreateSpec) {
	var (
		_node = &Magic{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: magic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: magic.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: magic.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// MagicCreateBulk is the builder for creating many Magic entities in bulk.
type MagicCreateBulk struct {
	config
	builders []*MagicCreate
}

// Save creates the Magic entities in the database.
func (mcb *MagicCreateBulk) Save(ctx context.Context) ([]*Magic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Magic, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MagicMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MagicCreateBulk) SaveX(ctx context.Context) []*Magic {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MagicCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MagicCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
