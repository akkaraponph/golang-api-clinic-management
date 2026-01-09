package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type referralRepository struct {
	db *gorm.DB
}

func NewReferralRepository(db *gorm.DB) ports.ReferralRepository {
	return &referralRepository{db: db}
}

func (r *referralRepository) Create(referral *domain.Referral) error {
	return r.db.Create(referral).Error
}

func (r *referralRepository) GetByID(id uuid.UUID) (*domain.Referral, error) {
	var referral domain.Referral
	err := r.db.Preload("Patient").Preload("ReferredDoctor").Preload("ReferredDoctor.Employee").
		Where("id = ?", id).First(&referral).Error
	if err != nil {
		return nil, err
	}
	return &referral, nil
}

func (r *referralRepository) GetAll() ([]*domain.Referral, error) {
	var referrals []*domain.Referral
	err := r.db.Preload("Patient").Preload("ReferredDoctor").Order("referral_date DESC").Find(&referrals).Error
	return referrals, err
}

func (r *referralRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Referral, error) {
	var referrals []*domain.Referral
	err := r.db.Where("patient_id = ?", patientID).Order("referral_date DESC").Find(&referrals).Error
	return referrals, err
}

func (r *referralRepository) Update(referral *domain.Referral) error {
	return r.db.Save(referral).Error
}

func (r *referralRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Referral{}, id).Error
}
