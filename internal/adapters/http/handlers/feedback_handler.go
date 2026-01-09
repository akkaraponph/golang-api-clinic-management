package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FeedbackHandler struct {
	service ports.FeedbackService
}

func NewFeedbackHandler(service ports.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{service: service}
}

func (h *FeedbackHandler) Create(c *fiber.Ctx) error {
	var req dto.FeedbackCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	feedback := dto.ToFeedbackDomain(&req)
	if err := h.service.CreateFeedback(feedback); err != nil {
		return utils.NewErrorResponse(c, "Failed to create feedback", err.Error())
	}

	return utils.NewSuccessResponse(c, "Feedback created successfully", dto.ToFeedbackResponse(feedback))
}

func (h *FeedbackHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid feedback ID", nil)
	}

	feedback, err := h.service.GetFeedbackByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Feedback not found", nil)
	}

	return utils.NewSuccessResponse(c, "Feedback retrieved successfully", dto.ToFeedbackResponse(feedback))
}

func (h *FeedbackHandler) GetAll(c *fiber.Ctx) error {
	feedbacks, err := h.service.GetAllFeedbacks()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve feedbacks", nil)
	}

	responses := make([]*dto.FeedbackResponse, len(feedbacks))
	for i, f := range feedbacks {
		responses[i] = dto.ToFeedbackResponse(f)
	}

	return utils.NewSuccessResponse(c, "Feedbacks retrieved successfully", responses)
}

func (h *FeedbackHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patient_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	feedbacks, err := h.service.GetFeedbacksByPatientID(patientID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve feedbacks", nil)
	}

	responses := make([]*dto.FeedbackResponse, len(feedbacks))
	for i, f := range feedbacks {
		responses[i] = dto.ToFeedbackResponse(f)
	}

	return utils.NewSuccessResponse(c, "Feedbacks retrieved successfully", responses)
}

func (h *FeedbackHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid feedback ID", nil)
	}

	var req dto.FeedbackUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	feedback, err := h.service.GetFeedbackByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Feedback not found", nil)
	}

	if req.PatientID != nil {
		feedback.PatientID = *req.PatientID
	}
	if req.FeedbackText != "" {
		feedback.FeedbackText = req.FeedbackText
	}

	if err := h.service.UpdateFeedback(feedback); err != nil {
		return utils.NewErrorResponse(c, "Failed to update feedback", err.Error())
	}

	return utils.NewSuccessResponse(c, "Feedback updated successfully", dto.ToFeedbackResponse(feedback))
}

func (h *FeedbackHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid feedback ID", nil)
	}

	if err := h.service.DeleteFeedback(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete feedback", nil)
	}

	return utils.NewSuccessResponse(c, "Feedback deleted successfully", nil)
}
