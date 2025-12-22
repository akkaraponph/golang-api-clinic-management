package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) ports.CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *domain.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) GetByID(id uuid.UUID) (*domain.Course, error) {
	var course domain.Course
	err := r.db.Preload("Patient").Preload("Employee").Where("id = ?", id).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) GetAll() ([]*domain.Course, error) {
	var courses []*domain.Course
	err := r.db.Preload("Patient").Preload("Employee").
		Order("date DESC").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Course, error) {
	var courses []*domain.Course
	err := r.db.Where("patient_id = ?", patientID).
		Order("date DESC").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) Update(course *domain.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Course{}, id).Error
}

func (r *courseRepository) SoftDelete(id uuid.UUID) error {
	// Note: Course table doesn't have is_remove in ER diagram, but following pattern
	return r.db.Delete(&domain.Course{}, id).Error
}

func (r *courseRepository) Restore(id uuid.UUID) error {
	// If using soft delete, restore logic would go here
	return nil
}

func (r *courseRepository) GetTrash() ([]*domain.Course, error) {
	// If using soft delete, return deleted records
	return []*domain.Course{}, nil
}

