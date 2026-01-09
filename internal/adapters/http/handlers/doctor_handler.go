package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DoctorHandler struct {
	service ports.DoctorService
}

func NewDoctorHandler(service ports.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

func (h *DoctorHandler) Create(c *fiber.Ctx) error {
	var req dto.DoctorCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	doctor := dto.ToDoctorDomain(&req)
	if err := h.service.CreateDoctor(doctor); err != nil {
		return utils.NewErrorResponse(c, "Failed to create doctor", err.Error())
	}

	return utils.NewSuccessResponse(c, "Doctor created successfully", dto.ToDoctorResponse(doctor))
}

func (h *DoctorHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid doctor ID", nil)
	}

	doctor, err := h.service.GetDoctorByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Doctor not found", nil)
	}

	return utils.NewSuccessResponse(c, "Doctor retrieved successfully", dto.ToDoctorResponse(doctor))
}

func (h *DoctorHandler) GetByEmployeeID(c *fiber.Ctx) error {
	employeeID, err := uuid.Parse(c.Params("employee_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid employee ID", nil)
	}

	doctor, err := h.service.GetDoctorByEmployeeID(employeeID)
	if err != nil {
		return utils.NewErrorResponse(c, "Doctor not found", nil)
	}

	return utils.NewSuccessResponse(c, "Doctor retrieved successfully", dto.ToDoctorResponse(doctor))
}

func (h *DoctorHandler) GetAll(c *fiber.Ctx) error {
	doctors, err := h.service.GetAllDoctors()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve doctors", nil)
	}

	responses := make([]*dto.DoctorResponse, len(doctors))
	for i, d := range doctors {
		responses[i] = dto.ToDoctorResponse(d)
	}

	return utils.NewSuccessResponse(c, "Doctors retrieved successfully", responses)
}

func (h *DoctorHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid doctor ID", nil)
	}

	var req dto.DoctorUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	doctor, err := h.service.GetDoctorByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Doctor not found", nil)
	}

	// Update fields if provided
	if req.EmployeeID != nil {
		doctor.EmployeeID = *req.EmployeeID
	}
	if req.Specialization != "" {
		doctor.Specialization = req.Specialization
	}
	if req.LicenseNumber != "" {
		doctor.LicenseNumber = req.LicenseNumber
	}
	if req.RegistrationDate != nil {
		doctor.RegistrationDate = req.RegistrationDate
	}
	if req.ExperienceYears != nil {
		doctor.ExperienceYears = req.ExperienceYears
	}

	if err := h.service.UpdateDoctor(doctor); err != nil {
		return utils.NewErrorResponse(c, "Failed to update doctor", err.Error())
	}

	return utils.NewSuccessResponse(c, "Doctor updated successfully", dto.ToDoctorResponse(doctor))
}

func (h *DoctorHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid doctor ID", nil)
	}

	if err := h.service.DeleteDoctor(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete doctor", nil)
	}

	return utils.NewSuccessResponse(c, "Doctor deleted successfully", nil)
}
