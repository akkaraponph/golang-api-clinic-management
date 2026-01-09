package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type materialOrderService struct {
	repo ports.MaterialOrderRepository
}

func NewMaterialOrderService(repo ports.MaterialOrderRepository) ports.MaterialOrderService {
	return &materialOrderService{repo: repo}
}

func (s *materialOrderService) CreateMaterialOrder(order *domain.MaterialOrder) error {
	return s.repo.Create(order)
}

func (s *materialOrderService) GetMaterialOrderByID(id uuid.UUID) (*domain.MaterialOrder, error) {
	return s.repo.GetByID(id)
}

func (s *materialOrderService) GetAllMaterialOrders() ([]*domain.MaterialOrder, error) {
	return s.repo.GetAll()
}

func (s *materialOrderService) UpdateMaterialOrder(order *domain.MaterialOrder) error {
	return s.repo.Update(order)
}

func (s *materialOrderService) DeleteMaterialOrder(id uuid.UUID) error {
	return s.repo.Delete(id)
}
