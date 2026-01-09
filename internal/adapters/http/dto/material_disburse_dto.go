package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MaterialDisburseDetailCreateRequest struct {
	MaterialID uuid.UUID `json:"material_id" validate:"required"`
	EmployeeID uuid.UUID `json:"employee_id" validate:"required"`
	Qty        int       `json:"qty" validate:"required,min=1"`
}

type MaterialDisburseCreateRequest struct {
	Date       time.Time                            `json:"date" validate:"required"`
	EmployeeID uuid.UUID                            `json:"employee_id" validate:"required"`
	Details    []MaterialDisburseDetailCreateRequest `json:"details"`
}

type MaterialDisburseDetailResponse struct {
	MaterialDisburseID uuid.UUID `json:"material_disburse_id"`
	MaterialID         uuid.UUID `json:"material_id"`
	EmployeeID         uuid.UUID `json:"employee_id"`
	Qty                int       `json:"qty"`
}

type MaterialDisburseResponse struct {
	ID       uuid.UUID                        `json:"id"`
	Date     time.Time                        `json:"date"`
	EmployeeID uuid.UUID                      `json:"employee_id"`
	Details  []MaterialDisburseDetailResponse `json:"details,omitempty"`
	CreatedAt time.Time                       `json:"created_at"`
	UpdatedAt time.Time                       `json:"updated_at"`
}

func ToMaterialDisburseDetailResponse(mdd *domain.MaterialDisburseDetail) *MaterialDisburseDetailResponse {
	return &MaterialDisburseDetailResponse{
		MaterialDisburseID: mdd.MaterialDisburseID,
		MaterialID:         mdd.MaterialID,
		EmployeeID:         mdd.EmployeeID,
		Qty:                mdd.Qty,
	}
}

func ToMaterialDisburseResponse(md *domain.MaterialDisburse) *MaterialDisburseResponse {
	details := make([]MaterialDisburseDetailResponse, len(md.MaterialDisburseDetails))
	for i, d := range md.MaterialDisburseDetails {
		details[i] = *ToMaterialDisburseDetailResponse(&d)
	}

	return &MaterialDisburseResponse{
		ID:         md.ID,
		Date:       md.Date,
		EmployeeID: md.EmployeeID,
		Details:    details,
		CreatedAt:  md.CreatedAt,
		UpdatedAt:  md.UpdatedAt,
	}
}

func ToMaterialDisburseDomain(req *MaterialDisburseCreateRequest, disburseID uuid.UUID) *domain.MaterialDisburse {
	details := make([]domain.MaterialDisburseDetail, len(req.Details))
	for i, d := range req.Details {
		details[i] = domain.MaterialDisburseDetail{
			MaterialDisburseID: disburseID,
			MaterialID:         d.MaterialID,
			EmployeeID:         d.EmployeeID,
			Qty:                d.Qty,
		}
	}

	return &domain.MaterialDisburse{
		Date:                  req.Date,
		EmployeeID:            req.EmployeeID,
		MaterialDisburseDetails: details,
	}
}
