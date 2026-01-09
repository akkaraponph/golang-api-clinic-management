package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineDisburseRepository interface {
	Create(disburse *domain.MedicineDisburse) error
	GetByID(id uuid.UUID) (*domain.MedicineDisburse, error)
	GetAll() ([]*domain.MedicineDisburse, error)
	GetByPrescriptionID(prescriptionID uuid.UUID) ([]*domain.MedicineDisburse, error)
	Update(disburse *domain.MedicineDisburse) error
	Delete(id uuid.UUID) error
}

type MedicineDisburseService interface {
	CreateMedicineDisburse(disburse *domain.MedicineDisburse) error
	GetMedicineDisburseByID(id uuid.UUID) (*domain.MedicineDisburse, error)
	GetAllMedicineDisburses() ([]*domain.MedicineDisburse, error)
	GetMedicineDisbursesByPrescriptionID(prescriptionID uuid.UUID) ([]*domain.MedicineDisburse, error)
	UpdateMedicineDisburse(disburse *domain.MedicineDisburse) error
	DeleteMedicineDisburse(id uuid.UUID) error
}
