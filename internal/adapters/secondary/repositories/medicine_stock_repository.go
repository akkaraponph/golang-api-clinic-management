package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type medicineStockRepository struct {
	db *gorm.DB
}

func NewMedicineStockRepository(db *gorm.DB) ports.MedicineStockRepository {
	return &medicineStockRepository{db: db}
}

func (r *medicineStockRepository) Create(stock *domain.MedicineStock) error {
	return r.db.Create(stock).Error
}

func (r *medicineStockRepository) GetByID(id uuid.UUID) (*domain.MedicineStock, error) {
	var stock domain.MedicineStock
	err := r.db.Preload("Medicine").Where("id = ?", id).First(&stock).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *medicineStockRepository) GetByMedicineID(medicineID uuid.UUID) ([]*domain.MedicineStock, error) {
	var stocks []*domain.MedicineStock
	err := r.db.Where("medicine_id = ?", medicineID).Find(&stocks).Error
	return stocks, err
}

func (r *medicineStockRepository) GetAll() ([]*domain.MedicineStock, error) {
	var stocks []*domain.MedicineStock
	err := r.db.Preload("Medicine").Find(&stocks).Error
	return stocks, err
}

func (r *medicineStockRepository) Update(stock *domain.MedicineStock) error {
	return r.db.Save(stock).Error
}

func (r *medicineStockRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.MedicineStock{}, id).Error
}

func (r *medicineStockRepository) GetRemainingStock() ([]*domain.MedicineStock, error) {
	var stocks []*domain.MedicineStock
	err := r.db.Preload("Medicine").Where("qty > 0").Find(&stocks).Error
	return stocks, err
}

