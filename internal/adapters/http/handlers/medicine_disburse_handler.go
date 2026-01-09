package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MedicineDisburseHandler struct {
	service ports.MedicineDisburseService
}

func NewMedicineDisburseHandler(service ports.MedicineDisburseService) *MedicineDisburseHandler {
	return &MedicineDisburseHandler{service: service}
}

func (h *MedicineDisburseHandler) Create(c *fiber.Ctx) error {
	var req dto.MedicineDisburseCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	// Create disburse first
	disburse := &domain.MedicineDisburse{
		PrescriptionID: req.PrescriptionID,
		Date:           req.Date,
	}

	if err := h.service.CreateMedicineDisburse(disburse); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine disburse", err.Error())
	}

	// Set disburse ID in details
	for _, d := range req.Details {
		disburse.MedicineDisburseDetails = append(disburse.MedicineDisburseDetails, domain.MedicineDisburseDetail{
			MedicineDisburseID: disburse.ID,
			MedicineStockID:    d.MedicineStockID,
			Price:              d.Price,
			Unit:               d.Unit,
			Qty:                d.Qty,
			Dosage:             d.Dosage,
			AdminMethod:        d.AdminMethod,
			ApplyMethod:        d.ApplyMethod,
			TimeOfAdmin:        d.TimeOfAdmin,
		})
	}

	// Update disburse with details
	if err := h.service.UpdateMedicineDisburse(disburse); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine disburse details", err.Error())
	}

	// Fetch full disburse with details
	fullDisburse, err := h.service.GetMedicineDisburseByID(disburse.ID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine disburse", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine disburse created successfully", dto.ToMedicineDisburseResponse(fullDisburse))
}

func (h *MedicineDisburseHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine disburse ID", nil)
	}

	disburse, err := h.service.GetMedicineDisburseByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine disburse not found", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine disburse retrieved successfully", dto.ToMedicineDisburseResponse(disburse))
}

func (h *MedicineDisburseHandler) GetAll(c *fiber.Ctx) error {
	disburses, err := h.service.GetAllMedicineDisburses()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine disburses", nil)
	}

	responses := make([]*dto.MedicineDisburseResponse, len(disburses))
	for i, d := range disburses {
		responses[i] = dto.ToMedicineDisburseResponse(d)
	}

	return utils.NewSuccessResponse(c, "Medicine disburses retrieved successfully", responses)
}

func (h *MedicineDisburseHandler) GetByPrescriptionID(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("prescription_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid prescription ID", nil)
	}

	disburses, err := h.service.GetMedicineDisbursesByPrescriptionID(prescriptionID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicine disburses", nil)
	}

	responses := make([]*dto.MedicineDisburseResponse, len(disburses))
	for i, d := range disburses {
		responses[i] = dto.ToMedicineDisburseResponse(d)
	}

	return utils.NewSuccessResponse(c, "Medicine disburses retrieved successfully", responses)
}

func (h *MedicineDisburseHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine disburse ID", nil)
	}

	var req dto.MedicineDisburseCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	disburse, err := h.service.GetMedicineDisburseByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine disburse not found", nil)
	}

	disburse.PrescriptionID = req.PrescriptionID
	disburse.Date = req.Date

	// Update details
	disburse.MedicineDisburseDetails = make([]domain.MedicineDisburseDetail, len(req.Details))
	for i, d := range req.Details {
		disburse.MedicineDisburseDetails[i] = domain.MedicineDisburseDetail{
			MedicineDisburseID: disburse.ID,
			MedicineStockID:    d.MedicineStockID,
			Price:              d.Price,
			Unit:               d.Unit,
			Qty:                d.Qty,
			Dosage:             d.Dosage,
			AdminMethod:        d.AdminMethod,
			ApplyMethod:        d.ApplyMethod,
			TimeOfAdmin:        d.TimeOfAdmin,
		}
	}

	if err := h.service.UpdateMedicineDisburse(disburse); err != nil {
		return utils.NewErrorResponse(c, "Failed to update medicine disburse", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine disburse updated successfully", dto.ToMedicineDisburseResponse(disburse))
}

func (h *MedicineDisburseHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine disburse ID", nil)
	}

	if err := h.service.DeleteMedicineDisburse(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete medicine disburse", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine disburse deleted successfully", nil)
}
