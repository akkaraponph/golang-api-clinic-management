package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineRepository interface {
	Create(medicine *domain.Medicine) error
	GetByID(id uuid.UUID) (*domain.Medicine, error)
	GetAll() ([]*domain.Medicine, error)
	Search(query string) ([]*domain.Medicine, error)
	Update(medicine *domain.Medicine) error
	Delete(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
	Restore(id uuid.UUID) error
}

type MedicineService interface {
	CreateMedicine(medicine *domain.Medicine) error
	GetMedicineByID(id uuid.UUID) (*domain.Medicine, error)
	GetAllMedicines() ([]*domain.Medicine, error)
	SearchMedicines(query string) ([]*domain.Medicine, error)
	UpdateMedicine(medicine *domain.Medicine) error
	DeleteMedicine(id uuid.UUID) error
	RestoreMedicine(id uuid.UUID) error
}

type MedicineStockRepository interface {
	Create(stock *domain.MedicineStock) error
	GetByID(id uuid.UUID) (*domain.MedicineStock, error)
	GetByMedicineID(medicineID uuid.UUID) ([]*domain.MedicineStock, error)
	GetAll() ([]*domain.MedicineStock, error)
	Update(stock *domain.MedicineStock) error
	Delete(id uuid.UUID) error
	GetRemainingStock() ([]*domain.MedicineStock, error)
}

type MedicineStockService interface {
	CreateStock(stock *domain.MedicineStock) error
	GetStockByID(id uuid.UUID) (*domain.MedicineStock, error)
	GetStockByMedicineID(medicineID uuid.UUID) ([]*domain.MedicineStock, error)
	GetAllStocks() ([]*domain.MedicineStock, error)
	UpdateStock(stock *domain.MedicineStock) error
	DeleteStock(id uuid.UUID) error
	GetRemainingStocks() ([]*domain.MedicineStock, error)
}

