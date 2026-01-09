package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReferralHandler struct {
	service ports.ReferralService
}

func NewReferralHandler(service ports.ReferralService) *ReferralHandler {
	return &ReferralHandler{service: service}
}

func (h *ReferralHandler) Create(c *fiber.Ctx) error {
	var req dto.ReferralCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	referral := dto.ToReferralDomain(&req)
	if err := h.service.CreateReferral(referral); err != nil {
		return utils.NewErrorResponse(c, "Failed to create referral", err.Error())
	}

	return utils.NewSuccessResponse(c, "Referral created successfully", dto.ToReferralResponse(referral))
}

func (h *ReferralHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid referral ID", nil)
	}

	referral, err := h.service.GetReferralByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Referral not found", nil)
	}

	return utils.NewSuccessResponse(c, "Referral retrieved successfully", dto.ToReferralResponse(referral))
}

func (h *ReferralHandler) GetAll(c *fiber.Ctx) error {
	referrals, err := h.service.GetAllReferrals()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve referrals", nil)
	}

	responses := make([]*dto.ReferralResponse, len(referrals))
	for i, r := range referrals {
		responses[i] = dto.ToReferralResponse(r)
	}

	return utils.NewSuccessResponse(c, "Referrals retrieved successfully", responses)
}

func (h *ReferralHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patient_id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid patient ID", nil)
	}

	referrals, err := h.service.GetReferralsByPatientID(patientID)
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve referrals", nil)
	}

	responses := make([]*dto.ReferralResponse, len(referrals))
	for i, r := range referrals {
		responses[i] = dto.ToReferralResponse(r)
	}

	return utils.NewSuccessResponse(c, "Referrals retrieved successfully", responses)
}

func (h *ReferralHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid referral ID", nil)
	}

	var req dto.ReferralUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	referral, err := h.service.GetReferralByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Referral not found", nil)
	}

	if req.PatientID != nil {
		referral.PatientID = *req.PatientID
	}
	if req.ReferredDoctorID != nil {
		referral.ReferredDoctorID = *req.ReferredDoctorID
	}
	if req.Reason != "" {
		referral.Reason = req.Reason
	}
	if req.ReferralDate != nil {
		referral.ReferralDate = *req.ReferralDate
	}

	if err := h.service.UpdateReferral(referral); err != nil {
		return utils.NewErrorResponse(c, "Failed to update referral", err.Error())
	}

	return utils.NewSuccessResponse(c, "Referral updated successfully", dto.ToReferralResponse(referral))
}

func (h *ReferralHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid referral ID", nil)
	}

	if err := h.service.DeleteReferral(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete referral", nil)
	}

	return utils.NewSuccessResponse(c, "Referral deleted successfully", nil)
}
