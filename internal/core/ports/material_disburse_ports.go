package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MaterialDisburseRepository interface {
	Create(disburse *domain.MaterialDisburse) error
	GetByID(id uuid.UUID) (*domain.MaterialDisburse, error)
	GetAll() ([]*domain.MaterialDisburse, error)
	Update(disburse *domain.MaterialDisburse) error
	Delete(id uuid.UUID) error
}

type MaterialDisburseService interface {
	CreateMaterialDisburse(disburse *domain.MaterialDisburse) error
	GetMaterialDisburseByID(id uuid.UUID) (*domain.MaterialDisburse, error)
	GetAllMaterialDisburses() ([]*domain.MaterialDisburse, error)
	UpdateMaterialDisburse(disburse *domain.MaterialDisburse) error
	DeleteMaterialDisburse(id uuid.UUID) error
}
