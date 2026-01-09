package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) ports.FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) Create(feedback *domain.Feedback) error {
	return r.db.Create(feedback).Error
}

func (r *feedbackRepository) GetByID(id uuid.UUID) (*domain.Feedback, error) {
	var feedback domain.Feedback
	err := r.db.Preload("Patient").Where("id = ?", id).First(&feedback).Error
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

func (r *feedbackRepository) GetAll() ([]*domain.Feedback, error) {
	var feedbacks []*domain.Feedback
	err := r.db.Preload("Patient").Order("created_at DESC").Find(&feedbacks).Error
	return feedbacks, err
}

func (r *feedbackRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Feedback, error) {
	var feedbacks []*domain.Feedback
	err := r.db.Where("patient_id = ?", patientID).Order("created_at DESC").Find(&feedbacks).Error
	return feedbacks, err
}

func (r *feedbackRepository) Update(feedback *domain.Feedback) error {
	return r.db.Save(feedback).Error
}

func (r *feedbackRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Feedback{}, id).Error
}
