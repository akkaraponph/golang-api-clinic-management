package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AgentHandler struct {
	service ports.AgentService
}

func NewAgentHandler(service ports.AgentService) *AgentHandler {
	return &AgentHandler{service: service}
}

func (h *AgentHandler) Create(c *fiber.Ctx) error {
	var req dto.AgentCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	agent := dto.ToAgentDomain(&req)
	if err := h.service.CreateAgent(agent); err != nil {
		return utils.NewErrorResponse(c, "Failed to create agent", err.Error())
	}

	return utils.NewSuccessResponse(c, "Agent created successfully", dto.ToAgentResponse(agent))
}

func (h *AgentHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid agent ID", nil)
	}

	agent, err := h.service.GetAgentByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Agent not found", nil)
	}

	return utils.NewSuccessResponse(c, "Agent retrieved successfully", dto.ToAgentResponse(agent))
}

func (h *AgentHandler) GetAll(c *fiber.Ctx) error {
	agents, err := h.service.GetAllAgents()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve agents", nil)
	}

	responses := make([]*dto.AgentResponse, len(agents))
	for i, a := range agents {
		responses[i] = dto.ToAgentResponse(a)
	}

	return utils.NewSuccessResponse(c, "Agents retrieved successfully", responses)
}

func (h *AgentHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid agent ID", nil)
	}

	var req dto.AgentUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	agent, err := h.service.GetAgentByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Agent not found", nil)
	}

	if req.Name != "" {
		agent.Name = req.Name
	}
	if req.Address != "" {
		agent.Address = req.Address
	}
	if req.Phone != "" {
		agent.Phone = req.Phone
	}
	if req.Email != "" {
		agent.Email = req.Email
	}

	if err := h.service.UpdateAgent(agent); err != nil {
		return utils.NewErrorResponse(c, "Failed to update agent", err.Error())
	}

	return utils.NewSuccessResponse(c, "Agent updated successfully", dto.ToAgentResponse(agent))
}

func (h *AgentHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid agent ID", nil)
	}

	if err := h.service.DeleteAgent(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete agent", nil)
	}

	return utils.NewSuccessResponse(c, "Agent deleted successfully", nil)
}
