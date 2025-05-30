// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/espitman/jbm-hr-backend/ent/hrteam"
	"github.com/espitman/jbm-hr-backend/ent/predicate"
)

// HRTeamDelete is the builder for deleting a HRTeam entity.
type HRTeamDelete struct {
	config
	hooks    []Hook
	mutation *HRTeamMutation
}

// Where appends a list predicates to the HRTeamDelete builder.
func (htd *HRTeamDelete) Where(ps ...predicate.HRTeam) *HRTeamDelete {
	htd.mutation.Where(ps...)
	return htd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (htd *HRTeamDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, htd.sqlExec, htd.mutation, htd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (htd *HRTeamDelete) ExecX(ctx context.Context) int {
	n, err := htd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (htd *HRTeamDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hrteam.Table, sqlgraph.NewFieldSpec(hrteam.FieldID, field.TypeInt))
	if ps := htd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, htd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	htd.mutation.done = true
	return affected, err
}

// HRTeamDeleteOne is the builder for deleting a single HRTeam entity.
type HRTeamDeleteOne struct {
	htd *HRTeamDelete
}

// Where appends a list predicates to the HRTeamDelete builder.
func (htdo *HRTeamDeleteOne) Where(ps ...predicate.HRTeam) *HRTeamDeleteOne {
	htdo.htd.mutation.Where(ps...)
	return htdo
}

// Exec executes the deletion query.
func (htdo *HRTeamDeleteOne) Exec(ctx context.Context) error {
	n, err := htdo.htd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hrteam.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (htdo *HRTeamDeleteOne) ExecX(ctx context.Context) {
	if err := htdo.Exec(ctx); err != nil {
		panic(err)
	}
}
