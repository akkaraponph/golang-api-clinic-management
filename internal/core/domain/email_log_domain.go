package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmailLogStatus string

const (
	EmailLogStatusSent   EmailLogStatus = "Sent"
	EmailLogStatusFailed EmailLogStatus = "Failed"
)

type EmailLog struct {
	BaseModel
	PatientID uuid.UUID      `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	Email     string         `json:"email" validate:"required,max=255,email" gorm:"size:255"`
	Subject   string         `json:"subject" validate:"required,max=255" gorm:"size:255"`
	Body      string         `json:"body" validate:"required" gorm:"type:text"`
	SentAt    *time.Time     `json:"sent_at"`
	Status    EmailLogStatus `json:"status" validate:"required" gorm:"size:50"`
	
	// Relations
	Patient Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
}

func (el *EmailLog) TableName() string {
	return "email_logs"
}

func (el *EmailLog) BeforeCreate(tx *gorm.DB) (err error) {
	if el.ID == uuidv7.NilUUID() {
		if el.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

