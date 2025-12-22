package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type medicineService struct {
	repo ports.MedicineRepository
}

func NewMedicineService(repo ports.MedicineRepository) ports.MedicineService {
	return &medicineService{repo: repo}
}

func (s *medicineService) CreateMedicine(medicine *domain.Medicine) error {
	return s.repo.Create(medicine)
}

func (s *medicineService) GetMedicineByID(id uuid.UUID) (*domain.Medicine, error) {
	return s.repo.GetByID(id)
}

func (s *medicineService) GetAllMedicines() ([]*domain.Medicine, error) {
	return s.repo.GetAll()
}

func (s *medicineService) SearchMedicines(query string) ([]*domain.Medicine, error) {
	return s.repo.Search(query)
}

func (s *medicineService) UpdateMedicine(medicine *domain.Medicine) error {
	return s.repo.Update(medicine)
}

func (s *medicineService) DeleteMedicine(id uuid.UUID) error {
	return s.repo.SoftDelete(id)
}

func (s *medicineService) RestoreMedicine(id uuid.UUID) error {
	return s.repo.Restore(id)
}

