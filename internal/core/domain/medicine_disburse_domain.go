package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicineDisburse struct {
	BaseModel
	PrescriptionID uuid.UUID `json:"prescription_id" validate:"required" gorm:"type:uuid;not null;index"`
	Date           time.Time `json:"date" validate:"required"`
	
	// Relations
	Prescription            Prescription                `json:"prescription,omitempty" gorm:"foreignKey:PrescriptionID"`
	MedicineDisburseDetails []MedicineDisburseDetail   `json:"medicine_disburse_details,omitempty" gorm:"foreignKey:MedicineDisburseID"`
}

func (md *MedicineDisburse) TableName() string {
	return "medicine_disburses"
}

func (md *MedicineDisburse) BeforeCreate(tx *gorm.DB) (err error) {
	if md.ID == uuidv7.NilUUID() {
		if md.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type MedicineDisburseDetail struct {
	MedicineDisburseID uuid.UUID  `json:"medicine_disburse_id" validate:"required" gorm:"type:uuid;not null;primaryKey"`
	MedicineStockID    uuid.UUID  `json:"medicine_stock_id" validate:"required" gorm:"type:uuid;not null;primaryKey"`
	Price              float64    `json:"price" gorm:"type:decimal(10,2)"`
	Unit               string     `json:"unit" validate:"max=50" gorm:"size:50"`
	Qty                int        `json:"qty" validate:"required,min=1"`
	Dosage             string     `json:"dosage" validate:"max=100" gorm:"size:100"`
	AdminMethod        string     `json:"admin_method" validate:"max=100" gorm:"size:100"`
	ApplyMethod        string     `json:"apply_method" validate:"max=100" gorm:"size:100"`
	TimeOfAdmin        *time.Time `json:"time_of_admin"`
	
	// Relations
	MedicineDisburse MedicineDisburse `json:"medicine_disburse,omitempty" gorm:"foreignKey:MedicineDisburseID"`
	MedicineStock    MedicineStock    `json:"medicine_stock,omitempty" gorm:"foreignKey:MedicineStockID"`
}

func (mdd *MedicineDisburseDetail) TableName() string {
	return "medicine_disburse_details"
}

