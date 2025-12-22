package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicineStock struct {
	BaseModel
	MedicineID  uuid.UUID  `json:"medicine_id" validate:"required" gorm:"type:uuid;not null;index"`
	Qty         int        `json:"qty" validate:"required,min=0"`
	ExpiredDate *time.Time `json:"expired_date"`
	
	// Relations
	Medicine Medicine `json:"medicine,omitempty" gorm:"foreignKey:MedicineID"`
}

func (ms *MedicineStock) TableName() string {
	return "medicine_stocks"
}

func (ms *MedicineStock) BeforeCreate(tx *gorm.DB) (err error) {
	if ms.ID == uuidv7.NilUUID() {
		if ms.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

