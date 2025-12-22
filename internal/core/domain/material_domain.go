package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Material struct {
	BaseModel
	Name         string    `json:"name" validate:"required,max=255" gorm:"size:255"`
	Detail       string    `json:"detail" gorm:"type:text"`
	Unit         string    `json:"unit" validate:"max=50" gorm:"size:50"`
	PurchasePrice float64  `json:"purchase_price" gorm:"type:decimal(10,2)"`
	Qty          int       `json:"qty" gorm:"default:0"`
	AgentID      uuid.UUID `json:"agent_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Agent Agent `json:"agent,omitempty" gorm:"foreignKey:AgentID"`
}

func (m *Material) TableName() string {
	return "materials"
}

func (m *Material) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuidv7.NilUUID() {
		if m.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

