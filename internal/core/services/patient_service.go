package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type patientService struct {
	repo ports.PatientRepository
}

func NewPatientService(repo ports.PatientRepository) ports.PatientService {
	return &patientService{repo: repo}
}

func (s *patientService) CreatePatient(patient *domain.Patient) error {
	return s.repo.Create(patient)
}

func (s *patientService) GetPatientByID(id uuid.UUID) (*domain.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *patientService) GetAllPatients() ([]*domain.Patient, error) {
	return s.repo.GetAll()
}

func (s *patientService) SearchPatients(query string) ([]*domain.Patient, error) {
	return s.repo.Search(query)
}

func (s *patientService) UpdatePatient(patient *domain.Patient) error {
	return s.repo.Update(patient)
}

func (s *patientService) DeletePatient(id uuid.UUID) error {
	return s.repo.SoftDelete(id)
}

func (s *patientService) RestorePatient(id uuid.UUID) error {
	return s.repo.Restore(id)
}

func (s *patientService) GetTrashPatients() ([]*domain.Patient, error) {
	return s.repo.GetTrash()
}

