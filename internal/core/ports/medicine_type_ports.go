package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineTypeRepository interface {
	Create(medicineType *domain.MedicineType) error
	GetByID(id uuid.UUID) (*domain.MedicineType, error)
	GetAll() ([]*domain.MedicineType, error)
	GetByAgentID(agentID uuid.UUID) ([]*domain.MedicineType, error)
	Update(medicineType *domain.MedicineType) error
	Delete(id uuid.UUID) error
}

type MedicineTypeService interface {
	CreateMedicineType(medicineType *domain.MedicineType) error
	GetMedicineTypeByID(id uuid.UUID) (*domain.MedicineType, error)
	GetAllMedicineTypes() ([]*domain.MedicineType, error)
	GetMedicineTypesByAgentID(agentID uuid.UUID) ([]*domain.MedicineType, error)
	UpdateMedicineType(medicineType *domain.MedicineType) error
	DeleteMedicineType(id uuid.UUID) error
}
