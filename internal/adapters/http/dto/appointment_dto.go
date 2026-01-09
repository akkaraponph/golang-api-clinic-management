package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type AppointmentCreateRequest struct {
	Subject    string    `json:"subject" validate:"max=255"`
	Detail     string    `json:"detail"`
	Date       time.Time `json:"date" validate:"required"`
	Status     string    `json:"status" validate:"required"`
	PatientID  uuid.UUID `json:"patient_id" validate:"required"`
	EmployeeID uuid.UUID `json:"employee_id" validate:"required"`
}

type AppointmentUpdateRequest struct {
	Subject    string     `json:"subject" validate:"max=255"`
	Detail     string     `json:"detail"`
	Date       *time.Time `json:"date,omitempty"`
	Status     string     `json:"status"`
	PatientID  *uuid.UUID `json:"patient_id,omitempty"`
	EmployeeID *uuid.UUID `json:"employee_id,omitempty"`
}

type AppointmentResponse struct {
	ID         uuid.UUID `json:"id"`
	Subject    string    `json:"subject"`
	Detail     string    `json:"detail"`
	Date       time.Time `json:"date"`
	Status     string    `json:"status"`
	PatientID  uuid.UUID `json:"patient_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToAppointmentResponse(a *domain.Appointment) *AppointmentResponse {
	return &AppointmentResponse{
		ID:         a.ID,
		Subject:    a.Subject,
		Detail:     a.Detail,
		Date:       a.Date,
		Status:     string(a.Status),
		PatientID:  a.PatientID,
		EmployeeID: a.EmployeeID,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

func ToAppointmentDomain(req *AppointmentCreateRequest) *domain.Appointment {
	return &domain.Appointment{
		Subject:    req.Subject,
		Detail:     req.Detail,
		Date:       req.Date,
		Status:     domain.AppointmentStatus(req.Status),
		PatientID:  req.PatientID,
		EmployeeID: req.EmployeeID,
	}
}
