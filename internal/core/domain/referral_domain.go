package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Referral struct {
	BaseModel
	PatientID        uuid.UUID  `json:"patient_id" validate:"required" gorm:"type:uuid;not null;index"`
	ReferredDoctorID uuid.UUID  `json:"referred_doctor_id" validate:"required" gorm:"type:uuid;not null;index"`
	Reason           string     `json:"reason" gorm:"type:text"`
	ReferralDate     time.Time  `json:"referral_date" validate:"required"`
	
	// Relations
	Patient        Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	ReferredDoctor Doctor  `json:"referred_doctor,omitempty" gorm:"foreignKey:ReferredDoctorID"`
}

func (r *Referral) TableName() string {
	return "referrals"
}

func (r *Referral) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuidv7.NilUUID() {
		if r.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

