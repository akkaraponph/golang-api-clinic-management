package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type materialService struct {
	repo ports.MaterialRepository
}

func NewMaterialService(repo ports.MaterialRepository) ports.MaterialService {
	return &materialService{repo: repo}
}

func (s *materialService) CreateMaterial(material *domain.Material) error {
	return s.repo.Create(material)
}

func (s *materialService) GetMaterialByID(id uuid.UUID) (*domain.Material, error) {
	return s.repo.GetByID(id)
}

func (s *materialService) GetAllMaterials() ([]*domain.Material, error) {
	return s.repo.GetAll()
}

func (s *materialService) UpdateMaterial(material *domain.Material) error {
	return s.repo.Update(material)
}

func (s *materialService) DeleteMaterial(id uuid.UUID) error {
	return s.repo.Delete(id)
}
