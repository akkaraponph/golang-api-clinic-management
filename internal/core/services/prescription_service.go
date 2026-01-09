package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type prescriptionService struct {
	repo ports.PrescriptionRepository
}

func NewPrescriptionService(repo ports.PrescriptionRepository) ports.PrescriptionService {
	return &prescriptionService{repo: repo}
}

func (s *prescriptionService) CreatePrescription(prescription *domain.Prescription) error {
	return s.repo.Create(prescription)
}

func (s *prescriptionService) GetPrescriptionByID(id uuid.UUID) (*domain.Prescription, error) {
	return s.repo.GetByID(id)
}

func (s *prescriptionService) GetAllPrescriptions() ([]*domain.Prescription, error) {
	return s.repo.GetAll()
}

func (s *prescriptionService) GetPrescriptionsByPatientID(patientID uuid.UUID) ([]*domain.Prescription, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *prescriptionService) GetPrescriptionByCourseID(courseID uuid.UUID) (*domain.Prescription, error) {
	return s.repo.GetByCourseID(courseID)
}

func (s *prescriptionService) UpdatePrescription(prescription *domain.Prescription) error {
	return s.repo.Update(prescription)
}

func (s *prescriptionService) DeletePrescription(id uuid.UUID) error {
	return s.repo.Delete(id)
}
