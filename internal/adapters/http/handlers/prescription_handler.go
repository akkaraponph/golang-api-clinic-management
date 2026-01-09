package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PrescriptionHandler struct {
	service ports.PrescriptionService
}

func NewPrescriptionHandler(service ports.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{service: service}
}

func (h *PrescriptionHandler) Create(c *fiber.Ctx) error {
	var req dto.PrescriptionCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	prescription := dto.ToPrescriptionDomain(&req)
	if err := h.service.CreatePrescription(prescription); err != nil {
		return utils.NewErrorResponse(c, "Failed to create prescription", err.Error())
	}

	return utils.NewSuccessResponse(c, "Prescription created successfully", dto.ToPrescriptionResponse(prescription))
}

func (h *PrescriptionHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid prescription ID", nil)
	}

	prescription, err := h.service.GetPrescriptionByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Prescription not found", nil)
	}

	return utils.NewSuccessResponse(c, "Prescription retrieved successfully", dto.ToPrescriptionResponse(prescription))
}

func (h *PrescriptionHandler) GetAll(c *fiber.Ctx) error {
	prescriptions, err := h.service.GetAllPrescriptions()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve prescriptions", nil)
	}

	responses := make([]*dto.PrescriptionResponse, len(prescriptions))
	for i, p := range prescriptions {
		responses[i] = dto.ToPrescriptionResponse(p)
	}

	return utils.NewSuccessResponse(c, "Prescriptions retrieved successfully", responses)
}

func (h *PrescriptionHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patient_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	prescriptions, err := h.service.GetPrescriptionsByPatientID(patientID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve prescriptions", nil)
	}

	responses := make([]*dto.PrescriptionResponse, len(prescriptions))
	for i, p := range prescriptions {
		responses[i] = dto.ToPrescriptionResponse(p)
	}

	return utils.NewSuccessResponse(c, "Prescriptions retrieved successfully", responses)
}

func (h *PrescriptionHandler) GetByCourseID(c *fiber.Ctx) error {
	courseID, err := uuid.Parse(c.Params("course_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid course ID", nil)
	}

	prescription, err := h.service.GetPrescriptionByCourseID(courseID)
	if err != nil {
		return utils.NewErrorResponse(c, "Prescription not found", nil)
	}

	return utils.NewSuccessResponse(c, "Prescription retrieved successfully", dto.ToPrescriptionResponse(prescription))
}

func (h *PrescriptionHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid prescription ID", nil)
	}

	var req dto.PrescriptionUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	prescription, err := h.service.GetPrescriptionByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Prescription not found", nil)
	}

	// Update fields if provided
	if req.PatientID != nil {
		prescription.PatientID = *req.PatientID
	}
	if req.DoctorID != nil {
		prescription.DoctorID = *req.DoctorID
	}
	if req.CourseID != nil {
		prescription.CourseID = *req.CourseID
	}
	if req.Date != nil {
		prescription.Date = *req.Date
	}
	if req.Instructions != "" {
		prescription.Instructions = req.Instructions
	}

	if err := h.service.UpdatePrescription(prescription); err != nil {
		return utils.NewErrorResponse(c, "Failed to update prescription", err.Error())
	}

	return utils.NewSuccessResponse(c, "Prescription updated successfully", dto.ToPrescriptionResponse(prescription))
}

func (h *PrescriptionHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid prescription ID", nil)
	}

	if err := h.service.DeletePrescription(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete prescription", nil)
	}

	return utils.NewSuccessResponse(c, "Prescription deleted successfully", nil)
}
