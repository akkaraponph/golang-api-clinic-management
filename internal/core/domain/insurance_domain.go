package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Insurance struct {
	BaseModel
	PatientID    uuid.UUID  `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	Provider     string     `json:"provider" validate:"required,max=255" gorm:"size:255"`
	PolicyNumber string     `json:"policy_number" validate:"required,max=100" gorm:"size:100"`
	ValidUntil   *time.Time `json:"valid_until"`
	
	// Relations
	Patient Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
}

func (i *Insurance) TableName() string {
	return "insurances"
}

func (i *Insurance) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuidv7.NilUUID() {
		if i.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

