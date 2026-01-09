package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type InsuranceCreateRequest struct {
	PatientID    uuid.UUID  `json:"patient_id" validate:"required"`
	Provider     string     `json:"provider" validate:"required,max=255"`
	PolicyNumber string     `json:"policy_number" validate:"required,max=100"`
	ValidUntil   *time.Time `json:"valid_until,omitempty"`
}

type InsuranceUpdateRequest struct {
	PatientID    *uuid.UUID `json:"patient_id,omitempty"`
	Provider     string     `json:"provider" validate:"max=255"`
	PolicyNumber string     `json:"policy_number" validate:"max=100"`
	ValidUntil   *time.Time `json:"valid_until,omitempty"`
}

type InsuranceResponse struct {
	ID           uuid.UUID  `json:"id"`
	PatientID    uuid.UUID  `json:"patient_id"`
	Provider     string     `json:"provider"`
	PolicyNumber string     `json:"policy_number"`
	ValidUntil   *time.Time `json:"valid_until,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func ToInsuranceResponse(i *domain.Insurance) *InsuranceResponse {
	return &InsuranceResponse{
		ID:           i.ID,
		PatientID:    i.PatientID,
		Provider:     i.Provider,
		PolicyNumber: i.PolicyNumber,
		ValidUntil:   i.ValidUntil,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}

func ToInsuranceDomain(req *InsuranceCreateRequest) *domain.Insurance {
	return &domain.Insurance{
		PatientID:    req.PatientID,
		Provider:     req.Provider,
		PolicyNumber: req.PolicyNumber,
		ValidUntil:   req.ValidUntil,
	}
}
