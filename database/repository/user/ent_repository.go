package user

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entUser "github.com/espitman/jbm-hr-backend/ent/user"
)

// EntRepository implements the Repository interface using Ent
type EntRepository struct {
	client *ent.Client
}

// NewEntRepository creates a new EntRepository
func NewEntRepository(client *ent.Client) *EntRepository {
	return &EntRepository{
		client: client,
	}
}

// convertToContractUser converts an ent.User to a contract.User
func convertToContractUser(entUser *ent.User) *contract.User {
	if entUser == nil {
		return nil
	}
	var departmentID *int
	if entUser.Edges.Department != nil {
		departmentID = &entUser.Edges.Department.ID
	}
	return &contract.User{
		ID:           entUser.ID,
		Email:        entUser.Email,
		Phone:        entUser.Phone,
		FirstName:    entUser.FirstName,
		LastName:     entUser.LastName,
		Role:         string(entUser.Role),
		Avatar:       entUser.Avatar,
		Password:     entUser.Password,
		DepartmentID: departmentID,
	}
}

// GetAll retrieves all users
func (r *EntRepository) GetAll(ctx context.Context) ([]*contract.User, error) {
	entUsers, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*contract.User, len(entUsers))
	for i, entUser := range entUsers {
		users[i] = convertToContractUser(entUser)
	}
	return users, nil
}

// GetByID retrieves a user by their ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.User, error) {
	entUser, err := r.client.User.Query().Where(entUser.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// GetByEmail retrieves a user by their email
func (r *EntRepository) GetByEmail(ctx context.Context, email string) (*contract.User, error) {
	entUser, err := r.client.User.Query().Where(entUser.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// Create creates a new user
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateUserInput) (*contract.User, error) {
	create := r.client.User.
		Create().
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetRole(entUser.Role(req.Role)).
		SetAvatar(req.Avatar)

	if req.DepartmentID != nil {
		create = create.SetDepartmentID(*req.DepartmentID)
	}

	entUser, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// Update updates an existing user
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.UpdateUserInput) (*contract.User, error) {
	update := r.client.User.
		UpdateOneID(id).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetRole(entUser.Role(req.Role)).
		SetAvatar(req.Avatar)

	if req.DepartmentID != nil {
		update = update.SetDepartmentID(*req.DepartmentID)
	} else {
		update = update.ClearDepartment()
	}

	entUser, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// Delete deletes a user by their ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}

// UpdatePassword updates a user's password
func (r *EntRepository) UpdatePassword(ctx context.Context, id int, req *contract.UpdatePasswordInput) error {
	return r.client.User.
		UpdateOneID(id).
		SetPassword(req.Password).
		Exec(ctx)
}

// UpdateAvatar updates only the avatar of a user
func (r *EntRepository) UpdateAvatar(ctx context.Context, id int, avatar string) (*contract.User, error) {
	entUser, err := r.client.User.
		UpdateOneID(id).
		SetAvatar(avatar).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}
