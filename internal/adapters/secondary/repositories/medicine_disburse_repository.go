package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type medicineDisburseRepository struct {
	db *gorm.DB
}

func NewMedicineDisburseRepository(db *gorm.DB) ports.MedicineDisburseRepository {
	return &medicineDisburseRepository{db: db}
}

func (r *medicineDisburseRepository) Create(disburse *domain.MedicineDisburse) error {
	if err := r.db.Create(disburse).Error; err != nil {
		return err
	}
	// Create details separately as they have composite primary keys
	if len(disburse.MedicineDisburseDetails) > 0 {
		for _, detail := range disburse.MedicineDisburseDetails {
			detail.MedicineDisburseID = disburse.ID
			if err := r.db.Create(&detail).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *medicineDisburseRepository) GetByID(id uuid.UUID) (*domain.MedicineDisburse, error) {
	var disburse domain.MedicineDisburse
	err := r.db.Preload("Prescription").
		Preload("MedicineDisburseDetails").Preload("MedicineDisburseDetails.MedicineStock").
		Where("id = ?", id).First(&disburse).Error
	if err != nil {
		return nil, err
	}
	return &disburse, nil
}

func (r *medicineDisburseRepository) GetAll() ([]*domain.MedicineDisburse, error) {
	var disburses []*domain.MedicineDisburse
	err := r.db.Preload("Prescription").Order("date DESC").Find(&disburses).Error
	return disburses, err
}

func (r *medicineDisburseRepository) GetByPrescriptionID(prescriptionID uuid.UUID) ([]*domain.MedicineDisburse, error) {
	var disburses []*domain.MedicineDisburse
	err := r.db.Where("prescription_id = ?", prescriptionID).Find(&disburses).Error
	return disburses, err
}

func (r *medicineDisburseRepository) Update(disburse *domain.MedicineDisburse) error {
	// Update main record
	if err := r.db.Save(disburse).Error; err != nil {
		return err
	}
	// Delete old details
	if err := r.db.Where("medicine_disburse_id = ?", disburse.ID).
		Delete(&domain.MedicineDisburseDetail{}).Error; err != nil {
		return err
	}
	// Create new details
	if len(disburse.MedicineDisburseDetails) > 0 {
		for _, detail := range disburse.MedicineDisburseDetails {
			detail.MedicineDisburseID = disburse.ID
			if err := r.db.Create(&detail).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *medicineDisburseRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.MedicineDisburse{}, id).Error
}
