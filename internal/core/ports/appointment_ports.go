package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type AppointmentRepository interface {
	Create(appointment *domain.Appointment) error
	GetByID(id uuid.UUID) (*domain.Appointment, error)
	GetAll() ([]*domain.Appointment, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Appointment, error)
	Update(appointment *domain.Appointment) error
	Delete(id uuid.UUID) error
}

type AppointmentService interface {
	CreateAppointment(appointment *domain.Appointment) error
	GetAppointmentByID(id uuid.UUID) (*domain.Appointment, error)
	GetAllAppointments() ([]*domain.Appointment, error)
	GetAppointmentsByPatientID(patientID uuid.UUID) ([]*domain.Appointment, error)
	UpdateAppointment(appointment *domain.Appointment) error
	DeleteAppointment(id uuid.UUID) error
}

