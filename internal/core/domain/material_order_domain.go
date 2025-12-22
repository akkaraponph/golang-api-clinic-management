package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialOrderStatus string

const (
	MaterialOrderStatusPending   MaterialOrderStatus = "Pending"
	MaterialOrderStatusReceived  MaterialOrderStatus = "Received"
	MaterialOrderStatusCancelled MaterialOrderStatus = "Cancelled"
)

type MaterialOrder struct {
	BaseModel
	TotalPrice   float64            `json:"total_price" gorm:"type:decimal(10,2)"`
	PurchaseDate *time.Time         `json:"purchase_date"`
	ReceiveDate  *time.Time         `json:"receive_date"`
	Status       MaterialOrderStatus `json:"status" validate:"required" gorm:"size:100"`
	EmployeeID   uuid.UUID          `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	AgentID      uuid.UUID          `json:"agent_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Employee          Employee               `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
	Agent             Agent                  `json:"agent,omitempty" gorm:"foreignKey:AgentID"`
	MaterialOrderDetails []MaterialOrderDetail `json:"material_order_details,omitempty" gorm:"foreignKey:MaterialOrderID"`
}

func (mo *MaterialOrder) TableName() string {
	return "material_orders"
}

func (mo *MaterialOrder) BeforeCreate(tx *gorm.DB) (err error) {
	if mo.ID == uuidv7.NilUUID() {
		if mo.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type MaterialOrderDetail struct {
	BaseModel
	MaterialOrderID uuid.UUID `json:"material_order_id" validate:"required" gorm:"type:uuid;not null;index"`
	MaterialID      uuid.UUID `json:"material_id" validate:"required" gorm:"type:uuid;not null;index"`
	Qty             int       `json:"qty" validate:"required,min=1"`
	Price           float64   `json:"price" gorm:"type:decimal(10,2)"`
	
	// Relations
	MaterialOrder MaterialOrder `json:"material_order,omitempty" gorm:"foreignKey:MaterialOrderID"`
	Material      Material      `json:"material,omitempty" gorm:"foreignKey:MaterialID"`
}

func (mod *MaterialOrderDetail) TableName() string {
	return "material_order_details"
}

func (mod *MaterialOrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	if mod.ID == uuidv7.NilUUID() {
		var err error
		if mod.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

