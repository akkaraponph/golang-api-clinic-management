package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type courseService struct {
	repo ports.CourseRepository
}

func NewCourseService(repo ports.CourseRepository) ports.CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(course *domain.Course) error {
	return s.repo.Create(course)
}

func (s *courseService) GetCourseByID(id uuid.UUID) (*domain.Course, error) {
	return s.repo.GetByID(id)
}

func (s *courseService) GetAllCourses() ([]*domain.Course, error) {
	return s.repo.GetAll()
}

func (s *courseService) GetCoursesByPatientID(patientID uuid.UUID) ([]*domain.Course, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *courseService) UpdateCourse(course *domain.Course) error {
	return s.repo.Update(course)
}

func (s *courseService) DeleteCourse(id uuid.UUID) error {
	return s.repo.SoftDelete(id)
}

func (s *courseService) RestoreCourse(id uuid.UUID) error {
	return s.repo.Restore(id)
}

func (s *courseService) GetTrashCourses() ([]*domain.Course, error) {
	return s.repo.GetTrash()
}

