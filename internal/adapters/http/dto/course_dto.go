package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type CourseCreateRequest struct {
	Date       time.Time `json:"date" validate:"required"`
	Weight     *float64  `json:"weight"`
	Height     *float64  `json:"height"`
	Systolic   *int      `json:"systolic"`
	Diastolic  *int      `json:"diastolic"`
	HeartRate  *int      `json:"heart_rate"`
	Diagnose   string    `json:"diagnose"`
	TotalPrice float64   `json:"total_price" validate:"required,min=0"`
	Status     string    `json:"status" validate:"required"`
	PatientID  uuid.UUID `json:"patient_id" validate:"required"`
	EmployeeID uuid.UUID `json:"employee_id" validate:"required"`
}

type CourseUpdateRequest struct {
	Date       *time.Time `json:"date,omitempty"`
	Weight     *float64   `json:"weight,omitempty"`
	Height     *float64   `json:"height,omitempty"`
	Systolic   *int       `json:"systolic,omitempty"`
	Diastolic  *int       `json:"diastolic,omitempty"`
	HeartRate  *int       `json:"heart_rate,omitempty"`
	Diagnose   string     `json:"diagnose,omitempty"`
	TotalPrice *float64   `json:"total_price,omitempty" validate:"omitempty,min=0"`
	Status     string     `json:"status,omitempty"`
	PatientID  *uuid.UUID `json:"patient_id,omitempty"`
	EmployeeID *uuid.UUID `json:"employee_id,omitempty"`
}

type CourseResponse struct {
	ID         uuid.UUID `json:"id"`
	Date       time.Time `json:"date"`
	Weight     *float64  `json:"weight"`
	Height     *float64  `json:"height"`
	Systolic   *int      `json:"systolic"`
	Diastolic  *int      `json:"diastolic"`
	HeartRate  *int      `json:"heart_rate"`
	Diagnose   string    `json:"diagnose"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	PatientID  uuid.UUID `json:"patient_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToCourseResponse(c *domain.Course) *CourseResponse {
	return &CourseResponse{
		ID:         c.ID,
		Date:       c.Date,
		Weight:     c.Weight,
		Height:     c.Height,
		Systolic:   c.Systolic,
		Diastolic:  c.Diastolic,
		HeartRate:  c.HeartRate,
		Diagnose:   c.Diagnose,
		TotalPrice: c.TotalPrice,
		Status:     string(c.Status),
		PatientID:  c.PatientID,
		EmployeeID: c.EmployeeID,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func ToCourseDomain(req *CourseCreateRequest) *domain.Course {
	return &domain.Course{
		Date:       req.Date,
		Weight:     req.Weight,
		Height:     req.Height,
		Systolic:   req.Systolic,
		Diastolic:  req.Diastolic,
		HeartRate:  req.HeartRate,
		Diagnose:   req.Diagnose,
		TotalPrice: req.TotalPrice,
		Status:     domain.CourseStatus(req.Status),
		PatientID:  req.PatientID,
		EmployeeID: req.EmployeeID,
	}
}
