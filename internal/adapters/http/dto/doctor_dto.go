package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type DoctorCreateRequest struct {
	EmployeeID      uuid.UUID  `json:"employee_id" validate:"required"`
	Specialization  string     `json:"specialization" validate:"max=100"`
	LicenseNumber   string     `json:"license_number" validate:"max=100"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
	ExperienceYears  *int       `json:"experience_years,omitempty"`
}

type DoctorUpdateRequest struct {
	EmployeeID       *uuid.UUID `json:"employee_id,omitempty"`
	Specialization   string     `json:"specialization" validate:"max=100"`
	LicenseNumber    string     `json:"license_number" validate:"max=100"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
	ExperienceYears  *int       `json:"experience_years,omitempty"`
}

type DoctorResponse struct {
	ID              uuid.UUID  `json:"id"`
	EmployeeID      uuid.UUID  `json:"employee_id"`
	Specialization  string     `json:"specialization"`
	LicenseNumber   string     `json:"license_number"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
	ExperienceYears  *int       `json:"experience_years,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func ToDoctorResponse(d *domain.Doctor) *DoctorResponse {
	return &DoctorResponse{
		ID:              d.ID,
		EmployeeID:      d.EmployeeID,
		Specialization:  d.Specialization,
		LicenseNumber:   d.LicenseNumber,
		RegistrationDate: d.RegistrationDate,
		ExperienceYears:  d.ExperienceYears,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

func ToDoctorDomain(req *DoctorCreateRequest) *domain.Doctor {
	return &domain.Doctor{
		EmployeeID:      req.EmployeeID,
		Specialization:  req.Specialization,
		LicenseNumber:   req.LicenseNumber,
		RegistrationDate: req.RegistrationDate,
		ExperienceYears:  req.ExperienceYears,
	}
}
