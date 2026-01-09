package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineTypeCreateRequest struct {
	Name    string    `json:"name" validate:"required"`
	Detail  string    `json:"detail"`
	AgentID uuid.UUID `json:"agent_id" validate:"required"`
}

type MedicineTypeUpdateRequest struct {
	Name    string     `json:"name"`
	Detail  string     `json:"detail"`
	AgentID *uuid.UUID `json:"agent_id,omitempty"`
}

type MedicineTypeResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	AgentID   uuid.UUID `json:"agent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToMedicineTypeResponse(mt *domain.MedicineType) *MedicineTypeResponse {
	return &MedicineTypeResponse{
		ID:        mt.ID,
		Name:      mt.Name,
		Detail:    mt.Detail,
		AgentID:   mt.AgentID,
		CreatedAt: mt.CreatedAt,
		UpdatedAt: mt.UpdatedAt,
	}
}

func ToMedicineTypeDomain(req *MedicineTypeCreateRequest) *domain.MedicineType {
	return &domain.MedicineType{
		Name:    req.Name,
		Detail:  req.Detail,
		AgentID: req.AgentID,
	}
}
