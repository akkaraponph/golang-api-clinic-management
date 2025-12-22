package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type PatientRepository interface {
	Create(patient *domain.Patient) error
	GetByID(id uuid.UUID) (*domain.Patient, error)
	GetAll() ([]*domain.Patient, error)
	Search(query string) ([]*domain.Patient, error)
	Update(patient *domain.Patient) error
	Delete(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
	Restore(id uuid.UUID) error
	GetTrash() ([]*domain.Patient, error)
}

type PatientService interface {
	CreatePatient(patient *domain.Patient) error
	GetPatientByID(id uuid.UUID) (*domain.Patient, error)
	GetAllPatients() ([]*domain.Patient, error)
	SearchPatients(query string) ([]*domain.Patient, error)
	UpdatePatient(patient *domain.Patient) error
	DeletePatient(id uuid.UUID) error
	RestorePatient(id uuid.UUID) error
	GetTrashPatients() ([]*domain.Patient, error)
}

