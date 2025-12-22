package domain

import (
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Geography struct {
	BaseModel
	Name string `json:"name" validate:"required,max=255" gorm:"size:255"`
	
	// Relations
	Provinces []Province `json:"provinces,omitempty" gorm:"foreignKey:GeographyID"`
}

func (g *Geography) TableName() string {
	return "geographies"
}

func (g *Geography) BeforeCreate(tx *gorm.DB) (err error) {
	if g.ID == uuidv7.NilUUID() {
		if g.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type Province struct {
	BaseModel
	Code        string     `json:"code" validate:"required" gorm:"size:10;uniqueIndex"`
	NameTh      string     `json:"name_th" validate:"required,max=255" gorm:"size:255"`
	NameEn      string     `json:"name_en" validate:"max=255" gorm:"size:255"`
	GeographyID *uuid.UUID `json:"geography_id,omitempty" gorm:"type:uuid;index"`
	
	// Relations
	Geography  *Geography  `json:"geography,omitempty" gorm:"foreignKey:GeographyID"`
	Districts  []District  `json:"districts,omitempty" gorm:"foreignKey:ProvinceID"`
}

func (p *Province) TableName() string {
	return "provinces"
}

func (p *Province) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuidv7.NilUUID() {
		if p.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type District struct {
	BaseModel
	Code      string    `json:"code" validate:"required" gorm:"size:10"`
	NameTh    string    `json:"name_th" validate:"required,max=255" gorm:"size:255"`
	NameEn    string    `json:"name_en" validate:"max=255" gorm:"size:255"`
	ProvinceID uuid.UUID `json:"province_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	Province     Province      `json:"province,omitempty" gorm:"foreignKey:ProvinceID"`
	Subdistricts []Subdistrict `json:"subdistricts,omitempty" gorm:"foreignKey:DistrictID"`
}

func (d *District) TableName() string {
	return "districts"
}

func (d *District) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuidv7.NilUUID() {
		if d.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

type Subdistrict struct {
	BaseModel
	ZipCode    string    `json:"zip_code" validate:"max=10" gorm:"size:10"`
	NameTh     string    `json:"name_th" validate:"required,max=255" gorm:"size:255"`
	NameEn     string    `json:"name_en" validate:"max=255" gorm:"size:255"`
	DistrictID uuid.UUID `json:"district_id" validate:"required" gorm:"type:uuid;not null;index"`
	
	// Relations
	District District `json:"district,omitempty" gorm:"foreignKey:DistrictID"`
}

func (s *Subdistrict) TableName() string {
	return "subdistricts"
}

func (s *Subdistrict) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuidv7.NilUUID() {
		if s.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
			return err
		}
	}
	return nil
}

