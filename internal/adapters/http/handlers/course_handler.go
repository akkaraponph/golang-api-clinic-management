package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CourseHandler struct {
	service ports.CourseService
}

func NewCourseHandler(service ports.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

func (h *CourseHandler) Create(c *fiber.Ctx) error {
	var req dto.CourseCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	course := dto.ToCourseDomain(&req)
	if err := h.service.CreateCourse(course); err != nil {
		return utils.NewErrorResponse(c, "Failed to create course", err.Error())
	}

	return utils.NewSuccessResponse(c, "Course created successfully", dto.ToCourseResponse(course))
}

func (h *CourseHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid course ID", nil)
	}

	course, err := h.service.GetCourseByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Course not found", nil)
	}

	return utils.NewSuccessResponse(c, "Course retrieved successfully", dto.ToCourseResponse(course))
}

func (h *CourseHandler) GetAll(c *fiber.Ctx) error {
	courses, err := h.service.GetAllCourses()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve courses", nil)
	}

	responses := make([]*dto.CourseResponse, len(courses))
	for i, course := range courses {
		responses[i] = dto.ToCourseResponse(course)
	}

	return utils.NewSuccessResponse(c, "Courses retrieved successfully", responses)
}

func (h *CourseHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patient_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	courses, err := h.service.GetCoursesByPatientID(patientID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve courses", nil)
	}

	responses := make([]*dto.CourseResponse, len(courses))
	for i, course := range courses {
		responses[i] = dto.ToCourseResponse(course)
	}

	return utils.NewSuccessResponse(c, "Courses retrieved successfully", responses)
}

func (h *CourseHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid course ID", nil)
	}

	var req dto.CourseUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	course, err := h.service.GetCourseByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Course not found", nil)
	}

	// Update fields if provided
	if req.Date != nil {
		course.Date = *req.Date
	}
	if req.Weight != nil {
		course.Weight = req.Weight
	}
	if req.Height != nil {
		course.Height = req.Height
	}
	if req.Systolic != nil {
		course.Systolic = req.Systolic
	}
	if req.Diastolic != nil {
		course.Diastolic = req.Diastolic
	}
	if req.HeartRate != nil {
		course.HeartRate = req.HeartRate
	}
	if req.Diagnose != "" {
		course.Diagnose = req.Diagnose
	}
	if req.TotalPrice != nil {
		course.TotalPrice = *req.TotalPrice
	}
	if req.Status != "" {
		course.Status = domain.CourseStatus(req.Status)
	}
	if req.PatientID != nil {
		course.PatientID = *req.PatientID
	}
	if req.EmployeeID != nil {
		course.EmployeeID = *req.EmployeeID
	}

	if err := h.service.UpdateCourse(course); err != nil {
		return utils.NewErrorResponse(c, "Failed to update course", err.Error())
	}

	return utils.NewSuccessResponse(c, "Course updated successfully", dto.ToCourseResponse(course))
}

func (h *CourseHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid course ID", nil)
	}

	if err := h.service.DeleteCourse(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete course", nil)
	}

	return utils.NewSuccessResponse(c, "Course deleted successfully", nil)
}

func (h *CourseHandler) Restore(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid course ID", nil)
	}

	if err := h.service.RestoreCourse(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to restore course", nil)
	}

	return utils.NewSuccessResponse(c, "Course restored successfully", nil)
}

func (h *CourseHandler) GetTrash(c *fiber.Ctx) error {
	courses, err := h.service.GetTrashCourses()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve trash courses", nil)
	}

	responses := make([]*dto.CourseResponse, len(courses))
	for i, course := range courses {
		responses[i] = dto.ToCourseResponse(course)
	}

	return utils.NewSuccessResponse(c, "Trash courses retrieved successfully", responses)
}
