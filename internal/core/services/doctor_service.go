package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type doctorService struct {
	repo ports.DoctorRepository
}

func NewDoctorService(repo ports.DoctorRepository) ports.DoctorService {
	return &doctorService{repo: repo}
}

func (s *doctorService) CreateDoctor(doctor *domain.Doctor) error {
	return s.repo.Create(doctor)
}

func (s *doctorService) GetDoctorByID(id uuid.UUID) (*domain.Doctor, error) {
	return s.repo.GetByID(id)
}

func (s *doctorService) GetDoctorByEmployeeID(employeeID uuid.UUID) (*domain.Doctor, error) {
	return s.repo.GetByEmployeeID(employeeID)
}

func (s *doctorService) GetAllDoctors() ([]*domain.Doctor, error) {
	return s.repo.GetAll()
}

func (s *doctorService) UpdateDoctor(doctor *domain.Doctor) error {
	return s.repo.Update(doctor)
}

func (s *doctorService) DeleteDoctor(id uuid.UUID) error {
	return s.repo.Delete(id)
}
