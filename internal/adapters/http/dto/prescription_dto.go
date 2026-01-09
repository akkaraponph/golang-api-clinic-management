package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type PrescriptionCreateRequest struct {
	PatientID    uuid.UUID `json:"patient_id" validate:"required"`
	DoctorID     uuid.UUID `json:"doctor_id" validate:"required"`
	CourseID     uuid.UUID `json:"course_id" validate:"required"`
	Date         time.Time `json:"date" validate:"required"`
	Instructions string    `json:"instructions"`
}

type PrescriptionUpdateRequest struct {
	PatientID    *uuid.UUID `json:"patient_id,omitempty"`
	DoctorID     *uuid.UUID `json:"doctor_id,omitempty"`
	CourseID     *uuid.UUID `json:"course_id,omitempty"`
	Date         *time.Time `json:"date,omitempty"`
	Instructions string     `json:"instructions"`
}

type PrescriptionResponse struct {
	ID           uuid.UUID `json:"id"`
	PatientID    uuid.UUID `json:"patient_id"`
	DoctorID     uuid.UUID `json:"doctor_id"`
	CourseID     uuid.UUID `json:"course_id"`
	Date         time.Time `json:"date"`
	Instructions string    `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToPrescriptionResponse(p *domain.Prescription) *PrescriptionResponse {
	return &PrescriptionResponse{
		ID:           p.ID,
		PatientID:    p.PatientID,
		DoctorID:     p.DoctorID,
		CourseID:     p.CourseID,
		Date:         p.Date,
		Instructions: p.Instructions,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func ToPrescriptionDomain(req *PrescriptionCreateRequest) *domain.Prescription {
	return &domain.Prescription{
		PatientID:    req.PatientID,
		DoctorID:     req.DoctorID,
		CourseID:     req.CourseID,
		Date:         req.Date,
		Instructions: req.Instructions,
	}
}
