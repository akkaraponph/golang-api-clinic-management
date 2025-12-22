package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type EmployeeRepository interface {
	Create(employee *domain.Employee) error
	GetByID(id uuid.UUID) (*domain.Employee, error)
	GetByUsername(username string) (*domain.Employee, error)
	GetAll() ([]*domain.Employee, error)
	Update(employee *domain.Employee) error
	Delete(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
	Restore(id uuid.UUID) error
}

type EmployeeService interface {
	CreateEmployee(employee *domain.Employee) error
	GetEmployeeByID(id uuid.UUID) (*domain.Employee, error)
	GetEmployeeByUsername(username string) (*domain.Employee, error)
	GetAllEmployees() ([]*domain.Employee, error)
	UpdateEmployee(employee *domain.Employee) error
	DeleteEmployee(id uuid.UUID) error
	Authenticate(username, password string) (*domain.Employee, error)
}

