package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Prescription struct {
	BaseModel
	PatientID    uuid.UUID  `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	DoctorID     uuid.UUID  `json:"doctor_id" validate:"required" gorm:"type:uuid;not null;index"`
	CourseID     uuid.UUID  `json:"course_id" validate:"required" gorm:"type:uuid;not null;index"`
	Date         time.Time  `json:"date" validate:"required"`
	Instructions string     `json:"instructions" gorm:"type:text"`
	
	// Relations
	Patient Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	Doctor  Doctor  `json:"doctor,omitempty" gorm:"foreignKey:DoctorID"`
	Course  Course  `json:"course,omitempty" gorm:"foreignKey:CourseID"`
}

func (p *Prescription) TableName() string {
	return "prescriptions"
}

func (p *Prescription) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuidv7.NilUUID() {
		if p.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

