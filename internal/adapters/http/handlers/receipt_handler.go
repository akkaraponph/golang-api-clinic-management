package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReceiptHandler struct {
	service ports.ReceiptService
}

func NewReceiptHandler(service ports.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{service: service}
}

func (h *ReceiptHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid receipt ID", nil)
	}

	receipt, err := h.service.GetReceiptByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Receipt not found", nil)
	}

	return utils.NewSuccessResponse(c, "Receipt retrieved successfully", receipt)
}

func (h *ReceiptHandler) GetAll(c *fiber.Ctx) error {
	receipts, err := h.service.GetAllReceipts()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve receipts", nil)
	}

	return utils.NewSuccessResponse(c, "Receipts retrieved successfully", receipts)
}

func (h *ReceiptHandler) DownloadPDF(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid receipt ID", nil)
	}

	pdfBytes, err := h.service.GenerateReceiptPDF(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to generate PDF", nil)
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=receipt.pdf")
	return c.Send(pdfBytes)
}

