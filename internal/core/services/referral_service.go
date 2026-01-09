package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type referralService struct {
	repo ports.ReferralRepository
}

func NewReferralService(repo ports.ReferralRepository) ports.ReferralService {
	return &referralService{repo: repo}
}

func (s *referralService) CreateReferral(referral *domain.Referral) error {
	return s.repo.Create(referral)
}

func (s *referralService) GetReferralByID(id uuid.UUID) (*domain.Referral, error) {
	return s.repo.GetByID(id)
}

func (s *referralService) GetAllReferrals() ([]*domain.Referral, error) {
	return s.repo.GetAll()
}

func (s *referralService) GetReferralsByPatientID(patientID uuid.UUID) ([]*domain.Referral, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *referralService) UpdateReferral(referral *domain.Referral) error {
	return s.repo.Update(referral)
}

func (s *referralService) DeleteReferral(id uuid.UUID) error {
	return s.repo.Delete(id)
}
