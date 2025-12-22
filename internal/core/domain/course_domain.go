package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseStatus string

const (
	CourseStatusPending   CourseStatus = "Pending"
	CourseStatusCompleted CourseStatus = "Completed"
	CourseStatusCancelled CourseStatus = "Cancelled"
)

type Course struct {
	BaseModel
	Date       time.Time   `json:"date" validate:"required"`
	Weight     *float64    `json:"weight" gorm:"type:decimal(10,2)"`
	Height     *float64    `json:"height" gorm:"type:decimal(10,2)"`
	Systolic   *int        `json:"systolic"`
	Diastolic  *int        `json:"diastolic"`
	HeartRate  *int        `json:"heart_rate"`
	Diagnose   string      `json:"diagnose" gorm:"type:text"`
	TotalPrice float64     `json:"total_price" gorm:"type:decimal(10,2)"`
	Status     CourseStatus `json:"status" validate:"required" gorm:"size:100"`
	PatientID  uuid.UUID   `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	EmployeeID uuid.UUID   `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Patient  Patient  `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	Employee Employee `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
}

func (c *Course) TableName() string {
	return "courses"
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuidv7.NilUUID() {
		if c.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

