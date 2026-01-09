package dto

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
	"time"
)

type EmployeeCreateRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Surname     string `json:"surname" validate:"required,max=100"`
	Email       string `json:"email" validate:"required,max=255,email"`
	Username    string `json:"username" validate:"required,max=100"`
	Password    string `json:"password" validate:"required"`
	Sex         string `json:"sex" validate:"required"`
	Phone       string `json:"phone" validate:"max=20"`
	Permission  string `json:"permission" validate:"required"`
	Address     string `json:"address" validate:"max=255"`
	Subdistrict string `json:"subdistrict" validate:"max=100"`
	District    string `json:"district" validate:"max=100"`
	Province    string `json:"province" validate:"max=100"`
	Zipcode     string `json:"zipcode" validate:"max=10"`
}

type EmployeeResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Sex         string    `json:"sex"`
	Phone       string    `json:"phone"`
	Permission  string    `json:"permission"`
	Address     string    `json:"address"`
	Subdistrict string    `json:"subdistrict"`
	District    string    `json:"district"`
	Province    string    `json:"province"`
	Zipcode     string    `json:"zipcode"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToEmployeeResponse(e *domain.Employee) *EmployeeResponse {
	return &EmployeeResponse{
		ID:          e.ID,
		Name:        e.Name,
		Surname:     e.Surname,
		Email:       e.Email,
		Username:    e.Username,
		Sex:         string(e.Sex),
		Phone:       e.Phone,
		Permission:  string(e.Permission),
		Address:     e.Address,
		Subdistrict: e.Subdistrict,
		District:    e.District,
		Province:    e.Province,
		Zipcode:     e.Zipcode,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToEmployeeDomain(req *EmployeeCreateRequest) *domain.Employee {
	return &domain.Employee{
		Name:        req.Name,
		Surname:     req.Surname,
		Email:       req.Email,
		Username:    req.Username,
		Password:    req.Password, // Will be hashed in service
		Sex:         domain.Sex(req.Sex),
		Phone:       req.Phone,
		Permission:  domain.PermissionEnum(req.Permission),
		Address:     req.Address,
		Subdistrict: req.Subdistrict,
		District:    req.District,
		Province:    req.Province,
		Zipcode:     req.Zipcode,
		IsRemove:    false,
	}
}

type EmployeeUpdateRequest struct {
	Name        string `json:"name" validate:"max=100"`
	Surname     string `json:"surname" validate:"max=100"`
	Email       string `json:"email" validate:"max=255,email"`
	Username    string `json:"username" validate:"max=100"`
	Password    string `json:"password"`
	Sex         string `json:"sex"`
	Phone       string `json:"phone" validate:"max=20"`
	Permission  string `json:"permission"`
	Address     string `json:"address" validate:"max=255"`
	Subdistrict string `json:"subdistrict" validate:"max=100"`
	District    string `json:"district" validate:"max=100"`
	Province    string `json:"province" validate:"max=100"`
	Zipcode     string `json:"zipcode" validate:"max=10"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

