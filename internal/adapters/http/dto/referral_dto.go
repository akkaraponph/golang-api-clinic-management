package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type ReferralCreateRequest struct {
	PatientID        uuid.UUID `json:"patient_id" validate:"required"`
	ReferredDoctorID uuid.UUID `json:"referred_doctor_id" validate:"required"`
	Reason           string    `json:"reason"`
	ReferralDate     time.Time `json:"referral_date" validate:"required"`
}

type ReferralUpdateRequest struct {
	PatientID        *uuid.UUID `json:"patient_id,omitempty"`
	ReferredDoctorID *uuid.UUID `json:"referred_doctor_id,omitempty"`
	Reason           string     `json:"reason"`
	ReferralDate     *time.Time `json:"referral_date,omitempty"`
}

type ReferralResponse struct {
	ID               uuid.UUID `json:"id"`
	PatientID        uuid.UUID `json:"patient_id"`
	ReferredDoctorID uuid.UUID `json:"referred_doctor_id"`
	Reason           string    `json:"reason"`
	ReferralDate     time.Time `json:"referral_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func ToReferralResponse(r *domain.Referral) *ReferralResponse {
	return &ReferralResponse{
		ID:               r.ID,
		PatientID:        r.PatientID,
		ReferredDoctorID: r.ReferredDoctorID,
		Reason:           r.Reason,
		ReferralDate:     r.ReferralDate,
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
	}
}

func ToReferralDomain(req *ReferralCreateRequest) *domain.Referral {
	return &domain.Referral{
		PatientID:        req.PatientID,
		ReferredDoctorID: req.ReferredDoctorID,
		Reason:           req.Reason,
		ReferralDate:     req.ReferralDate,
	}
}
