package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"gorm.io/gorm"
)

type Agent struct {
	BaseModel
	Name    string `json:"name" validate:"required,max=255" gorm:"size:255"`
	Address string `json:"address" validate:"max=255" gorm:"size:255"`
	Phone   string `json:"phone" validate:"max=20" gorm:"size:20"`
	Email   string `json:"email" validate:"max=255,email" gorm:"size:255"`
}

func (a *Agent) TableName() string {
	return "agents"
}

func (a *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuidv7.NilUUID() {
		var err error
		if a.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

