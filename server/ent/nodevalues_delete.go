// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nullsploit01/iosync/ent/nodevalues"
	"github.com/nullsploit01/iosync/ent/predicate"
)

// NodeValuesDelete is the builder for deleting a NodeValues entity.
type NodeValuesDelete struct {
	config
	hooks    []Hook
	mutation *NodeValuesMutation
}

// Where appends a list predicates to the NodeValuesDelete builder.
func (nvd *NodeValuesDelete) Where(ps ...predicate.NodeValues) *NodeValuesDelete {
	nvd.mutation.Where(ps...)
	return nvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nvd *NodeValuesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nvd.sqlExec, nvd.mutation, nvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nvd *NodeValuesDelete) ExecX(ctx context.Context) int {
	n, err := nvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nvd *NodeValuesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(nodevalues.Table, sqlgraph.NewFieldSpec(nodevalues.FieldID, field.TypeInt))
	if ps := nvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nvd.mutation.done = true
	return affected, err
}

// NodeValuesDeleteOne is the builder for deleting a single NodeValues entity.
type NodeValuesDeleteOne struct {
	nvd *NodeValuesDelete
}

// Where appends a list predicates to the NodeValuesDelete builder.
func (nvdo *NodeValuesDeleteOne) Where(ps ...predicate.NodeValues) *NodeValuesDeleteOne {
	nvdo.nvd.mutation.Where(ps...)
	return nvdo
}

// Exec executes the deletion query.
func (nvdo *NodeValuesDeleteOne) Exec(ctx context.Context) error {
	n, err := nvdo.nvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nodevalues.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nvdo *NodeValuesDeleteOne) ExecX(ctx context.Context) {
	if err := nvdo.Exec(ctx); err != nil {
		panic(err)
	}
}
