package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MaterialRepository interface {
	Create(material *domain.Material) error
	GetByID(id uuid.UUID) (*domain.Material, error)
	GetAll() ([]*domain.Material, error)
	Update(material *domain.Material) error
	Delete(id uuid.UUID) error
}

type MaterialService interface {
	CreateMaterial(material *domain.Material) error
	GetMaterialByID(id uuid.UUID) (*domain.Material, error)
	GetAllMaterials() ([]*domain.Material, error)
	UpdateMaterial(material *domain.Material) error
	DeleteMaterial(id uuid.UUID) error
}

type MaterialOrderRepository interface {
	Create(order *domain.MaterialOrder) error
	GetByID(id uuid.UUID) (*domain.MaterialOrder, error)
	GetAll() ([]*domain.MaterialOrder, error)
	Update(order *domain.MaterialOrder) error
	Delete(id uuid.UUID) error
}

type MaterialOrderService interface {
	CreateMaterialOrder(order *domain.MaterialOrder) error
	GetMaterialOrderByID(id uuid.UUID) (*domain.MaterialOrder, error)
	GetAllMaterialOrders() ([]*domain.MaterialOrder, error)
	UpdateMaterialOrder(order *domain.MaterialOrder) error
	DeleteMaterialOrder(id uuid.UUID) error
}

