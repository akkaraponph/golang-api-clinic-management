package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type feedbackService struct {
	repo ports.FeedbackRepository
}

func NewFeedbackService(repo ports.FeedbackRepository) ports.FeedbackService {
	return &feedbackService{repo: repo}
}

func (s *feedbackService) CreateFeedback(feedback *domain.Feedback) error {
	return s.repo.Create(feedback)
}

func (s *feedbackService) GetFeedbackByID(id uuid.UUID) (*domain.Feedback, error) {
	return s.repo.GetByID(id)
}

func (s *feedbackService) GetAllFeedbacks() ([]*domain.Feedback, error) {
	return s.repo.GetAll()
}

func (s *feedbackService) GetFeedbacksByPatientID(patientID uuid.UUID) ([]*domain.Feedback, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *feedbackService) UpdateFeedback(feedback *domain.Feedback) error {
	return s.repo.Update(feedback)
}

func (s *feedbackService) DeleteFeedback(id uuid.UUID) error {
	return s.repo.Delete(id)
}
