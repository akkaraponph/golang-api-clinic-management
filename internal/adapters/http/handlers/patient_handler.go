package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/uuidv7"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PatientHandler struct {
	service ports.PatientService
}

func NewPatientHandler(service ports.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

func (h *PatientHandler) Create(c *fiber.Ctx) error {
	var req dto.PatientCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	patient := dto.ToPatientDomain(&req)
	if err := h.service.CreatePatient(patient); err != nil {
		return utils.NewErrorResponse(c, "Failed to create patient", err.Error())
	}

	return utils.NewSuccessResponse(c, "Patient created successfully", dto.ToPatientResponse(patient))
}

func (h *PatientHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	patient, err := h.service.GetPatientByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Patient not found", nil)
	}

	return utils.NewSuccessResponse(c, "Patient retrieved successfully", dto.ToPatientResponse(patient))
}

func (h *PatientHandler) GetAll(c *fiber.Ctx) error {
	patients, err := h.service.GetAllPatients()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve patients", nil)
	}

	responses := make([]*dto.PatientResponse, len(patients))
	for i, p := range patients {
		responses[i] = dto.ToPatientResponse(p)
	}

	return utils.NewSuccessResponse(c, "Patients retrieved successfully", responses)
}

func (h *PatientHandler) Search(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return utils.NewErrorResponse(c, "Query parameter is required", nil)
	}

	patients, err := h.service.SearchPatients(query)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to search patients", nil)
	}

	responses := make([]*dto.PatientResponse, len(patients))
	for i, p := range patients {
		responses[i] = dto.ToPatientResponse(p)
	}

	return utils.NewSuccessResponse(c, "Search completed successfully", responses)
}

func (h *PatientHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	var req dto.PatientUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	patient, err := h.service.GetPatientByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Patient not found", nil)
	}

	// Update fields
	if req.Name != "" {
		patient.Name = req.Name
	}
	if req.Surname != "" {
		patient.Surname = req.Surname
	}
	// Update other fields similarly...

	if err := h.service.UpdatePatient(patient); err != nil {
		return utils.NewErrorResponse(c, "Failed to update patient", nil)
	}

	return utils.NewSuccessResponse(c, "Patient updated successfully", dto.ToPatientResponse(patient))
}

func (h *PatientHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	if err := h.service.DeletePatient(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete patient", nil)
	}

	return utils.NewSuccessResponse(c, "Patient deleted successfully", nil)
}

