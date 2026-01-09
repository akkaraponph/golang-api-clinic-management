package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MedicineStockHandler struct {
	service ports.MedicineStockService
}

func NewMedicineStockHandler(service ports.MedicineStockService) *MedicineStockHandler {
	return &MedicineStockHandler{service: service}
}

func (h *MedicineStockHandler) Create(c *fiber.Ctx) error {
	var req dto.MedicineStockCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	stock := dto.ToMedicineStockDomain(&req)
	if err := h.service.CreateStock(stock); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine stock", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine stock created successfully", dto.ToMedicineStockResponse(stock))
}

func (h *MedicineStockHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine stock ID", nil)
	}

	stock, err := h.service.GetStockByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine stock not found", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine stock retrieved successfully", dto.ToMedicineStockResponse(stock))
}

func (h *MedicineStockHandler) GetAll(c *fiber.Ctx) error {
	stocks, err := h.service.GetAllStocks()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine stocks", nil)
	}

	responses := make([]*dto.MedicineStockResponse, len(stocks))
	for i, s := range stocks {
		responses[i] = dto.ToMedicineStockResponse(s)
	}

	return utils.NewSuccessResponse(c, "Medicine stocks retrieved successfully", responses)
}

func (h *MedicineStockHandler) GetByMedicineID(c *fiber.Ctx) error {
	medicineID, err := uuid.Parse(c.Params("medicine_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine ID", nil)
	}

	stocks, err := h.service.GetStockByMedicineID(medicineID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine stocks", nil)
	}

	responses := make([]*dto.MedicineStockResponse, len(stocks))
	for i, s := range stocks {
		responses[i] = dto.ToMedicineStockResponse(s)
	}

	return utils.NewSuccessResponse(c, "Medicine stocks retrieved successfully", responses)
}

func (h *MedicineStockHandler) GetRemaining(c *fiber.Ctx) error {
	stocks, err := h.service.GetRemainingStocks()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve remaining stocks", nil)
	}

	responses := make([]*dto.MedicineStockResponse, len(stocks))
	for i, s := range stocks {
		responses[i] = dto.ToMedicineStockResponse(s)
	}

	return utils.NewSuccessResponse(c, "Remaining stocks retrieved successfully", responses)
}

func (h *MedicineStockHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine stock ID", nil)
	}

	var req dto.MedicineStockUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	stock, err := h.service.GetStockByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine stock not found", nil)
	}

	// Update fields if provided
	if req.MedicineID != nil {
		stock.MedicineID = *req.MedicineID
	}
	if req.Qty != nil {
		stock.Qty = *req.Qty
	}
	if req.ExpiredDate != nil {
		stock.ExpiredDate = req.ExpiredDate
	}

	if err := h.service.UpdateStock(stock); err != nil {
		return utils.NewErrorResponse(c, "Failed to update medicine stock", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine stock updated successfully", dto.ToMedicineStockResponse(stock))
}

func (h *MedicineStockHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine stock ID", nil)
	}

	if err := h.service.DeleteStock(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete medicine stock", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine stock deleted successfully", nil)
}
