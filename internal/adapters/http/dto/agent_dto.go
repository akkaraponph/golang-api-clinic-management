package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
)

type AgentCreateRequest struct {
	Name    string `json:"name" validate:"required,max=255"`
	Address string `json:"address" validate:"max=255"`
	Phone   string `json:"phone" validate:"max=20"`
	Email   string `json:"email" validate:"max=255,email"`
}

type AgentUpdateRequest struct {
	Name    string `json:"name" validate:"max=255"`
	Address string `json:"address" validate:"max=255"`
	Phone   string `json:"phone" validate:"max=20"`
	Email   string `json:"email" validate:"max=255,email"`
}

type AgentResponse struct {
	ID        interface{} `json:"id"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func ToAgentResponse(a *domain.Agent) *AgentResponse {
	return &AgentResponse{
		ID:        a.ID,
		Name:      a.Name,
		Address:   a.Address,
		Phone:     a.Phone,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ToAgentDomain(req *AgentCreateRequest) *domain.Agent {
	return &domain.Agent{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}
}
