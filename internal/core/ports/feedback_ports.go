package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type FeedbackRepository interface {
	Create(feedback *domain.Feedback) error
	GetByID(id uuid.UUID) (*domain.Feedback, error)
	GetAll() ([]*domain.Feedback, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Feedback, error)
	Update(feedback *domain.Feedback) error
	Delete(id uuid.UUID) error
}

type FeedbackService interface {
	CreateFeedback(feedback *domain.Feedback) error
	GetFeedbackByID(id uuid.UUID) (*domain.Feedback, error)
	GetAllFeedbacks() ([]*domain.Feedback, error)
	GetFeedbacksByPatientID(patientID uuid.UUID) ([]*domain.Feedback, error)
	UpdateFeedback(feedback *domain.Feedback) error
	DeleteFeedback(id uuid.UUID) error
}
