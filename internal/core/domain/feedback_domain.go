package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	BaseModel
	PatientID    uuid.UUID `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	FeedbackText string    `json:"feedback_text" validate:"required" gorm:"type:text"`
	
	// Relations
	Patient Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
}

func (f *Feedback) TableName() string {
	return "feedbacks"
}

func (f *Feedback) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuidv7.NilUUID() {
		if f.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

