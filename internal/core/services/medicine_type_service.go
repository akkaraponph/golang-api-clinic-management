package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type medicineTypeService struct {
	repo ports.MedicineTypeRepository
}

func NewMedicineTypeService(repo ports.MedicineTypeRepository) ports.MedicineTypeService {
	return &medicineTypeService{repo: repo}
}

func (s *medicineTypeService) CreateMedicineType(medicineType *domain.MedicineType) error {
	return s.repo.Create(medicineType)
}

func (s *medicineTypeService) GetMedicineTypeByID(id uuid.UUID) (*domain.MedicineType, error) {
	return s.repo.GetByID(id)
}

func (s *medicineTypeService) GetAllMedicineTypes() ([]*domain.MedicineType, error) {
	return s.repo.GetAll()
}

func (s *medicineTypeService) GetMedicineTypesByAgentID(agentID uuid.UUID) ([]*domain.MedicineType, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *medicineTypeService) UpdateMedicineType(medicineType *domain.MedicineType) error {
	return s.repo.Update(medicineType)
}

func (s *medicineTypeService) DeleteMedicineType(id uuid.UUID) error {
	return s.repo.Delete(id)
}
