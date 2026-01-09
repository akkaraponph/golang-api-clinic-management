package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MaterialOrderHandler struct {
	service ports.MaterialOrderService
}

func NewMaterialOrderHandler(service ports.MaterialOrderService) *MaterialOrderHandler {
	return &MaterialOrderHandler{service: service}
}

func (h *MaterialOrderHandler) Create(c *fiber.Ctx) error {
	var req dto.MaterialOrderCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	order := dto.ToMaterialOrderDomain(&req)
	if err := h.service.CreateMaterialOrder(order); err != nil {
		return utils.NewErrorResponse(c, "Failed to create material order", err.Error())
	}

	return utils.NewSuccessResponse(c, "Material order created successfully", dto.ToMaterialOrderResponse(order))
}

func (h *MaterialOrderHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material order ID", nil)
	}

	order, err := h.service.GetMaterialOrderByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Material order not found", nil)
	}

	return utils.NewSuccessResponse(c, "Material order retrieved successfully", dto.ToMaterialOrderResponse(order))
}

func (h *MaterialOrderHandler) GetAll(c *fiber.Ctx) error {
	orders, err := h.service.GetAllMaterialOrders()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve material orders", nil)
	}

	responses := make([]*dto.MaterialOrderResponse, len(orders))
	for i, o := range orders {
		responses[i] = dto.ToMaterialOrderResponse(o)
	}

	return utils.NewSuccessResponse(c, "Material orders retrieved successfully", responses)
}

func (h *MaterialOrderHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material order ID", nil)
	}

	var req dto.MaterialOrderUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	order, err := h.service.GetMaterialOrderByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Material order not found", nil)
	}

	// Update fields if provided
	if req.TotalPrice != nil {
		order.TotalPrice = *req.TotalPrice
	}
	if req.PurchaseDate != nil {
		order.PurchaseDate = req.PurchaseDate
	}
	if req.ReceiveDate != nil {
		order.ReceiveDate = req.ReceiveDate
	}
	if req.Status != "" {
		order.Status = domain.MaterialOrderStatus(req.Status)
	}
	if req.EmployeeID != nil {
		order.EmployeeID = *req.EmployeeID
	}
	if req.AgentID != nil {
		order.AgentID = *req.AgentID
	}

	if err := h.service.UpdateMaterialOrder(order); err != nil {
		return utils.NewErrorResponse(c, "Failed to update material order", err.Error())
	}

	return utils.NewSuccessResponse(c, "Material order updated successfully", dto.ToMaterialOrderResponse(order))
}

func (h *MaterialOrderHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid material order ID", nil)
	}

	if err := h.service.DeleteMaterialOrder(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete material order", nil)
	}

	return utils.NewSuccessResponse(c, "Material order deleted successfully", nil)
}
