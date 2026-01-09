package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type FeedbackCreateRequest struct {
	PatientID    uuid.UUID `json:"patient_id" validate:"required"`
	FeedbackText string    `json:"feedback_text" validate:"required"`
}

type FeedbackUpdateRequest struct {
	PatientID    *uuid.UUID `json:"patient_id,omitempty"`
	FeedbackText string     `json:"feedback_text"`
}

type FeedbackResponse struct {
	ID           uuid.UUID `json:"id"`
	PatientID    uuid.UUID `json:"patient_id"`
	FeedbackText string    `json:"feedback_text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToFeedbackResponse(f *domain.Feedback) *FeedbackResponse {
	return &FeedbackResponse{
		ID:           f.ID,
		PatientID:    f.PatientID,
		FeedbackText: f.FeedbackText,
		CreatedAt:    f.CreatedAt,
		UpdatedAt:    f.UpdatedAt,
	}
}

func ToFeedbackDomain(req *FeedbackCreateRequest) *domain.Feedback {
	return &domain.Feedback{
		PatientID:    req.PatientID,
		FeedbackText: req.FeedbackText,
	}
}
