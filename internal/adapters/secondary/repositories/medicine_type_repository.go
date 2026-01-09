package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type medicineTypeRepository struct {
	db *gorm.DB
}

func NewMedicineTypeRepository(db *gorm.DB) ports.MedicineTypeRepository {
	return &medicineTypeRepository{db: db}
}

func (r *medicineTypeRepository) Create(medicineType *domain.MedicineType) error {
	return r.db.Create(medicineType).Error
}

func (r *medicineTypeRepository) GetByID(id uuid.UUID) (*domain.MedicineType, error) {
	var medicineType domain.MedicineType
	err := r.db.Preload("Agent").Where("id = ?", id).First(&medicineType).Error
	if err != nil {
		return nil, err
	}
	return &medicineType, nil
}

func (r *medicineTypeRepository) GetAll() ([]*domain.MedicineType, error) {
	var medicineTypes []*domain.MedicineType
	err := r.db.Preload("Agent").Find(&medicineTypes).Error
	return medicineTypes, err
}

func (r *medicineTypeRepository) GetByAgentID(agentID uuid.UUID) ([]*domain.MedicineType, error) {
	var medicineTypes []*domain.MedicineType
	err := r.db.Where("agent_id = ?", agentID).Find(&medicineTypes).Error
	return medicineTypes, err
}

func (r *medicineTypeRepository) Update(medicineType *domain.MedicineType) error {
	return r.db.Save(medicineType).Error
}

func (r *medicineTypeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.MedicineType{}, id).Error
}
