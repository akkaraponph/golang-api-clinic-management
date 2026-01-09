package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MaterialCreateRequest struct {
	Name          string    `json:"name" validate:"required,max=255"`
	Detail        string    `json:"detail"`
	Unit          string    `json:"unit" validate:"max=50"`
	PurchasePrice float64   `json:"purchase_price" validate:"min=0"`
	Qty           int       `json:"qty" validate:"min=0"`
	AgentID       uuid.UUID `json:"agent_id" validate:"required"`
}

type MaterialUpdateRequest struct {
	Name          string     `json:"name" validate:"max=255"`
	Detail        string     `json:"detail"`
	Unit          string     `json:"unit" validate:"max=50"`
	PurchasePrice *float64   `json:"purchase_price,omitempty" validate:"omitempty,min=0"`
	Qty           *int       `json:"qty,omitempty" validate:"omitempty,min=0"`
	AgentID       *uuid.UUID `json:"agent_id,omitempty"`
}

type MaterialResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Detail        string    `json:"detail"`
	Unit          string    `json:"unit"`
	PurchasePrice float64   `json:"purchase_price"`
	Qty           int       `json:"qty"`
	AgentID       uuid.UUID `json:"agent_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToMaterialResponse(m *domain.Material) *MaterialResponse {
	return &MaterialResponse{
		ID:            m.ID,
		Name:          m.Name,
		Detail:        m.Detail,
		Unit:          m.Unit,
		PurchasePrice: m.PurchasePrice,
		Qty:           m.Qty,
		AgentID:       m.AgentID,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

func ToMaterialDomain(req *MaterialCreateRequest) *domain.Material {
	return &domain.Material{
		Name:          req.Name,
		Detail:        req.Detail,
		Unit:          req.Unit,
		PurchasePrice: req.PurchasePrice,
		Qty:           req.Qty,
		AgentID:       req.AgentID,
	}
}
