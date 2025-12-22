package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Medicine struct {
	BaseModel
	Name        string    `json:"name" validate:"required,max=255" gorm:"size:255"`
	Detail      string    `json:"detail" gorm:"type:text"`
	Dosage      string    `json:"dosage" validate:"max=100" gorm:"size:100"`
	Unit        string    `json:"unit" validate:"max=50" gorm:"size:50"`
	MedicineTypeID *uuid.UUID `json:"medicine_type_id,omitempty" gorm:"type:uuid;index"`
	
	// Relations
	MedicineType *MedicineType `json:"medicine_type,omitempty" gorm:"foreignKey:MedicineTypeID"`
}

func (m *Medicine) TableName() string {
	return "medicines"
}

func (m *Medicine) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuidv7.NilUUID() {
		if m.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

