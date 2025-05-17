package user

import (
	"context"
	"fmt"
	"time"

	"github.com/espitman/jbm-hr-backend/utils"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entDepartment "github.com/espitman/jbm-hr-backend/ent/department"
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
	var departmentTitle *string
	var departmentIcon *string
	var departmentShortName *string
	if entUser.Edges.Department != nil {
		departmentID = &entUser.Edges.Department.ID
		departmentTitle = &entUser.Edges.Department.Title
		departmentIcon = &entUser.Edges.Department.Icon
		departmentShortName = &entUser.Edges.Department.ShortName
	}
	var birthdate *string
	var age *int
	if !entUser.Birthdate.IsZero() {
		birthdateStr := entUser.Birthdate.Format("2006-01-02")
		birthdate = &birthdateStr

		// Calculate age based on year difference only
		today := time.Now()
		ageValue := today.Year() - entUser.Birthdate.Year()
		age = &ageValue
	}
	var cooperationStartDate *string
	var cooperationDuration *int
	if !entUser.CooperationStartDate.IsZero() {
		startDateStr := entUser.CooperationStartDate.Format("2006-01-02")
		cooperationStartDate = &startDateStr

		// Calculate cooperation duration based on year difference only
		today := time.Now()
		durationValue := today.Year() - entUser.CooperationStartDate.Year()
		cooperationDuration = &durationValue
	}

	return &contract.User{
		ID:                   entUser.ID,
		Email:                entUser.Email,
		Phone:                entUser.Phone,
		FirstName:            entUser.FirstName,
		LastName:             entUser.LastName,
		FullName:             entUser.FullName,
		Role:                 string(entUser.Role),
		Avatar:               entUser.Avatar,
		Password:             entUser.Password,
		DepartmentID:         departmentID,
		DepartmentTitle:      departmentTitle,
		DepartmentIcon:       departmentIcon,
		DepartmentShortName:  departmentShortName,
		Birthdate:            birthdate,
		CooperationStartDate: cooperationStartDate,
		PersonnelNumber:      entUser.PersonnelNumber,
		NationalCode:         entUser.NationalCode,
		Confirmed:            entUser.Confirmed,
		Age:                  age,
		CooperationDuration:  cooperationDuration,
	}
}

// GetAll retrieves all users
func (r *EntRepository) GetAll(ctx context.Context, filters *contract.UserFilters) ([]*contract.User, error) {
	query := r.client.User.Query().
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		})

	// Apply filters
	if filters != nil {
		if filters.FullName != nil {
			query = query.Where(entUser.FullNameContains(*filters.FullName))
		}
		if filters.Role != nil {
			query = query.Where(entUser.RoleEQ(entUser.Role(*filters.Role)))
		}
		if filters.PersonnelNumber != nil {
			query = query.Where(entUser.PersonnelNumberContains(*filters.PersonnelNumber))
		}
		if filters.NationalCode != nil {
			query = query.Where(entUser.NationalCodeContains(*filters.NationalCode))
		}
		if filters.Phone != nil {
			query = query.Where(entUser.PhoneContains(*filters.Phone))
		}
		if filters.DepartmentID != nil {
			query = query.Where(entUser.HasDepartmentWith(entDepartment.IDEQ(*filters.DepartmentID)))
		}
	}

	entUsers, err := query.Order(ent.Asc(entUser.FieldID)).All(ctx)
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
	entUser, err := r.client.User.Query().
		Where(entUser.ID(id)).
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// GetByEmail retrieves a user by their email
func (r *EntRepository) GetByEmail(ctx context.Context, email string) (*contract.User, error) {
	entUser, err := r.client.User.Query().
		Where(entUser.Email(email)).
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		}).
		Only(ctx)
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
		SetFullName(fmt.Sprintf("%s %s", req.FirstName, req.LastName)).
		SetRole(entUser.Role(req.Role)).
		SetAvatar(req.Avatar)

	if req.PersonnelNumber != "" {
		create = create.SetPersonnelNumber(req.PersonnelNumber)
	}
	if req.NationalCode != "" {
		create = create.SetNationalCode(req.NationalCode)
	}
	if req.DepartmentID != nil {
		create = create.SetDepartmentID(*req.DepartmentID)
	}
	if req.Birthdate != nil {
		birthdate, err := time.Parse("2006-01-02", *req.Birthdate)
		if err != nil {
			return nil, err
		}
		create = create.SetBirthdate(birthdate)
	}
	if req.CooperationStartDate != nil {
		startDate, err := time.Parse("2006-01-02", *req.CooperationStartDate)
		if err != nil {
			return nil, err
		}
		create = create.SetCooperationStartDate(startDate)
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
		SetFullName(fmt.Sprintf("%s %s", req.FirstName, req.LastName)).
		SetRole(entUser.Role(req.Role))

	if req.PersonnelNumber != "" {
		update = update.SetPersonnelNumber(req.PersonnelNumber)
	}
	if req.NationalCode != "" {
		update = update.SetNationalCode(req.NationalCode)
	}
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

// UpdateBirthdate updates a user's birthdate
func (r *EntRepository) UpdateBirthdate(ctx context.Context, id int, birthdate string) (*contract.User, error) {
	// Parse the birthdate string to time.Time
	birthdateTime, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return nil, err
	}

	entUser, err := r.client.User.
		UpdateOneID(id).
		SetBirthdate(birthdateTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// UpdateCooperationStartDate updates a user's cooperation start date
func (r *EntRepository) UpdateCooperationStartDate(ctx context.Context, id int, startDate string) (*contract.User, error) {
	// Parse the start date string to time.Time
	startDateTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}

	entUser, err := r.client.User.
		UpdateOneID(id).
		SetCooperationStartDate(startDateTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractUser(entUser), nil
}

// SearchUsers searches users by term (full name, email, or phone)
func (r *EntRepository) SearchUsers(ctx context.Context, term string) ([]*contract.User, error) {
	users, err := r.client.User.Query().
		Where(
			entUser.Or(
				entUser.FullNameHasPrefix(term),
				entUser.EmailHasPrefix(term),
				entUser.PhoneHasPrefix(term),
			),
		).
		WithDepartment().
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*contract.User, len(users))
	for i, u := range users {
		result[i] = convertToContractUser(u)
	}
	return result, nil
}

// UpdateConfirmed updates a user's confirmed status
func (r *EntRepository) UpdateConfirmed(ctx context.Context, id int) (*contract.User, error) {
	entUser, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	entUser, err = entUser.Update().
		SetConfirmed(true).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return convertToContractUser(entUser), nil
}

// GetUsersWithTodayBirthdate retrieves all users whose birthdate is today (ignoring year)
func (r *EntRepository) GetUsersWithTodayBirthdate(ctx context.Context) ([]*contract.User, error) {
	birthDates := utils.GenerateDatesForTodayMonthDay(100)

	users, err := r.client.User.Query().
		Where(
			entUser.BirthdateIn(birthDates...),
		).
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*contract.User, len(users))
	for i, u := range users {
		result[i] = convertToContractUser(u)
	}
	return result, nil
}

// GetUsersWithTodayCooperationStartDate retrieves all users whose cooperation start date is today (ignoring year)
func (r *EntRepository) GetUsersWithTodayCooperationStartDate(ctx context.Context) ([]*contract.User, error) {
	cooperationDates := utils.GenerateDatesForTodayMonthDay(20)

	users, err := r.client.User.Query().
		Where(
			entUser.CooperationStartDateIn(cooperationDates...),
		).
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*contract.User, len(users))
	for i, u := range users {
		result[i] = convertToContractUser(u)
	}
	return result, nil
}

// GetUsersWithBirthdateInJalaliMonth retrieves all users whose birthdate is in the current Jalali month
func (r *EntRepository) GetUsersWithBirthdateInJalaliMonth(ctx context.Context) ([]*contract.User, error) {
	dates := utils.GenerateDatesForJalaliMonth(100)

	users, err := r.client.User.Query().
		Where(
			entUser.BirthdateIn(dates...),
		).
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		}).
		Order(ent.Asc(entUser.FieldBirthdate)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*contract.User, len(users))
	for i, u := range users {
		result[i] = convertToContractUser(u)
	}
	return result, nil
}

// GetUsersWithCooperationStartDateInJalaliMonth retrieves all users whose cooperation start date is in the current Jalali month
func (r *EntRepository) GetUsersWithCooperationStartDateInJalaliMonth(ctx context.Context) ([]*contract.User, error) {
	dates := utils.GenerateDatesForJalaliMonth(100)

	users, err := r.client.User.Query().
		Where(
			entUser.CooperationStartDateIn(dates...),
		).
		WithDepartment(func(q *ent.DepartmentQuery) {
			q.Select("id", "title", "icon", "short_name")
		}).
		Order(ent.Asc(entUser.FieldCooperationStartDate)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*contract.User, len(users))
	for i, u := range users {
		result[i] = convertToContractUser(u)
	}
	return result, nil
}
