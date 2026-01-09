package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type agentService struct {
	repo ports.AgentRepository
}

func NewAgentService(repo ports.AgentRepository) ports.AgentService {
	return &agentService{repo: repo}
}

func (s *agentService) CreateAgent(agent *domain.Agent) error {
	return s.repo.Create(agent)
}

func (s *agentService) GetAgentByID(id uuid.UUID) (*domain.Agent, error) {
	return s.repo.GetByID(id)
}

func (s *agentService) GetAllAgents() ([]*domain.Agent, error) {
	return s.repo.GetAll()
}

func (s *agentService) UpdateAgent(agent *domain.Agent) error {
	return s.repo.Update(agent)
}

func (s *agentService) DeleteAgent(id uuid.UUID) error {
	return s.repo.Delete(id)
}
