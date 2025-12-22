package domain

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"gorm.io/gorm"
)

type Patient struct {
	BaseModel
	Name        string     `json:"name" validate:"required,max=100" gorm:"size:100"`
	Surname     string     `json:"surname" validate:"required,max=100" gorm:"size:100"`
	IDCard      string     `json:"id_card" validate:"max=20" gorm:"size:20"`
	Sex         Sex        `json:"sex" validate:"required" gorm:"size:10"`
	BirthDate   *time.Time `json:"birth_date"`
	BloodType   string     `json:"blood_type" validate:"max=10" gorm:"size:10"`
	Disease     string     `json:"disease" gorm:"type:text"`
	Allergic    string     `json:"allergic" gorm:"type:text"`
	Phone       string     `json:"phone" validate:"max=20" gorm:"size:20"`
	Address     string     `json:"address" validate:"max=255" gorm:"size:255"`
	Subdistrict string     `json:"subdistrict" validate:"max=100" gorm:"size:100"`
	District    string     `json:"district" validate:"max=100" gorm:"size:100"`
	Province    string     `json:"province" validate:"max=100" gorm:"size:100"`
	Zipcode     string     `json:"zipcode" validate:"max=10" gorm:"size:10"`
	IsRemove    bool       `json:"is_remove" gorm:"default:false"`
}

func (p *Patient) TableName() string {
	return "patients"
}

func (p *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuidv7.NilUUID() {
		if p.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

