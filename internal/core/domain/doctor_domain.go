package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	BaseModel
	EmployeeID       uuid.UUID  `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	Specialization   string     `json:"specialization" validate:"max=100" gorm:"size:100"`
	LicenseNumber    string     `json:"license_number" validate:"max=100" gorm:"size:100"`
	RegistrationDate *time.Time `json:"registration_date"`
	ExperienceYears  *int       `json:"experience_years"`
	
	// Relations
	Employee Employee `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
}

func (d *Doctor) TableName() string {
	return "doctors"
}

func (d *Doctor) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuidv7.NilUUID() {
		if d.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

