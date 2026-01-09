package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MaterialHandler struct {
	service ports.MaterialService
}

func NewMaterialHandler(service ports.MaterialService) *MaterialHandler {
	return &MaterialHandler{service: service}
}

func (h *MaterialHandler) Create(c *fiber.Ctx) error {
	var req dto.MaterialCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	material := dto.ToMaterialDomain(&req)
	if err := h.service.CreateMaterial(material); err != nil {
		return utils.NewErrorResponse(c, "Failed to create material", err.Error())
	}

	return utils.NewSuccessResponse(c, "Material created successfully", dto.ToMaterialResponse(material))
}

func (h *MaterialHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material ID", nil)
	}

	material, err := h.service.GetMaterialByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Material not found", nil)
	}

	return utils.NewSuccessResponse(c, "Material retrieved successfully", dto.ToMaterialResponse(material))
}

func (h *MaterialHandler) GetAll(c *fiber.Ctx) error {
	materials, err := h.service.GetAllMaterials()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve materials", nil)
	}

	responses := make([]*dto.MaterialResponse, len(materials))
	for i, m := range materials {
		responses[i] = dto.ToMaterialResponse(m)
	}

	return utils.NewSuccessResponse(c, "Materials retrieved successfully", responses)
}

func (h *MaterialHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material ID", nil)
	}

	var req dto.MaterialUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	material, err := h.service.GetMaterialByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Material not found", nil)
	}

	// Update fields if provided
	if req.Name != "" {
		material.Name = req.Name
	}
	if req.Detail != "" {
		material.Detail = req.Detail
	}
	if req.Unit != "" {
		material.Unit = req.Unit
	}
	if req.PurchasePrice != nil {
		material.PurchasePrice = *req.PurchasePrice
	}
	if req.Qty != nil {
		material.Qty = *req.Qty
	}
	if req.AgentID != nil {
		material.AgentID = *req.AgentID
	}

	if err := h.service.UpdateMaterial(material); err != nil {
		return utils.NewErrorResponse(c, "Failed to update material", err.Error())
	}

	return utils.NewSuccessResponse(c, "Material updated successfully", dto.ToMaterialResponse(material))
}

func (h *MaterialHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material ID", nil)
	}

	if err := h.service.DeleteMaterial(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete material", nil)
	}

	return utils.NewSuccessResponse(c, "Material deleted successfully", nil)
}
