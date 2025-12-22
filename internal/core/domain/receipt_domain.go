package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Receipt struct {
	BaseModel
	Date               time.Time  `json:"date" validate:"required"`
	TotalPrice         float64    `json:"total_price" gorm:"type:decimal(10,2)"`
	CourseID           *uuid.UUID `json:"course_id,omitempty" gorm:"type:uuid;index"`
	MedicineDisburseID *uuid.UUID `json:"medicine_disburse_id,omitempty" gorm:"type:uuid;index"`
	PatientID          uuid.UUID  `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	EmployeeID         uuid.UUID  `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Course           *Course           `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	MedicineDisburse *MedicineDisburse `json:"medicine_disburse,omitempty" gorm:"foreignKey:MedicineDisburseID"`
	Patient          Patient           `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	Employee         Employee          `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
}

func (r *Receipt) TableName() string {
	return "receipts"
}

func (r *Receipt) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuidv7.NilUUID() {
		if r.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

