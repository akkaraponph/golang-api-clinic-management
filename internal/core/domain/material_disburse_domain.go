package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialDisburse struct {
	BaseModel
	Date       time.Time `json:"date" validate:"required"`
	EmployeeID uuid.UUID `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Employee                Employee                  `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
	MaterialDisburseDetails []MaterialDisburseDetail  `json:"material_disburse_details,omitempty" gorm:"foreignKey:MaterialDisburseID"`
}

func (md *MaterialDisburse) TableName() string {
	return "material_disburses"
}

func (md *MaterialDisburse) BeforeCreate(tx *gorm.DB) (err error) {
	if md.ID == uuidv7.NilUUID() {
		if md.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type MaterialDisburseDetail struct {
	MaterialDisburseID uuid.UUID `json:"material_disburse_id" validate:"required" gorm:"type:uuid;not null;primaryKey"`
	MaterialID         uuid.UUID `json:"material_id" validate:"required" gorm:"type:uuid;not null;primaryKey"`
	EmployeeID         uuid.UUID `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	Qty                int       `json:"qty" validate:"required,min=1"`
	
	// Relations
	MaterialDisburse MaterialDisburse `json:"material_disburse,omitempty" gorm:"foreignKey:MaterialDisburseID"`
	Material         Material          `json:"material,omitempty" gorm:"foreignKey:MaterialID"`
	Employee         Employee          `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
}

func (mdd *MaterialDisburseDetail) TableName() string {
	return "material_disburse_details"
}

