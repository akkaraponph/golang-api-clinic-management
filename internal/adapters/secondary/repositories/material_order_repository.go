package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type materialOrderRepository struct {
	db *gorm.DB
}

func NewMaterialOrderRepository(db *gorm.DB) ports.MaterialOrderRepository {
	return &materialOrderRepository{db: db}
}

func (r *materialOrderRepository) Create(order *domain.MaterialOrder) error {
	return r.db.Create(order).Error
}

func (r *materialOrderRepository) GetByID(id uuid.UUID) (*domain.MaterialOrder, error) {
	var order domain.MaterialOrder
	err := r.db.Preload("Employee").Preload("Agent").
		Preload("MaterialOrderDetails").Preload("MaterialOrderDetails.Material").
		Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *materialOrderRepository) GetAll() ([]*domain.MaterialOrder, error) {
	var orders []*domain.MaterialOrder
	err := r.db.Preload("Employee").Preload("Agent").
		Order("created_at DESC").Find(&orders).Error
	return orders, err
}

func (r *materialOrderRepository) Update(order *domain.MaterialOrder) error {
	return r.db.Save(order).Error
}

func (r *materialOrderRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.MaterialOrder{}, id).Error
}

