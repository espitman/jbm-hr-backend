package otp

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	"github.com/espitman/jbm-hr-backend/ent/otp"
	"github.com/espitman/jbm-hr-backend/ent/user"
)

type repository struct {
	client *ent.Client
}

// New creates a new OTP repository
func NewEntRepository(client *ent.Client) Repository {
	return &repository{
		client: client,
	}
}

// Create creates a new OTP for a user
func (r *repository) Create(ctx context.Context, email string, code string, expiresAt time.Time) (*contract.OTP, error) {
	// First find the user by email
	user, err := r.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	otp, err := r.client.OTP.
		Create().
		SetCode(code).
		SetExpiresAt(expiresAt).
		SetUserID(user.ID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &contract.OTP{
		ID:        otp.ID,
		Code:      otp.Code,
		ExpiresAt: otp.ExpiresAt,
		Used:      otp.Used,
		CreatedAt: otp.CreatedAt,
		UsedAt:    otp.UsedAt,
		UserID:    user.ID,
	}, nil
}

// GetByCode retrieves an OTP by its code
func (r *repository) GetByCode(ctx context.Context, code string) (*contract.OTP, error) {
	otp, err := r.client.OTP.
		Query().
		Where(otp.Code(code)).
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	user, err := otp.QueryUser().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &contract.OTP{
		ID:        otp.ID,
		Code:      otp.Code,
		ExpiresAt: otp.ExpiresAt,
		Used:      otp.Used,
		CreatedAt: otp.CreatedAt,
		UsedAt:    otp.UsedAt,
		UserID:    user.ID,
	}, nil
}

// GetActiveByEmail retrieves active (unused and not expired) OTPs for a user
func (r *repository) GetActiveByEmail(ctx context.Context, email string) ([]*contract.OTP, error) {
	// First find the newUser by email
	newUser, err := r.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	otps, err := r.client.OTP.
		Query().
		Where(
			otp.HasUserWith(user.ID(newUser.ID)),
			otp.Used(false),
			otp.ExpiresAtGT(time.Now()),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*contract.OTP, len(otps))
	for i, otp := range otps {
		result[i] = &contract.OTP{
			ID:        otp.ID,
			Code:      otp.Code,
			ExpiresAt: otp.ExpiresAt,
			Used:      otp.Used,
			CreatedAt: otp.CreatedAt,
			UsedAt:    otp.UsedAt,
			UserID:    newUser.ID,
		}
	}
	return result, nil
}

// MarkAsUsed marks an OTP as used
func (r *repository) MarkAsUsed(ctx context.Context, id int) error {
	return r.client.OTP.
		UpdateOneID(id).
		SetUsed(true).
		SetUsedAt(time.Now()).
		Exec(ctx)
}

// DeleteExpired deletes all expired OTPs
func (r *repository) DeleteExpired(ctx context.Context) error {
	_, err := r.client.OTP.
		Delete().
		Where(otp.ExpiresAtLT(time.Now())).
		Exec(ctx)
	return err
}
