package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MedicineTypeHandler struct {
	service ports.MedicineTypeService
}

func NewMedicineTypeHandler(service ports.MedicineTypeService) *MedicineTypeHandler {
	return &MedicineTypeHandler{service: service}
}

func (h *MedicineTypeHandler) Create(c *fiber.Ctx) error {
	var req dto.MedicineTypeCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	medicineType := dto.ToMedicineTypeDomain(&req)
	if err := h.service.CreateMedicineType(medicineType); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine type", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine type created successfully", dto.ToMedicineTypeResponse(medicineType))
}

func (h *MedicineTypeHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine type ID", nil)
	}

	medicineType, err := h.service.GetMedicineTypeByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine type not found", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine type retrieved successfully", dto.ToMedicineTypeResponse(medicineType))
}

func (h *MedicineTypeHandler) GetAll(c *fiber.Ctx) error {
	medicineTypes, err := h.service.GetAllMedicineTypes()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine types", nil)
	}

	responses := make([]*dto.MedicineTypeResponse, len(medicineTypes))
	for i, mt := range medicineTypes {
		responses[i] = dto.ToMedicineTypeResponse(mt)
	}

	return utils.NewSuccessResponse(c, "Medicine types retrieved successfully", responses)
}

func (h *MedicineTypeHandler) GetByAgentID(c *fiber.Ctx) error {
	agentID, err := uuid.Parse(c.Params("agent_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid agent ID", nil)
	}

	medicineTypes, err := h.service.GetMedicineTypesByAgentID(agentID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine types", nil)
	}

	responses := make([]*dto.MedicineTypeResponse, len(medicineTypes))
	for i, mt := range medicineTypes {
		responses[i] = dto.ToMedicineTypeResponse(mt)
	}

	return utils.NewSuccessResponse(c, "Medicine types retrieved successfully", responses)
}

func (h *MedicineTypeHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine type ID", nil)
	}

	var req dto.MedicineTypeUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	medicineType, err := h.service.GetMedicineTypeByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine type not found", nil)
	}

	if req.Name != "" {
		medicineType.Name = req.Name
	}
	if req.Detail != "" {
		medicineType.Detail = req.Detail
	}
	if req.AgentID != nil {
		medicineType.AgentID = *req.AgentID
	}

	if err := h.service.UpdateMedicineType(medicineType); err != nil {
		return utils.NewErrorResponse(c, "Failed to update medicine type", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine type updated successfully", dto.ToMedicineTypeResponse(medicineType))
}

func (h *MedicineTypeHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine type ID", nil)
	}

	if err := h.service.DeleteMedicineType(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete medicine type", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine type deleted successfully", nil)
}
