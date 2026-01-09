package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type DoctorRepository interface {
	Create(doctor *domain.Doctor) error
	GetByID(id uuid.UUID) (*domain.Doctor, error)
	GetByEmployeeID(employeeID uuid.UUID) (*domain.Doctor, error)
	GetAll() ([]*domain.Doctor, error)
	Update(doctor *domain.Doctor) error
	Delete(id uuid.UUID) error
}

type DoctorService interface {
	CreateDoctor(doctor *domain.Doctor) error
	GetDoctorByID(id uuid.UUID) (*domain.Doctor, error)
	GetDoctorByEmployeeID(employeeID uuid.UUID) (*domain.Doctor, error)
	GetAllDoctors() ([]*domain.Doctor, error)
	UpdateDoctor(doctor *domain.Doctor) error
	DeleteDoctor(id uuid.UUID) error
}
