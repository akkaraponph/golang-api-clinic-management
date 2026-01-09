package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type InsuranceHandler struct {
	service ports.InsuranceService
}

func NewInsuranceHandler(service ports.InsuranceService) *InsuranceHandler {
	return &InsuranceHandler{service: service}
}

func (h *InsuranceHandler) Create(c *fiber.Ctx) error {
	var req dto.InsuranceCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	insurance := dto.ToInsuranceDomain(&req)
	if err := h.service.CreateInsurance(insurance); err != nil {
		return utils.NewErrorResponse(c, "Failed to create insurance", err.Error())
	}

	return utils.NewSuccessResponse(c, "Insurance created successfully", dto.ToInsuranceResponse(insurance))
}

func (h *InsuranceHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid insurance ID", nil)
	}

	insurance, err := h.service.GetInsuranceByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Insurance not found", nil)
	}

	return utils.NewSuccessResponse(c, "Insurance retrieved successfully", dto.ToInsuranceResponse(insurance))
}

func (h *InsuranceHandler) GetAll(c *fiber.Ctx) error {
	insurances, err := h.service.GetAllInsurances()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve insurances", nil)
	}

	responses := make([]*dto.InsuranceResponse, len(insurances))
	for i, ins := range insurances {
		responses[i] = dto.ToInsuranceResponse(ins)
	}

	return utils.NewSuccessResponse(c, "Insurances retrieved successfully", responses)
}

func (h *InsuranceHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patient_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	insurances, err := h.service.GetInsurancesByPatientID(patientID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve insurances", nil)
	}

	responses := make([]*dto.InsuranceResponse, len(insurances))
	for i, ins := range insurances {
		responses[i] = dto.ToInsuranceResponse(ins)
	}

	return utils.NewSuccessResponse(c, "Insurances retrieved successfully", responses)
}

func (h *InsuranceHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid insurance ID", nil)
	}

	var req dto.InsuranceUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	insurance, err := h.service.GetInsuranceByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Insurance not found", nil)
	}

	if req.PatientID != nil {
		insurance.PatientID = *req.PatientID
	}
	if req.Provider != "" {
		insurance.Provider = req.Provider
	}
	if req.PolicyNumber != "" {
		insurance.PolicyNumber = req.PolicyNumber
	}
	if req.ValidUntil != nil {
		insurance.ValidUntil = req.ValidUntil
	}

	if err := h.service.UpdateInsurance(insurance); err != nil {
		return utils.NewErrorResponse(c, "Failed to update insurance", err.Error())
	}

	return utils.NewSuccessResponse(c, "Insurance updated successfully", dto.ToInsuranceResponse(insurance))
}

func (h *InsuranceHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid insurance ID", nil)
	}

	if err := h.service.DeleteInsurance(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete insurance", nil)
	}

	return utils.NewSuccessResponse(c, "Insurance deleted successfully", nil)
}
