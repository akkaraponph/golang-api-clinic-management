package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type medicineOrderRepository struct {
	db *gorm.DB
}

func NewMedicineOrderRepository(db *gorm.DB) ports.MedicineOrderRepository {
	return &medicineOrderRepository{db: db}
}

func (r *medicineOrderRepository) Create(order *domain.MedicineOrder) error {
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	// Create details separately as they have composite primary keys
	if len(order.MedicineOrderDetails) > 0 {
		for _, detail := range order.MedicineOrderDetails {
			detail.MedicineOrderID = order.ID
			if err := r.db.Create(&detail).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *medicineOrderRepository) GetByID(id uuid.UUID) (*domain.MedicineOrder, error) {
	var order domain.MedicineOrder
	err := r.db.Preload("Employee").Preload("Agent").
		Preload("MedicineOrderDetails").Preload("MedicineOrderDetails.Medicine").
		Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *medicineOrderRepository) GetAll() ([]*domain.MedicineOrder, error) {
	var orders []*domain.MedicineOrder
	err := r.db.Preload("Employee").Preload("Agent").
		Order("date DESC").Find(&orders).Error
	return orders, err
}

func (r *medicineOrderRepository) Update(order *domain.MedicineOrder) error {
	// Update main record
	if err := r.db.Save(order).Error; err != nil {
		return err
	}
	// Delete old details
	if err := r.db.Where("medicine_order_id = ?", order.ID).
		Delete(&domain.MedicineOrderDetail{}).Error; err != nil {
		return err
	}
	// Create new details
	if len(order.MedicineOrderDetails) > 0 {
		for _, detail := range order.MedicineOrderDetails {
			detail.MedicineOrderID = order.ID
			if err := r.db.Create(&detail).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *medicineOrderRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.MedicineOrder{}, id).Error
}
