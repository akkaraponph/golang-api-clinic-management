package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type medicineDisburseService struct {
	repo ports.MedicineDisburseRepository
}

func NewMedicineDisburseService(repo ports.MedicineDisburseRepository) ports.MedicineDisburseService {
	return &medicineDisburseService{repo: repo}
}

func (s *medicineDisburseService) CreateMedicineDisburse(disburse *domain.MedicineDisburse) error {
	return s.repo.Create(disburse)
}

func (s *medicineDisburseService) GetMedicineDisburseByID(id uuid.UUID) (*domain.MedicineDisburse, error) {
	return s.repo.GetByID(id)
}

func (s *medicineDisburseService) GetAllMedicineDisburses() ([]*domain.MedicineDisburse, error) {
	return s.repo.GetAll()
}

func (s *medicineDisburseService) GetMedicineDisbursesByPrescriptionID(prescriptionID uuid.UUID) ([]*domain.MedicineDisburse, error) {
	return s.repo.GetByPrescriptionID(prescriptionID)
}

func (s *medicineDisburseService) UpdateMedicineDisburse(disburse *domain.MedicineDisburse) error {
	return s.repo.Update(disburse)
}

func (s *medicineDisburseService) DeleteMedicineDisburse(id uuid.UUID) error {
	return s.repo.Delete(id)
}
