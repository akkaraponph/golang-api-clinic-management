package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MedicineHandler struct {
	service ports.MedicineService
}

func NewMedicineHandler(service ports.MedicineService) *MedicineHandler {
	return &MedicineHandler{service: service}
}

func (h *MedicineHandler) Create(c *fiber.Ctx) error {
	var req dto.MedicineCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	medicine := dto.ToMedicineDomain(&req)
	if err := h.service.CreateMedicine(medicine); err != nil {
		return utils.NewErrorResponse(c, "Failed to create medicine", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine created successfully", dto.ToMedicineResponse(medicine))
}

func (h *MedicineHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine ID", nil)
	}

	medicine, err := h.service.GetMedicineByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine not found", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine retrieved successfully", dto.ToMedicineResponse(medicine))
}

func (h *MedicineHandler) GetAll(c *fiber.Ctx) error {
	medicines, err := h.service.GetAllMedicines()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve medicines", nil)
	}

	responses := make([]*dto.MedicineResponse, len(medicines))
	for i, m := range medicines {
		responses[i] = dto.ToMedicineResponse(m)
	}

	return utils.NewSuccessResponse(c, "Medicines retrieved successfully", responses)
}

func (h *MedicineHandler) Search(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return utils.NewErrorResponse(c, "Query parameter is required", nil)
	}

	medicines, err := h.service.SearchMedicines(query)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to search medicines", nil)
	}

	responses := make([]*dto.MedicineResponse, len(medicines))
	for i, m := range medicines {
		responses[i] = dto.ToMedicineResponse(m)
	}

	return utils.NewSuccessResponse(c, "Search completed successfully", responses)
}

func (h *MedicineHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine ID", nil)
	}

	var req dto.MedicineUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	medicine, err := h.service.GetMedicineByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Medicine not found", nil)
	}

	// Update fields if provided
	if req.Name != "" {
		medicine.Name = req.Name
	}
	if req.Detail != "" {
		medicine.Detail = req.Detail
	}
	if req.Dosage != "" {
		medicine.Dosage = req.Dosage
	}
	if req.Unit != "" {
		medicine.Unit = req.Unit
	}
	if req.MedicineTypeID != nil {
		medicine.MedicineTypeID = req.MedicineTypeID
	}

	if err := h.service.UpdateMedicine(medicine); err != nil {
		return utils.NewErrorResponse(c, "Failed to update medicine", err.Error())
	}

	return utils.NewSuccessResponse(c, "Medicine updated successfully", dto.ToMedicineResponse(medicine))
}

func (h *MedicineHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine ID", nil)
	}

	if err := h.service.DeleteMedicine(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete medicine", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine deleted successfully", nil)
}

func (h *MedicineHandler) Restore(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid medicine ID", nil)
	}

	if err := h.service.RestoreMedicine(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to restore medicine", nil)
	}

	return utils.NewSuccessResponse(c, "Medicine restored successfully", nil)
}
