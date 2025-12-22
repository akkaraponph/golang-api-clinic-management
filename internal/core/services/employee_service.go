package services

import (
	"errors"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/argon2id"
	"github.com/google/uuid"
)

type employeeService struct {
	repo ports.EmployeeRepository
}

func NewEmployeeService(repo ports.EmployeeRepository) ports.EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(employee *domain.Employee) error {
	// Hash password before saving
	if employee.Password != "" {
		employee.Password = argon2id.NewPassword(employee.Password)
	}
	return s.repo.Create(employee)
}

func (s *employeeService) GetEmployeeByID(id uuid.UUID) (*domain.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) GetEmployeeByUsername(username string) (*domain.Employee, error) {
	return s.repo.GetByUsername(username)
}

func (s *employeeService) GetAllEmployees() ([]*domain.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) UpdateEmployee(employee *domain.Employee) error {
	// If password is being updated, hash it
	if employee.Password != "" && len(employee.Password) < 100 { // Not already hashed
		employee.Password = argon2id.NewPassword(employee.Password)
	}
	return s.repo.Update(employee)
}

func (s *employeeService) DeleteEmployee(id uuid.UUID) error {
	return s.repo.SoftDelete(id)
}

func (s *employeeService) Authenticate(username, password string) (*domain.Employee, error) {
	employee, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	
	if !argon2id.ComparePassword(employee.Password, password) {
		return nil, errors.New("invalid credentials")
	}
	
	return employee, nil
}

