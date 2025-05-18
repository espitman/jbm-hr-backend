package infohandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// DashboardInfoResponse represents the response for dashboard info
type DashboardInfoResponse struct {
	dto.Response
	contract.DashboardInfo
}
