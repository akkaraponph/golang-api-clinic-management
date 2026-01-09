package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type insuranceRepository struct {
	db *gorm.DB
}

func NewInsuranceRepository(db *gorm.DB) ports.InsuranceRepository {
	return &insuranceRepository{db: db}
}

func (r *insuranceRepository) Create(insurance *domain.Insurance) error {
	return r.db.Create(insurance).Error
}

func (r *insuranceRepository) GetByID(id uuid.UUID) (*domain.Insurance, error) {
	var insurance domain.Insurance
	err := r.db.Preload("Patient").Where("id = ?", id).First(&insurance).Error
	if err != nil {
		return nil, err
	}
	return &insurance, nil
}

func (r *insuranceRepository) GetAll() ([]*domain.Insurance, error) {
	var insurances []*domain.Insurance
	err := r.db.Preload("Patient").Find(&insurances).Error
	return insurances, err
}

func (r *insuranceRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Insurance, error) {
	var insurances []*domain.Insurance
	err := r.db.Where("patient_id = ?", patientID).Find(&insurances).Error
	return insurances, err
}

func (r *insuranceRepository) Update(insurance *domain.Insurance) error {
	return r.db.Save(insurance).Error
}

func (r *insuranceRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Insurance{}, id).Error
}
