// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/espitman/jbm-hr-backend/ent/alibabacode"
	"github.com/espitman/jbm-hr-backend/ent/department"
	"github.com/espitman/jbm-hr-backend/ent/digikalacode"
	"github.com/espitman/jbm-hr-backend/ent/otp"
	"github.com/espitman/jbm-hr-backend/ent/request"
	"github.com/espitman/jbm-hr-backend/ent/resume"
	"github.com/espitman/jbm-hr-backend/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetFirstName sets the "first_name" field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetFullName sets the "full_name" field.
func (uc *UserCreate) SetFullName(s string) *UserCreate {
	uc.mutation.SetFullName(s)
	return uc
}

// SetRole sets the "role" field.
func (uc *UserCreate) SetRole(u user.Role) *UserCreate {
	uc.mutation.SetRole(u)
	return uc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(u *user.Role) *UserCreate {
	if u != nil {
		uc.SetRole(*u)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetPersonnelNumber sets the "personnel_number" field.
func (uc *UserCreate) SetPersonnelNumber(s string) *UserCreate {
	uc.mutation.SetPersonnelNumber(s)
	return uc
}

// SetNationalCode sets the "national_code" field.
func (uc *UserCreate) SetNationalCode(s string) *UserCreate {
	uc.mutation.SetNationalCode(s)
	return uc
}

// SetBirthdate sets the "birthdate" field.
func (uc *UserCreate) SetBirthdate(t time.Time) *UserCreate {
	uc.mutation.SetBirthdate(t)
	return uc
}

// SetNillableBirthdate sets the "birthdate" field if the given value is not nil.
func (uc *UserCreate) SetNillableBirthdate(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetBirthdate(*t)
	}
	return uc
}

// SetCooperationStartDate sets the "cooperation_start_date" field.
func (uc *UserCreate) SetCooperationStartDate(t time.Time) *UserCreate {
	uc.mutation.SetCooperationStartDate(t)
	return uc
}

// SetNillableCooperationStartDate sets the "cooperation_start_date" field if the given value is not nil.
func (uc *UserCreate) SetNillableCooperationStartDate(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCooperationStartDate(*t)
	}
	return uc
}

// SetConfirmed sets the "confirmed" field.
func (uc *UserCreate) SetConfirmed(b bool) *UserCreate {
	uc.mutation.SetConfirmed(b)
	return uc
}

// SetNillableConfirmed sets the "confirmed" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmed(b *bool) *UserCreate {
	if b != nil {
		uc.SetConfirmed(*b)
	}
	return uc
}

// SetActive sets the "active" field.
func (uc *UserCreate) SetActive(b bool) *UserCreate {
	uc.mutation.SetActive(b)
	return uc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uc *UserCreate) SetNillableActive(b *bool) *UserCreate {
	if b != nil {
		uc.SetActive(*b)
	}
	return uc
}

// AddOtpIDs adds the "otps" edge to the OTP entity by IDs.
func (uc *UserCreate) AddOtpIDs(ids ...int) *UserCreate {
	uc.mutation.AddOtpIDs(ids...)
	return uc
}

// AddOtps adds the "otps" edges to the OTP entity.
func (uc *UserCreate) AddOtps(o ...*OTP) *UserCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uc.AddOtpIDs(ids...)
}

// AddResumeIDs adds the "resumes" edge to the Resume entity by IDs.
func (uc *UserCreate) AddResumeIDs(ids ...int) *UserCreate {
	uc.mutation.AddResumeIDs(ids...)
	return uc
}

// AddResumes adds the "resumes" edges to the Resume entity.
func (uc *UserCreate) AddResumes(r ...*Resume) *UserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddResumeIDs(ids...)
}

// AddRequestIDs adds the "requests" edge to the Request entity by IDs.
func (uc *UserCreate) AddRequestIDs(ids ...int) *UserCreate {
	uc.mutation.AddRequestIDs(ids...)
	return uc
}

// AddRequests adds the "requests" edges to the Request entity.
func (uc *UserCreate) AddRequests(r ...*Request) *UserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRequestIDs(ids...)
}

// SetDepartmentID sets the "department" edge to the Department entity by ID.
func (uc *UserCreate) SetDepartmentID(id int) *UserCreate {
	uc.mutation.SetDepartmentID(id)
	return uc
}

// SetNillableDepartmentID sets the "department" edge to the Department entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableDepartmentID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetDepartmentID(*id)
	}
	return uc
}

// SetDepartment sets the "department" edge to the Department entity.
func (uc *UserCreate) SetDepartment(d *Department) *UserCreate {
	return uc.SetDepartmentID(d.ID)
}

// AddDigikalaCodeIDs adds the "digikala_codes" edge to the DigikalaCode entity by IDs.
func (uc *UserCreate) AddDigikalaCodeIDs(ids ...int) *UserCreate {
	uc.mutation.AddDigikalaCodeIDs(ids...)
	return uc
}

// AddDigikalaCodes adds the "digikala_codes" edges to the DigikalaCode entity.
func (uc *UserCreate) AddDigikalaCodes(d ...*DigikalaCode) *UserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDigikalaCodeIDs(ids...)
}

// AddAlibabaCodeIDs adds the "alibaba_codes" edge to the AlibabaCode entity by IDs.
func (uc *UserCreate) AddAlibabaCodeIDs(ids ...int) *UserCreate {
	uc.mutation.AddAlibabaCodeIDs(ids...)
	return uc
}

// AddAlibabaCodes adds the "alibaba_codes" edges to the AlibabaCode entity.
func (uc *UserCreate) AddAlibabaCodes(a ...*AlibabaCode) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAlibabaCodeIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Role(); !ok {
		v := user.DefaultRole
		uc.mutation.SetRole(v)
	}
	if _, ok := uc.mutation.Confirmed(); !ok {
		v := user.DefaultConfirmed
		uc.mutation.SetConfirmed(v)
	}
	if _, ok := uc.mutation.Active(); !ok {
		v := user.DefaultActive
		uc.mutation.SetActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "User.phone"`)}
	}
	if v, ok := uc.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if _, ok := uc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "User.first_name"`)}
	}
	if v, ok := uc.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "User.last_name"`)}
	}
	if v, ok := uc.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "User.full_name"`)}
	}
	if v, ok := uc.mutation.FullName(); ok {
		if err := user.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "User.full_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "User.role"`)}
	}
	if v, ok := uc.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if _, ok := uc.mutation.PersonnelNumber(); !ok {
		return &ValidationError{Name: "personnel_number", err: errors.New(`ent: missing required field "User.personnel_number"`)}
	}
	if v, ok := uc.mutation.PersonnelNumber(); ok {
		if err := user.PersonnelNumberValidator(v); err != nil {
			return &ValidationError{Name: "personnel_number", err: fmt.Errorf(`ent: validator failed for field "User.personnel_number": %w`, err)}
		}
	}
	if _, ok := uc.mutation.NationalCode(); !ok {
		return &ValidationError{Name: "national_code", err: errors.New(`ent: missing required field "User.national_code"`)}
	}
	if v, ok := uc.mutation.NationalCode(); ok {
		if err := user.NationalCodeValidator(v); err != nil {
			return &ValidationError{Name: "national_code", err: fmt.Errorf(`ent: validator failed for field "User.national_code": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Confirmed(); !ok {
		return &ValidationError{Name: "confirmed", err: errors.New(`ent: missing required field "User.confirmed"`)}
	}
	if _, ok := uc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "User.active"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := uc.mutation.FullName(); ok {
		_spec.SetField(user.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.PersonnelNumber(); ok {
		_spec.SetField(user.FieldPersonnelNumber, field.TypeString, value)
		_node.PersonnelNumber = value
	}
	if value, ok := uc.mutation.NationalCode(); ok {
		_spec.SetField(user.FieldNationalCode, field.TypeString, value)
		_node.NationalCode = value
	}
	if value, ok := uc.mutation.Birthdate(); ok {
		_spec.SetField(user.FieldBirthdate, field.TypeTime, value)
		_node.Birthdate = value
	}
	if value, ok := uc.mutation.CooperationStartDate(); ok {
		_spec.SetField(user.FieldCooperationStartDate, field.TypeTime, value)
		_node.CooperationStartDate = value
	}
	if value, ok := uc.mutation.Confirmed(); ok {
		_spec.SetField(user.FieldConfirmed, field.TypeBool, value)
		_node.Confirmed = value
	}
	if value, ok := uc.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if nodes := uc.mutation.OtpsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OtpsTable,
			Columns: []string{user.OtpsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(otp.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ResumesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ResumesTable,
			Columns: []string{user.ResumesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resume.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RequestsTable,
			Columns: []string{user.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.department_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DigikalaCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DigikalaCodesTable,
			Columns: []string{user.DigikalaCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(digikalacode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AlibabaCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AlibabaCodesTable,
			Columns: []string{user.AlibabaCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(alibabacode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
