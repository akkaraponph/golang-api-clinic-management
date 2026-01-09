package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type materialDisburseService struct {
	repo ports.MaterialDisburseRepository
}

func NewMaterialDisburseService(repo ports.MaterialDisburseRepository) ports.MaterialDisburseService {
	return &materialDisburseService{repo: repo}
}

func (s *materialDisburseService) CreateMaterialDisburse(disburse *domain.MaterialDisburse) error {
	return s.repo.Create(disburse)
}

func (s *materialDisburseService) GetMaterialDisburseByID(id uuid.UUID) (*domain.MaterialDisburse, error) {
	return s.repo.GetByID(id)
}

func (s *materialDisburseService) GetAllMaterialDisburses() ([]*domain.MaterialDisburse, error) {
	return s.repo.GetAll()
}

func (s *materialDisburseService) UpdateMaterialDisburse(disburse *domain.MaterialDisburse) error {
	return s.repo.Update(disburse)
}

func (s *materialDisburseService) DeleteMaterialDisburse(id uuid.UUID) error {
	return s.repo.Delete(id)
}
