package infoservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

type Service interface {
	GetDashboardInfo(ctx context.Context) (*contract.DashboardInfo, error)
}
