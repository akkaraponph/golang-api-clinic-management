package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type ReferralRepository interface {
	Create(referral *domain.Referral) error
	GetByID(id uuid.UUID) (*domain.Referral, error)
	GetAll() ([]*domain.Referral, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Referral, error)
	Update(referral *domain.Referral) error
	Delete(id uuid.UUID) error
}

type ReferralService interface {
	CreateReferral(referral *domain.Referral) error
	GetReferralByID(id uuid.UUID) (*domain.Referral, error)
	GetAllReferrals() ([]*domain.Referral, error)
	GetReferralsByPatientID(patientID uuid.UUID) ([]*domain.Referral, error)
	UpdateReferral(referral *domain.Referral) error
	DeleteReferral(id uuid.UUID) error
}
