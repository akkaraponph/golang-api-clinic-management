package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
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

func (h *ReceiptHandler) Create(c *fiber.Ctx) error {
	var req dto.ReceiptCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	receipt := dto.ToReceiptDomain(&req)
	if err := h.service.CreateReceipt(receipt); err != nil {
		return utils.NewErrorResponse(c, "Failed to create receipt", err.Error())
	}

	return utils.NewSuccessResponse(c, "Receipt created successfully", dto.ToReceiptResponse(receipt))
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

	return utils.NewSuccessResponse(c, "Receipt retrieved successfully", dto.ToReceiptResponse(receipt))
}

func (h *ReceiptHandler) GetAll(c *fiber.Ctx) error {
	receipts, err := h.service.GetAllReceipts()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve receipts", nil)
	}

	responses := make([]*dto.ReceiptResponse, len(receipts))
	for i, r := range receipts {
		responses[i] = dto.ToReceiptResponse(r)
	}

	return utils.NewSuccessResponse(c, "Receipts retrieved successfully", responses)
}

func (h *ReceiptHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid receipt ID", nil)
	}

	var req dto.ReceiptUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	receipt, err := h.service.GetReceiptByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Receipt not found", nil)
	}

	// Update fields if provided
	if req.Date != nil {
		receipt.Date = *req.Date
	}
	if req.TotalPrice != nil {
		receipt.TotalPrice = *req.TotalPrice
	}
	if req.CourseID != nil {
		receipt.CourseID = req.CourseID
	}
	if req.MedicineDisburseID != nil {
		receipt.MedicineDisburseID = req.MedicineDisburseID
	}
	if req.PatientID != nil {
		receipt.PatientID = *req.PatientID
	}
	if req.EmployeeID != nil {
		receipt.EmployeeID = *req.EmployeeID
	}

	if err := h.service.UpdateReceipt(receipt); err != nil {
		return utils.NewErrorResponse(c, "Failed to update receipt", err.Error())
	}

	return utils.NewSuccessResponse(c, "Receipt updated successfully", dto.ToReceiptResponse(receipt))
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

