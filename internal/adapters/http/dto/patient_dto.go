package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type PatientCreateRequest struct {
	Name        string     `json:"name" validate:"required,max=100"`
	Surname     string     `json:"surname" validate:"required,max=100"`
	IDCard      string     `json:"id_card" validate:"max=20"`
	Sex         string     `json:"sex" validate:"required"`
	BirthDate   *time.Time `json:"birth_date"`
	BloodType   string     `json:"blood_type" validate:"max=10"`
	Disease     string     `json:"disease"`
	Allergic    string     `json:"allergic"`
	Phone       string     `json:"phone" validate:"max=20"`
	Address     string     `json:"address" validate:"max=255"`
	Subdistrict string     `json:"subdistrict" validate:"max=100"`
	District    string     `json:"district" validate:"max=100"`
	Province    string     `json:"province" validate:"max=100"`
	Zipcode     string     `json:"zipcode" validate:"max=10"`
}

type PatientUpdateRequest struct {
	Name        string     `json:"name" validate:"max=100"`
	Surname     string     `json:"surname" validate:"max=100"`
	IDCard      string     `json:"id_card" validate:"max=20"`
	Sex         string     `json:"sex"`
	BirthDate   *time.Time `json:"birth_date"`
	BloodType   string     `json:"blood_type" validate:"max=10"`
	Disease     string     `json:"disease"`
	Allergic    string     `json:"allergic"`
	Phone       string     `json:"phone" validate:"max=20"`
	Address     string     `json:"address" validate:"max=255"`
	Subdistrict string     `json:"subdistrict" validate:"max=100"`
	District    string     `json:"district" validate:"max=100"`
	Province    string     `json:"province" validate:"max=100"`
	Zipcode     string     `json:"zipcode" validate:"max=10"`
}

type PatientResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Surname     string     `json:"surname"`
	IDCard      string     `json:"id_card"`
	Sex         string     `json:"sex"`
	BirthDate   *time.Time `json:"birth_date"`
	BloodType   string     `json:"blood_type"`
	Disease     string     `json:"disease"`
	Allergic    string     `json:"allergic"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
	Subdistrict string     `json:"subdistrict"`
	District    string     `json:"district"`
	Province    string     `json:"province"`
	Zipcode     string     `json:"zipcode"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func ToPatientResponse(p *domain.Patient) *PatientResponse {
	return &PatientResponse{
		ID:          p.ID,
		Name:        p.Name,
		Surname:     p.Surname,
		IDCard:      p.IDCard,
		Sex:         string(p.Sex),
		BirthDate:   p.BirthDate,
		BloodType:   p.BloodType,
		Disease:     p.Disease,
		Allergic:    p.Allergic,
		Phone:       p.Phone,
		Address:     p.Address,
		Subdistrict: p.Subdistrict,
		District:    p.District,
		Province:    p.Province,
		Zipcode:     p.Zipcode,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ToPatientDomain(req *PatientCreateRequest) *domain.Patient {
	return &domain.Patient{
		Name:        req.Name,
		Surname:     req.Surname,
		IDCard:      req.IDCard,
		Sex:         domain.Sex(req.Sex),
		BirthDate:   req.BirthDate,
		BloodType:   req.BloodType,
		Disease:     req.Disease,
		Allergic:    req.Allergic,
		Phone:       req.Phone,
		Address:     req.Address,
		Subdistrict: req.Subdistrict,
		District:    req.District,
		Province:    req.Province,
		Zipcode:     req.Zipcode,
		IsRemove:    false,
	}
}

