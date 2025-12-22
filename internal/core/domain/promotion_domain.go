package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"gorm.io/gorm"
)

type PromotionType string

const (
	PromotionTypePercentage  PromotionType = "percentage"
	PromotionTypeFixedAmount PromotionType = "fixed_amount"
)

type Promotion struct {
	BaseModel
	Name      string        `json:"name" validate:"required,max=255" gorm:"size:255"`
	Type      PromotionType `json:"type" validate:"required" gorm:"size:50"`
	Value     float64       `json:"value" gorm:"type:decimal(10,2)"`
	StartDate *time.Time    `json:"start_date"`
	EndDate   *time.Time    `json:"end_date"`
	Detail    string        `json:"detail" gorm:"type:text"`
	IsRemove  bool          `json:"is_remove" gorm:"default:false"`
}

func (p *Promotion) TableName() string {
	return "promotions"
}

func (p *Promotion) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuidv7.NilUUID() {
		if p.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

