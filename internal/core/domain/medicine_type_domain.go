package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicineType struct {
	BaseModel
	Name    string    `json:"name" validate:"required" gorm:"size:255"`
	Detail  string    `json:"detail" gorm:"type:text"`
	AgentID uuid.UUID `json:"agent_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Agent Agent `json:"agent,omitempty" gorm:"foreignKey:AgentID"`
}

func (mt *MedicineType) TableName() string {
	return "medicine_types"
}

func (mt *MedicineType) BeforeCreate(tx *gorm.DB) (err error) {
	if mt.ID == uuidv7.NilUUID() {
		var err error
		if mt.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

