package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type insuranceService struct {
	repo ports.InsuranceRepository
}

func NewInsuranceService(repo ports.InsuranceRepository) ports.InsuranceService {
	return &insuranceService{repo: repo}
}

func (s *insuranceService) CreateInsurance(insurance *domain.Insurance) error {
	return s.repo.Create(insurance)
}

func (s *insuranceService) GetInsuranceByID(id uuid.UUID) (*domain.Insurance, error) {
	return s.repo.GetByID(id)
}

func (s *insuranceService) GetAllInsurances() ([]*domain.Insurance, error) {
	return s.repo.GetAll()
}

func (s *insuranceService) GetInsurancesByPatientID(patientID uuid.UUID) ([]*domain.Insurance, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *insuranceService) UpdateInsurance(insurance *domain.Insurance) error {
	return s.repo.Update(insurance)
}

func (s *insuranceService) DeleteInsurance(id uuid.UUID) error {
	return s.repo.Delete(id)
}
