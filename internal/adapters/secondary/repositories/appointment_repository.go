package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) ports.AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Create(appointment *domain.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetByID(id uuid.UUID) (*domain.Appointment, error) {
	var appointment domain.Appointment
	err := r.db.Preload("Patient").Preload("Employee").
		Where("id = ?", id).First(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *appointmentRepository) GetAll() ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment
	err := r.db.Preload("Patient").Preload("Employee").
		Order("date DESC").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment
	err := r.db.Where("patient_id = ?", patientID).
		Order("date DESC").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) Update(appointment *domain.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *appointmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Appointment{}, id).Error
}

