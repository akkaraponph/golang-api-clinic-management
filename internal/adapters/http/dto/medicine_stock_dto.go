package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineStockCreateRequest struct {
	MedicineID  uuid.UUID  `json:"medicine_id" validate:"required"`
	Qty         int        `json:"qty" validate:"required,min=0"`
	ExpiredDate *time.Time `json:"expired_date,omitempty"`
}

type MedicineStockUpdateRequest struct {
	MedicineID  *uuid.UUID `json:"medicine_id,omitempty"`
	Qty         *int       `json:"qty,omitempty" validate:"omitempty,min=0"`
	ExpiredDate *time.Time `json:"expired_date,omitempty"`
}

type MedicineStockResponse struct {
	ID          uuid.UUID  `json:"id"`
	MedicineID  uuid.UUID  `json:"medicine_id"`
	Qty         int        `json:"qty"`
	ExpiredDate *time.Time `json:"expired_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func ToMedicineStockResponse(ms *domain.MedicineStock) *MedicineStockResponse {
	return &MedicineStockResponse{
		ID:          ms.ID,
		MedicineID:  ms.MedicineID,
		Qty:         ms.Qty,
		ExpiredDate: ms.ExpiredDate,
		CreatedAt:   ms.CreatedAt,
		UpdatedAt:   ms.UpdatedAt,
	}
}

func ToMedicineStockDomain(req *MedicineStockCreateRequest) *domain.MedicineStock {
	return &domain.MedicineStock{
		MedicineID:  req.MedicineID,
		Qty:         req.Qty,
		ExpiredDate: req.ExpiredDate,
	}
}
