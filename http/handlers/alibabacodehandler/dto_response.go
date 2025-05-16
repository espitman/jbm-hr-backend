package alibabacodehandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// CreateAlibabaCodeResponse represents the response for creating an Alibaba code
type CreateAlibabaCodeResponse struct {
	dto.Response
	Data contract.AlibabaCode `json:"data"`
}

// GetAlibabaCodeResponse represents the response for getting an Alibaba code
type GetAlibabaCodeResponse struct {
	dto.Response
	Data contract.AlibabaCode `json:"data"`
}

// AlibabaCodeListData represents the data structure for listing Alibaba codes
type AlibabaCodeListData struct {
	Items []contract.AlibabaCode `json:"items"`
	Total int                    `json:"total"`
}

// ListAlibabaCodeResponse represents the response for listing Alibaba codes
type ListAlibabaCodeResponse struct {
	dto.Response
	Data AlibabaCodeListData `json:"data"`
}

// AssignAlibabaCodeResponse represents the response for assigning an Alibaba code
type AssignAlibabaCodeResponse struct {
	dto.Response
	Data contract.AlibabaCode `json:"data"`
}
