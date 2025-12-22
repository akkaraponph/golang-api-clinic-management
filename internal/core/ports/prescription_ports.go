package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type PrescriptionRepository interface {
	Create(prescription *domain.Prescription) error
	GetByID(id uuid.UUID) (*domain.Prescription, error)
	GetAll() ([]*domain.Prescription, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Prescription, error)
	GetByCourseID(courseID uuid.UUID) (*domain.Prescription, error)
	Update(prescription *domain.Prescription) error
	Delete(id uuid.UUID) error
}

type PrescriptionService interface {
	CreatePrescription(prescription *domain.Prescription) error
	GetPrescriptionByID(id uuid.UUID) (*domain.Prescription, error)
	GetAllPrescriptions() ([]*domain.Prescription, error)
	GetPrescriptionsByPatientID(patientID uuid.UUID) ([]*domain.Prescription, error)
	GetPrescriptionByCourseID(courseID uuid.UUID) (*domain.Prescription, error)
	UpdatePrescription(prescription *domain.Prescription) error
	DeletePrescription(id uuid.UUID) error
}

