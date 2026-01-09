package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MedicineOrderHandler struct {
	service ports.MedicineOrderService
}

func NewMedicineOrderHandler(service ports.MedicineOrderService) *MedicineOrderHandler {
	return &MedicineOrderHandler{service: service}
}

func (h *MedicineOrderHandler) Create(c *fiber.Ctx) error {
	var req dto.MedicineOrderCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	// Create order first
	order := &domain.MedicineOrder{
		Date:       req.Date,
		TotalPrice: req.TotalPrice,
		ReceiveDate: req.ReceiveDate,
		Status:     domain.MedicineOrderStatus(req.Status),
		EmployeeID: req.EmployeeID,
		AgentID:    req.AgentID,
	}

	if err := h.service.CreateMedicineOrder(order); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine order", err.Error())
	}

	// Set order ID in details and add details
	for _, d := range req.Details {
		order.MedicineOrderDetails = append(order.MedicineOrderDetails, domain.MedicineOrderDetail{
			MedicineOrderID: order.ID,
			MedicineID:      d.MedicineID,
			Price:           d.Price,
			Qty:             d.Qty,
			Remain:          d.Remain,
		})
	}

	// Update order with details
	if err := h.service.UpdateMedicineOrder(order); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine order details", err.Error())
	}

	// Fetch full order with details
	fullOrder, err := h.service.GetMedicineOrderByID(order.ID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine order", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine order created successfully", dto.ToMedicineOrderResponse(fullOrder))
}

func (h *MedicineOrderHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine order ID", nil)
	}

	order, err := h.service.GetMedicineOrderByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine order not found", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine order retrieved successfully", dto.ToMedicineOrderResponse(order))
}

func (h *MedicineOrderHandler) GetAll(c *fiber.Ctx) error {
	orders, err := h.service.GetAllMedicineOrders()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine orders", nil)
	}

	responses := make([]*dto.MedicineOrderResponse, len(orders))
	for i, o := range orders {
		responses[i] = dto.ToMedicineOrderResponse(o)
	}

	return utils.NewSuccessResponse(c, "Medicine orders retrieved successfully", responses)
}

func (h *MedicineOrderHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine order ID", nil)
	}

	var req dto.MedicineOrderCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	order, err := h.service.GetMedicineOrderByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine order not found", nil)
	}

	// Update basic fields
	order.Date = req.Date
	order.TotalPrice = req.TotalPrice
	order.ReceiveDate = req.ReceiveDate
	order.Status = domain.MedicineOrderStatus(req.Status)
	order.EmployeeID = req.EmployeeID
	order.AgentID = req.AgentID

	// Update details
	order.MedicineOrderDetails = make([]domain.MedicineOrderDetail, len(req.Details))
	for i, d := range req.Details {
		order.MedicineOrderDetails[i] = domain.MedicineOrderDetail{
			MedicineOrderID: order.ID,
			MedicineID:      d.MedicineID,
			Price:           d.Price,
			Qty:             d.Qty,
			Remain:          d.Remain,
		}
	}

	if err := h.service.UpdateMedicineOrder(order); err != nil {
		return utils.NewErrorResponse(c, "Failed to update medicine order", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine order updated successfully", dto.ToMedicineOrderResponse(order))
}

func (h *MedicineOrderHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine order ID", nil)
	}

	if err := h.service.DeleteMedicineOrder(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete medicine order", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine order deleted successfully", nil)
}
