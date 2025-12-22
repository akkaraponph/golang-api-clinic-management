package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type CourseRepository interface {
	Create(course *domain.Course) error
	GetByID(id uuid.UUID) (*domain.Course, error)
	GetAll() ([]*domain.Course, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Course, error)
	Update(course *domain.Course) error
	Delete(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
	Restore(id uuid.UUID) error
	GetTrash() ([]*domain.Course, error)
}

type CourseService interface {
	CreateCourse(course *domain.Course) error
	GetCourseByID(id uuid.UUID) (*domain.Course, error)
	GetAllCourses() ([]*domain.Course, error)
	GetCoursesByPatientID(patientID uuid.UUID) ([]*domain.Course, error)
	UpdateCourse(course *domain.Course) error
	DeleteCourse(id uuid.UUID) error
	RestoreCourse(id uuid.UUID) error
	GetTrashCourses() ([]*domain.Course, error)
}

