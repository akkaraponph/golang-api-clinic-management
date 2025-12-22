package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type EmployeeHandler struct {
	service ports.EmployeeService
}

func NewEmployeeHandler(service ports.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) Create(c *fiber.Ctx) error {
	var req dto.EmployeeCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	employee := dto.ToEmployeeDomain(&req)
	if err := h.service.CreateEmployee(employee); err != nil {
		return utils.NewErrorResponse(c, "Failed to create employee", err.Error())
	}

	return utils.NewSuccessResponse(c, "Employee created successfully", dto.ToEmployeeResponse(employee))
}

func (h *EmployeeHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid employee ID", nil)
	}

	employee, err := h.service.GetEmployeeByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Employee not found", nil)
	}

	return utils.NewSuccessResponse(c, "Employee retrieved successfully", dto.ToEmployeeResponse(employee))
}

func (h *EmployeeHandler) GetAll(c *fiber.Ctx) error {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		return utils.NewErrorResponse(c, "Failed to retrieve employees", nil)
	}

	responses := make([]*dto.EmployeeResponse, len(employees))
	for i, e := range employees {
		responses[i] = dto.ToEmployeeResponse(e)
	}

	return utils.NewSuccessResponse(c, "Employees retrieved successfully", responses)
}

func (h *EmployeeHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	employee, err := h.service.Authenticate(req.Username, req.Password)
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid credentials", nil)
	}

	return utils.NewSuccessResponse(c, "Login successful", dto.ToEmployeeResponse(employee))
}

