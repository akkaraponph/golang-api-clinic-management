package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type InsuranceRepository interface {
	Create(insurance *domain.Insurance) error
	GetByID(id uuid.UUID) (*domain.Insurance, error)
	GetAll() ([]*domain.Insurance, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Insurance, error)
	Update(insurance *domain.Insurance) error
	Delete(id uuid.UUID) error
}

type InsuranceService interface {
	CreateInsurance(insurance *domain.Insurance) error
	GetInsuranceByID(id uuid.UUID) (*domain.Insurance, error)
	GetAllInsurances() ([]*domain.Insurance, error)
	GetInsurancesByPatientID(patientID uuid.UUID) ([]*domain.Insurance, error)
	UpdateInsurance(insurance *domain.Insurance) error
	DeleteInsurance(id uuid.UUID) error
}
