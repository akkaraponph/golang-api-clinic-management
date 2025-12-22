package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type prescriptionRepository struct {
	db *gorm.DB
}

func NewPrescriptionRepository(db *gorm.DB) ports.PrescriptionRepository {
	return &prescriptionRepository{db: db}
}

func (r *prescriptionRepository) Create(prescription *domain.Prescription) error {
	return r.db.Create(prescription).Error
}

func (r *prescriptionRepository) GetByID(id uuid.UUID) (*domain.Prescription, error) {
	var prescription domain.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Doctor.Employee").
		Preload("Course").Where("id = ?", id).First(&prescription).Error
	if err != nil {
		return nil, err
	}
	return &prescription, nil
}

func (r *prescriptionRepository) GetAll() ([]*domain.Prescription, error) {
	var prescriptions []*domain.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Course").
		Order("date DESC").Find(&prescriptions).Error
	return prescriptions, err
}

func (r *prescriptionRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Prescription, error) {
	var prescriptions []*domain.Prescription
	err := r.db.Where("patient_id = ?", patientID).
		Order("date DESC").Find(&prescriptions).Error
	return prescriptions, err
}

func (r *prescriptionRepository) GetByCourseID(courseID uuid.UUID) (*domain.Prescription, error) {
	var prescription domain.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Course").
		Where("course_id = ?", courseID).First(&prescription).Error
	if err != nil {
		return nil, err
	}
	return &prescription, nil
}

func (r *prescriptionRepository) Update(prescription *domain.Prescription) error {
	return r.db.Save(prescription).Error
}

func (r *prescriptionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Prescription{}, id).Error
}

