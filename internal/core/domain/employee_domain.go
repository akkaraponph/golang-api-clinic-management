package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"gorm.io/gorm"
)

type PermissionEnum string

const (
	PermissionOwner    PermissionEnum = "owner"
	PermissionEmployee PermissionEnum = "employee"
)

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

type Employee struct {
	BaseModel
	Name        string        `json:"name" validate:"required,max=100" gorm:"size:100"`
	Surname     string        `json:"surname" validate:"required,max=100" gorm:"size:100"`
	Email       string        `json:"email" validate:"required,max=255,email" gorm:"size:255"`
	Username    string        `json:"username" validate:"required,max=100" gorm:"size:100;uniqueIndex"`
	Password    string        `json:"-" gorm:"size:255"` // Hidden from JSON
	Sex         Sex           `json:"sex" validate:"required" gorm:"size:10"`
	Phone       string        `json:"phone" validate:"max=20" gorm:"size:20"`
	Permission  PermissionEnum `json:"permission" validate:"required" gorm:"size:50"`
	Address     string        `json:"address" validate:"max=255" gorm:"size:255"`
	Subdistrict string        `json:"subdistrict" validate:"max=100" gorm:"size:100"`
	District    string        `json:"district" validate:"max=100" gorm:"size:100"`
	Province    string        `json:"province" validate:"max=100" gorm:"size:100"`
	Zipcode     string        `json:"zipcode" validate:"max=10" gorm:"size:10"`
	IsRemove    bool          `json:"is_remove" gorm:"default:false"`
}

func (e *Employee) TableName() string {
	return "employees"
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == uuidv7.NilUUID() {
		if e.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

