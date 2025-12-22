package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) ports.EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(employee *domain.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetByID(id uuid.UUID) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.Where("id = ? AND is_remove = ?", id, false).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) GetByUsername(username string) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.Where("username = ? AND is_remove = ?", username, false).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) GetAll() ([]*domain.Employee, error) {
	var employees []*domain.Employee
	err := r.db.Where("is_remove = ?", false).Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) Update(employee *domain.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Employee{}, id).Error
}

func (r *employeeRepository) SoftDelete(id uuid.UUID) error {
	return r.db.Model(&domain.Employee{}).Where("id = ?", id).Update("is_remove", true).Error
}

func (r *employeeRepository) Restore(id uuid.UUID) error {
	return r.db.Model(&domain.Employee{}).Where("id = ?", id).Update("is_remove", false).Error
}

