package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type appointmentService struct {
	repo ports.AppointmentRepository
}

func NewAppointmentService(repo ports.AppointmentRepository) ports.AppointmentService {
	return &appointmentService{repo: repo}
}

func (s *appointmentService) CreateAppointment(appointment *domain.Appointment) error {
	return s.repo.Create(appointment)
}

func (s *appointmentService) GetAppointmentByID(id uuid.UUID) (*domain.Appointment, error) {
	return s.repo.GetByID(id)
}

func (s *appointmentService) GetAllAppointments() ([]*domain.Appointment, error) {
	return s.repo.GetAll()
}

func (s *appointmentService) GetAppointmentsByPatientID(patientID uuid.UUID) ([]*domain.Appointment, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *appointmentService) UpdateAppointment(appointment *domain.Appointment) error {
	return s.repo.Update(appointment)
}

func (s *appointmentService) DeleteAppointment(id uuid.UUID) error {
	return s.repo.Delete(id)
}
