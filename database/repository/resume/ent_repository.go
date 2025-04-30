package resume

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	"github.com/espitman/jbm-hr-backend/ent/resume"
)

type repository struct {
	client *ent.Client
}

// NewRepository creates a new resume repository
func NewRepository(client *ent.Client) Repository {
	return &repository{
		client: client,
	}
}

// GetAll retrieves all resumes
func (r *repository) GetAll(ctx context.Context) ([]*contract.Resume, error) {
	resumes, err := r.client.Resume.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*contract.Resume
	for _, res := range resumes {
		result = append(result, &contract.Resume{
			ID:              res.ID,
			IntroducedName:  res.IntroducedName,
			IntroducedPhone: res.IntroducedPhone,
			Position:        res.Position,
			File:            res.File,
			Status:          string(res.Status),
			UserID:          res.UserID,
			CreatedAt:       res.CreatedAt.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
		})
	}

	return result, nil
}

// GetByID retrieves a resume by their ID
func (r *repository) GetByID(ctx context.Context, id int) (*contract.Resume, error) {
	res, err := r.client.Resume.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &contract.Resume{
		ID:              res.ID,
		IntroducedName:  res.IntroducedName,
		IntroducedPhone: res.IntroducedPhone,
		Position:        res.Position,
		File:            res.File,
		Status:          string(res.Status),
		UserID:          res.UserID,
		CreatedAt:       res.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// Create creates a new resume
func (r *repository) Create(ctx context.Context, req *contract.ResumeInput) (*contract.Resume, error) {
	res, err := r.client.Resume.Create().
		SetIntroducedName(req.IntroducedName).
		SetIntroducedPhone(req.IntroducedPhone).
		SetPosition(req.Position).
		SetFile(req.File).
		SetUserID(req.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &contract.Resume{
		ID:              res.ID,
		IntroducedName:  res.IntroducedName,
		IntroducedPhone: res.IntroducedPhone,
		Position:        res.Position,
		File:            res.File,
		Status:          string(res.Status),
		UserID:          res.UserID,
		CreatedAt:       res.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// Update updates an existing resume
func (r *repository) Update(ctx context.Context, id int, req *contract.ResumeInput) (*contract.Resume, error) {
	res, err := r.client.Resume.UpdateOneID(id).
		SetIntroducedName(req.IntroducedName).
		SetIntroducedPhone(req.IntroducedPhone).
		SetPosition(req.Position).
		SetFile(req.File).
		SetUserID(req.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &contract.Resume{
		ID:              res.ID,
		IntroducedName:  res.IntroducedName,
		IntroducedPhone: res.IntroducedPhone,
		Position:        res.Position,
		File:            res.File,
		Status:          string(res.Status),
		UserID:          res.UserID,
		CreatedAt:       res.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// Delete deletes a resume by their ID
func (r *repository) Delete(ctx context.Context, id int) error {
	return r.client.Resume.DeleteOneID(id).Exec(ctx)
}

// List retrieves a paginated list of resumes
func (r *repository) List(ctx context.Context, page, limit int) ([]*contract.Resume, int, error) {
	offset := (page - 1) * limit

	// Get total count
	total, err := r.client.Resume.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	resumes, err := r.client.Resume.Query().
		Offset(offset).
		Limit(limit).
		Order(ent.Desc(resume.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	var result []*contract.Resume
	for _, res := range resumes {
		result = append(result, &contract.Resume{
			ID:              res.ID,
			IntroducedName:  res.IntroducedName,
			IntroducedPhone: res.IntroducedPhone,
			Position:        res.Position,
			File:            res.File,
			Status:          string(res.Status),
			UserID:          res.UserID,
			CreatedAt:       res.CreatedAt.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
		})
	}

	return result, total, nil
}
