package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineOrderRepository interface {
	Create(order *domain.MedicineOrder) error
	GetByID(id uuid.UUID) (*domain.MedicineOrder, error)
	GetAll() ([]*domain.MedicineOrder, error)
	Update(order *domain.MedicineOrder) error
	Delete(id uuid.UUID) error
}

type MedicineOrderService interface {
	CreateMedicineOrder(order *domain.MedicineOrder) error
	GetMedicineOrderByID(id uuid.UUID) (*domain.MedicineOrder, error)
	GetAllMedicineOrders() ([]*domain.MedicineOrder, error)
	UpdateMedicineOrder(order *domain.MedicineOrder) error
	DeleteMedicineOrder(id uuid.UUID) error
}
