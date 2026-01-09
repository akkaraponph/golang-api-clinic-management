package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type medicineOrderService struct {
	repo ports.MedicineOrderRepository
}

func NewMedicineOrderService(repo ports.MedicineOrderRepository) ports.MedicineOrderService {
	return &medicineOrderService{repo: repo}
}

func (s *medicineOrderService) CreateMedicineOrder(order *domain.MedicineOrder) error {
	return s.repo.Create(order)
}

func (s *medicineOrderService) GetMedicineOrderByID(id uuid.UUID) (*domain.MedicineOrder, error) {
	return s.repo.GetByID(id)
}

func (s *medicineOrderService) GetAllMedicineOrders() ([]*domain.MedicineOrder, error) {
	return s.repo.GetAll()
}

func (s *medicineOrderService) UpdateMedicineOrder(order *domain.MedicineOrder) error {
	return s.repo.Update(order)
}

func (s *medicineOrderService) DeleteMedicineOrder(id uuid.UUID) error {
	return s.repo.Delete(id)
}
