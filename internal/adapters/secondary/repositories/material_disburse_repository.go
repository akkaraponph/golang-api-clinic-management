package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type materialDisburseRepository struct {
	db *gorm.DB
}

func NewMaterialDisburseRepository(db *gorm.DB) ports.MaterialDisburseRepository {
	return &materialDisburseRepository{db: db}
}

func (r *materialDisburseRepository) Create(disburse *domain.MaterialDisburse) error {
	if err := r.db.Create(disburse).Error; err != nil {
		return err
	}
	// Create details separately as they have composite primary keys
	if len(disburse.MaterialDisburseDetails) > 0 {
		for _, detail := range disburse.MaterialDisburseDetails {
			detail.MaterialDisburseID = disburse.ID
			if err := r.db.Create(&detail).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *materialDisburseRepository) GetByID(id uuid.UUID) (*domain.MaterialDisburse, error) {
	var disburse domain.MaterialDisburse
	err := r.db.Preload("Employee").
		Preload("MaterialDisburseDetails").Preload("MaterialDisburseDetails.Material").
		Where("id = ?", id).First(&disburse).Error
	if err != nil {
		return nil, err
	}
	return &disburse, nil
}

func (r *materialDisburseRepository) GetAll() ([]*domain.MaterialDisburse, error) {
	var disburses []*domain.MaterialDisburse
	err := r.db.Preload("Employee").Order("date DESC").Find(&disburses).Error
	return disburses, err
}

func (r *materialDisburseRepository) Update(disburse *domain.MaterialDisburse) error {
	// Update main record
	if err := r.db.Save(disburse).Error; err != nil {
		return err
	}
	// Delete old details
	if err := r.db.Where("material_disburse_id = ?", disburse.ID).
		Delete(&domain.MaterialDisburseDetail{}).Error; err != nil {
		return err
	}
	// Create new details
	if len(disburse.MaterialDisburseDetails) > 0 {
		for _, detail := range disburse.MaterialDisburseDetails {
			detail.MaterialDisburseID = disburse.ID
			if err := r.db.Create(&detail).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *materialDisburseRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.MaterialDisburse{}, id).Error
}
