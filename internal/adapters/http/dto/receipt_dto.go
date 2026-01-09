package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type ReceiptCreateRequest struct {
	Date               time.Time  `json:"date" validate:"required"`
	TotalPrice         float64    `json:"total_price" validate:"required,min=0"`
	CourseID           *uuid.UUID `json:"course_id,omitempty"`
	MedicineDisburseID *uuid.UUID `json:"medicine_disburse_id,omitempty"`
	PatientID          uuid.UUID  `json:"patient_id" validate:"required"`
	EmployeeID         uuid.UUID  `json:"employee_id" validate:"required"`
}

type ReceiptUpdateRequest struct {
	Date               *time.Time `json:"date,omitempty"`
	TotalPrice         *float64   `json:"total_price,omitempty" validate:"omitempty,min=0"`
	CourseID           *uuid.UUID `json:"course_id,omitempty"`
	MedicineDisburseID *uuid.UUID `json:"medicine_disburse_id,omitempty"`
	PatientID          *uuid.UUID `json:"patient_id,omitempty"`
	EmployeeID         *uuid.UUID `json:"employee_id,omitempty"`
}

type ReceiptResponse struct {
	ID                 uuid.UUID  `json:"id"`
	Date               time.Time  `json:"date"`
	TotalPrice         float64    `json:"total_price"`
	CourseID           *uuid.UUID `json:"course_id,omitempty"`
	MedicineDisburseID *uuid.UUID `json:"medicine_disburse_id,omitempty"`
	PatientID          uuid.UUID  `json:"patient_id"`
	EmployeeID         uuid.UUID  `json:"employee_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

func ToReceiptResponse(r *domain.Receipt) *ReceiptResponse {
	return &ReceiptResponse{
		ID:                 r.ID,
		Date:               r.Date,
		TotalPrice:         r.TotalPrice,
		CourseID:           r.CourseID,
		MedicineDisburseID: r.MedicineDisburseID,
		PatientID:          r.PatientID,
		EmployeeID:         r.EmployeeID,
		CreatedAt:          r.CreatedAt,
		UpdatedAt:          r.UpdatedAt,
	}
}

func ToReceiptDomain(req *ReceiptCreateRequest) *domain.Receipt {
	return &domain.Receipt{
		Date:               req.Date,
		TotalPrice:         req.TotalPrice,
		CourseID:           req.CourseID,
		MedicineDisburseID: req.MedicineDisburseID,
		PatientID:          req.PatientID,
		EmployeeID:         req.EmployeeID,
	}
}
