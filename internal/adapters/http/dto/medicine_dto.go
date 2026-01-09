package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineCreateRequest struct {
	Name           string     `json:"name" validate:"required,max=255"`
	Detail         string     `json:"detail"`
	Dosage         string     `json:"dosage" validate:"max=100"`
	Unit           string     `json:"unit" validate:"max=50"`
	MedicineTypeID *uuid.UUID `json:"medicine_type_id,omitempty"`
}

type MedicineUpdateRequest struct {
	Name           string     `json:"name" validate:"max=255"`
	Detail         string     `json:"detail"`
	Dosage         string     `json:"dosage" validate:"max=100"`
	Unit           string     `json:"unit" validate:"max=50"`
	MedicineTypeID *uuid.UUID `json:"medicine_type_id,omitempty"`
}

type MedicineResponse struct {
	ID             uuid.UUID  `json:"id"`
	Name           string     `json:"name"`
	Detail         string     `json:"detail"`
	Dosage         string     `json:"dosage"`
	Unit           string     `json:"unit"`
	MedicineTypeID *uuid.UUID `json:"medicine_type_id,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func ToMedicineResponse(m *domain.Medicine) *MedicineResponse {
	return &MedicineResponse{
		ID:             m.ID,
		Name:           m.Name,
		Detail:         m.Detail,
		Dosage:         m.Dosage,
		Unit:           m.Unit,
		MedicineTypeID: m.MedicineTypeID,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func ToMedicineDomain(req *MedicineCreateRequest) *domain.Medicine {
	return &domain.Medicine{
		Name:           req.Name,
		Detail:         req.Detail,
		Dosage:         req.Dosage,
		Unit:           req.Unit,
		MedicineTypeID: req.MedicineTypeID,
	}
}
