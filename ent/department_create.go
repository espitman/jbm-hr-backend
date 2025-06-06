// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/espitman/jbm-hr-backend/ent/department"
	"github.com/espitman/jbm-hr-backend/ent/user"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (dc *DepartmentCreate) SetTitle(s string) *DepartmentCreate {
	dc.mutation.SetTitle(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DepartmentCreate) SetDescription(s string) *DepartmentCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetImage sets the "image" field.
func (dc *DepartmentCreate) SetImage(s string) *DepartmentCreate {
	dc.mutation.SetImage(s)
	return dc
}

// SetIcon sets the "icon" field.
func (dc *DepartmentCreate) SetIcon(s string) *DepartmentCreate {
	dc.mutation.SetIcon(s)
	return dc
}

// SetColor sets the "color" field.
func (dc *DepartmentCreate) SetColor(s string) *DepartmentCreate {
	dc.mutation.SetColor(s)
	return dc
}

// SetShortName sets the "shortName" field.
func (dc *DepartmentCreate) SetShortName(s string) *DepartmentCreate {
	dc.mutation.SetShortName(s)
	return dc
}

// SetDisplayOrder sets the "display_order" field.
func (dc *DepartmentCreate) SetDisplayOrder(i int) *DepartmentCreate {
	dc.mutation.SetDisplayOrder(i)
	return dc
}

// SetNillableDisplayOrder sets the "display_order" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableDisplayOrder(i *int) *DepartmentCreate {
	if i != nil {
		dc.SetDisplayOrder(*i)
	}
	return dc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (dc *DepartmentCreate) AddUserIDs(ids ...int) *DepartmentCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the "users" edges to the User entity.
func (dc *DepartmentCreate) AddUsers(u ...*User) *DepartmentCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() {
	if _, ok := dc.mutation.DisplayOrder(); !ok {
		v := department.DefaultDisplayOrder
		dc.mutation.SetDisplayOrder(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Department.title"`)}
	}
	if v, ok := dc.mutation.Title(); ok {
		if err := department.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Department.title": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Department.description"`)}
	}
	if v, ok := dc.mutation.Description(); ok {
		if err := department.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Department.description": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "Department.image"`)}
	}
	if v, ok := dc.mutation.Image(); ok {
		if err := department.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Department.image": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "Department.icon"`)}
	}
	if v, ok := dc.mutation.Icon(); ok {
		if err := department.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Department.icon": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "Department.color"`)}
	}
	if v, ok := dc.mutation.Color(); ok {
		if err := department.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "Department.color": %w`, err)}
		}
	}
	if _, ok := dc.mutation.ShortName(); !ok {
		return &ValidationError{Name: "shortName", err: errors.New(`ent: missing required field "Department.shortName"`)}
	}
	if v, ok := dc.mutation.ShortName(); ok {
		if err := department.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "shortName", err: fmt.Errorf(`ent: validator failed for field "Department.shortName": %w`, err)}
		}
	}
	if _, ok := dc.mutation.DisplayOrder(); !ok {
		return &ValidationError{Name: "display_order", err: errors.New(`ent: missing required field "Department.display_order"`)}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt))
	)
	if value, ok := dc.mutation.Title(); ok {
		_spec.SetField(department.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(department.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dc.mutation.Image(); ok {
		_spec.SetField(department.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := dc.mutation.Icon(); ok {
		_spec.SetField(department.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := dc.mutation.Color(); ok {
		_spec.SetField(department.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if value, ok := dc.mutation.ShortName(); ok {
		_spec.SetField(department.FieldShortName, field.TypeString, value)
		_node.ShortName = value
	}
	if value, ok := dc.mutation.DisplayOrder(); ok {
		_spec.SetField(department.FieldDisplayOrder, field.TypeInt, value)
		_node.DisplayOrder = value
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	err      error
	builders []*DepartmentCreate
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
