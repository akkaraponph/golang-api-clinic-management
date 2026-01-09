package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AppointmentHandler struct {
	service ports.AppointmentService
}

func NewAppointmentHandler(service ports.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	var req dto.AppointmentCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	appointment := dto.ToAppointmentDomain(&req)
	if err := h.service.CreateAppointment(appointment); err != nil {
		return utils.NewErrorResponse(c, "Failed to create appointment", err.Error())
	}

	return utils.NewSuccessResponse(c, "Appointment created successfully", dto.ToAppointmentResponse(appointment))
}

func (h *AppointmentHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid appointment ID", nil)
	}

	appointment, err := h.service.GetAppointmentByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Appointment not found", nil)
	}

	return utils.NewSuccessResponse(c, "Appointment retrieved successfully", dto.ToAppointmentResponse(appointment))
}

func (h *AppointmentHandler) GetAll(c *fiber.Ctx) error {
	appointments, err := h.service.GetAllAppointments()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve appointments", nil)
	}

	responses := make([]*dto.AppointmentResponse, len(appointments))
	for i, a := range appointments {
		responses[i] = dto.ToAppointmentResponse(a)
	}

	return utils.NewSuccessResponse(c, "Appointments retrieved successfully", responses)
}

func (h *AppointmentHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patient_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	appointments, err := h.service.GetAppointmentsByPatientID(patientID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve appointments", nil)
	}

	responses := make([]*dto.AppointmentResponse, len(appointments))
	for i, a := range appointments {
		responses[i] = dto.ToAppointmentResponse(a)
	}

	return utils.NewSuccessResponse(c, "Appointments retrieved successfully", responses)
}

func (h *AppointmentHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid appointment ID", nil)
	}

	var req dto.AppointmentUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	appointment, err := h.service.GetAppointmentByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Appointment not found", nil)
	}

	// Update fields if provided
	if req.Subject != "" {
		appointment.Subject = req.Subject
	}
	if req.Detail != "" {
		appointment.Detail = req.Detail
	}
	if req.Date != nil {
		appointment.Date = *req.Date
	}
	if req.Status != "" {
		appointment.Status = domain.AppointmentStatus(req.Status)
	}
	if req.PatientID != nil {
		appointment.PatientID = *req.PatientID
	}
	if req.EmployeeID != nil {
		appointment.EmployeeID = *req.EmployeeID
	}

	if err := h.service.UpdateAppointment(appointment); err != nil {
		return utils.NewErrorResponse(c, "Failed to update appointment", err.Error())
	}

	return utils.NewSuccessResponse(c, "Appointment updated successfully", dto.ToAppointmentResponse(appointment))
}

func (h *AppointmentHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid appointment ID", nil)
	}

	if err := h.service.DeleteAppointment(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete appointment", nil)
	}

	return utils.NewSuccessResponse(c, "Appointment deleted successfully", nil)
}
