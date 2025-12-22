package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppointmentStatus string

const (
	AppointmentStatusPending   AppointmentStatus = "Pending"
	AppointmentStatusCompleted AppointmentStatus = "Completed"
	AppointmentStatusCancelled AppointmentStatus = "Cancelled"
)

type Appointment struct {
	BaseModel
	Subject    string           `json:"subject" validate:"max=255" gorm:"size:255"`
	Detail     string           `json:"detail" gorm:"type:text"`
	Date       time.Time        `json:"date" validate:"required"`
	Status     AppointmentStatus `json:"status" validate:"required" gorm:"size:100"`
	PatientID  uuid.UUID        `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	EmployeeID uuid.UUID        `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Patient  Patient  `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	Employee Employee `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
}

func (a *Appointment) TableName() string {
	return "appointments"
}

func (a *Appointment) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuidv7.NilUUID() {
		if a.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

