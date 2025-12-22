package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicineOrderStatus string

const (
	MedicineOrderStatusPending   MedicineOrderStatus = "Pending"
	MedicineOrderStatusReceived  MedicineOrderStatus = "Received"
	MedicineOrderStatusCancelled MedicineOrderStatus = "Cancelled"
)

type MedicineOrder struct {
	BaseModel
	Date        time.Time           `json:"date" validate:"required"`
	TotalPrice  float64             `json:"total_price" gorm:"type:decimal(10,2)"`
	ReceiveDate *time.Time          `json:"receive_date"`
	Status      MedicineOrderStatus `json:"status" validate:"required" gorm:"size:100"`
	EmployeeID  uuid.UUID           `json:"employee_id" validate:"required" gorm:"type:uuid;not null;index"`
	AgentID     uuid.UUID           `json:"agent_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Employee            Employee                `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
	Agent               Agent                   `json:"agent,omitempty" gorm:"foreignKey:AgentID"`
	MedicineOrderDetails []MedicineOrderDetail `json:"medicine_order_details,omitempty" gorm:"foreignKey:MedicineOrderID"`
}

func (mo *MedicineOrder) TableName() string {
	return "medicine_orders"
}

func (mo *MedicineOrder) BeforeCreate(tx *gorm.DB) (err error) {
	if mo.ID == uuidv7.NilUUID() {
		if mo.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type MedicineOrderDetail struct {
	MedicineOrderID uuid.UUID `json:"medicine_order_id" validate:"required" gorm:"type:uuid;not null;primaryKey"`
	MedicineID      uuid.UUID `json:"medicine_id" validate:"required" gorm:"type:uuid;not null;primaryKey"`
	Price           float64   `json:"price" gorm:"type:decimal(10,2)"`
	Qty             int       `json:"qty" validate:"required,min=1"`
	Remain          int       `json:"remain" gorm:"default:0"`
	
	// Relations
	MedicineOrder MedicineOrder `json:"medicine_order,omitempty" gorm:"foreignKey:MedicineOrderID"`
	Medicine      Medicine      `json:"medicine,omitempty" gorm:"foreignKey:MedicineID"`
}

func (mod *MedicineOrderDetail) TableName() string {
	return "medicine_order_details"
}

