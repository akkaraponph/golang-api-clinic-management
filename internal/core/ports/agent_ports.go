package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type AgentRepository interface {
	Create(agent *domain.Agent) error
	GetByID(id uuid.UUID) (*domain.Agent, error)
	GetAll() ([]*domain.Agent, error)
	Update(agent *domain.Agent) error
	Delete(id uuid.UUID) error
}

type AgentService interface {
	CreateAgent(agent *domain.Agent) error
	GetAgentByID(id uuid.UUID) (*domain.Agent, error)
	GetAllAgents() ([]*domain.Agent, error)
	UpdateAgent(agent *domain.Agent) error
	DeleteAgent(id uuid.UUID) error
}
