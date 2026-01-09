package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) ports.DoctorRepository {
	return &doctorRepository{db: db}
}

func (r *doctorRepository) Create(doctor *domain.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *doctorRepository) GetByID(id uuid.UUID) (*domain.Doctor, error) {
	var doctor domain.Doctor
	err := r.db.Preload("Employee").Where("id = ?", id).First(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetByEmployeeID(employeeID uuid.UUID) (*domain.Doctor, error) {
	var doctor domain.Doctor
	err := r.db.Preload("Employee").Where("employee_id = ?", employeeID).First(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetAll() ([]*domain.Doctor, error) {
	var doctors []*domain.Doctor
	err := r.db.Preload("Employee").Find(&doctors).Error
	return doctors, err
}

func (r *doctorRepository) Update(doctor *domain.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *doctorRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Doctor{}, id).Error
}
