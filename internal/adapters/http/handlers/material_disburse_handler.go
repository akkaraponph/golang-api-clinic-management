package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MaterialDisburseHandler struct {
	service ports.MaterialDisburseService
}

func NewMaterialDisburseHandler(service ports.MaterialDisburseService) *MaterialDisburseHandler {
	return &MaterialDisburseHandler{service: service}
}

func (h *MaterialDisburseHandler) Create(c *fiber.Ctx) error {
	var req dto.MaterialDisburseCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	// Create disburse first
	disburse := &domain.MaterialDisburse{
		Date:       req.Date,
		EmployeeID: req.EmployeeID,
	}

	if err := h.service.CreateMaterialDisburse(disburse); err != nil {
		return utils.NewErrorResponse(c, "Failed to create material disburse", err.Error())
	}

	// Set disburse ID in details
	for _, d := range req.Details {
		disburse.MaterialDisburseDetails = append(disburse.MaterialDisburseDetails, domain.MaterialDisburseDetail{
			MaterialDisburseID: disburse.ID,
			MaterialID:         d.MaterialID,
			EmployeeID:         d.EmployeeID,
			Qty:                d.Qty,
		})
	}

	// Update disburse with details
	if err := h.service.UpdateMaterialDisburse(disburse); err != nil {
		return utils.NewErrorResponse(c, "Failed to create material disburse details", err.Error())
	}

	// Fetch full disburse with details
	fullDisburse, err := h.service.GetMaterialDisburseByID(disburse.ID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve material disburse", err.Error())
	}

	return utils.NewSuccessResponse(c, "Material disburse created successfully", dto.ToMaterialDisburseResponse(fullDisburse))
}

func (h *MaterialDisburseHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material disburse ID", nil)
	}

	disburse, err := h.service.GetMaterialDisburseByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Material disburse not found", nil)
	}

	return utils.NewSuccessResponse(c, "Material disburse retrieved successfully", dto.ToMaterialDisburseResponse(disburse))
}

func (h *MaterialDisburseHandler) GetAll(c *fiber.Ctx) error {
	disburses, err := h.service.GetAllMaterialDisburses()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve material disburses", nil)
	}

	responses := make([]*dto.MaterialDisburseResponse, len(disburses))
	for i, d := range disburses {
		responses[i] = dto.ToMaterialDisburseResponse(d)
	}

	return utils.NewSuccessResponse(c, "Material disburses retrieved successfully", responses)
}

func (h *MaterialDisburseHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material disburse ID", nil)
	}

	var req dto.MaterialDisburseCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	disburse, err := h.service.GetMaterialDisburseByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Material disburse not found", nil)
	}

	disburse.Date = req.Date
	disburse.EmployeeID = req.EmployeeID

	// Update details
	disburse.MaterialDisburseDetails = make([]domain.MaterialDisburseDetail, len(req.Details))
	for i, d := range req.Details {
		disburse.MaterialDisburseDetails[i] = domain.MaterialDisburseDetail{
			MaterialDisburseID: disburse.ID,
			MaterialID:         d.MaterialID,
			EmployeeID:         d.EmployeeID,
			Qty:                d.Qty,
		}
	}

	if err := h.service.UpdateMaterialDisburse(disburse); err != nil {
		return utils.NewErrorResponse(c, "Failed to update material disburse", err.Error())
	}

	return utils.NewSuccessResponse(c, "Material disburse updated successfully", dto.ToMaterialDisburseResponse(disburse))
}

func (h *MaterialDisburseHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material disburse ID", nil)
	}

	if err := h.service.DeleteMaterialDisburse(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete material disburse", nil)
	}

	return utils.NewSuccessResponse(c, "Material disburse deleted successfully", nil)
}
