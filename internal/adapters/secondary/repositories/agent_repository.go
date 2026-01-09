package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type agentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) ports.AgentRepository {
	return &agentRepository{db: db}
}

func (r *agentRepository) Create(agent *domain.Agent) error {
	return r.db.Create(agent).Error
}

func (r *agentRepository) GetByID(id uuid.UUID) (*domain.Agent, error) {
	var agent domain.Agent
	err := r.db.Where("id = ?", id).First(&agent).Error
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *agentRepository) GetAll() ([]*domain.Agent, error) {
	var agents []*domain.Agent
	err := r.db.Find(&agents).Error
	return agents, err
}

func (r *agentRepository) Update(agent *domain.Agent) error {
	return r.db.Save(agent).Error
}

func (r *agentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Agent{}, id).Error
}
