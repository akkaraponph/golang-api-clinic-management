package handlers

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/dto"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
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

func (h *EmployeeHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid employee ID", nil)
	}

	var req dto.EmployeeUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.NewErrorResponse(c, "Invalid request body", nil)
	}

	employee, err := h.service.GetEmployeeByID(id)
	if err != nil {
		return utils.NewErrorResponse(c, "Employee not found", nil)
	}

	// Update fields if provided
	if req.Name != "" {
		employee.Name = req.Name
	}
	if req.Surname != "" {
		employee.Surname = req.Surname
	}
	if req.Email != "" {
		employee.Email = req.Email
	}
	if req.Username != "" {
		employee.Username = req.Username
	}
	if req.Password != "" {
		employee.Password = req.Password
	}
	if req.Sex != "" {
		employee.Sex = domain.Sex(req.Sex)
	}
	if req.Phone != "" {
		employee.Phone = req.Phone
	}
	if req.Permission != "" {
		employee.Permission = domain.PermissionEnum(req.Permission)
	}
	if req.Address != "" {
		employee.Address = req.Address
	}
	if req.Subdistrict != "" {
		employee.Subdistrict = req.Subdistrict
	}
	if req.District != "" {
		employee.District = req.District
	}
	if req.Province != "" {
		employee.Province = req.Province
	}
	if req.Zipcode != "" {
		employee.Zipcode = req.Zipcode
	}

	if err := h.service.UpdateEmployee(employee); err != nil {
		return utils.NewErrorResponse(c, "Failed to update employee", err.Error())
	}

	return utils.NewSuccessResponse(c, "Employee updated successfully", dto.ToEmployeeResponse(employee))
}

func (h *EmployeeHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid employee ID", nil)
	}

	if err := h.service.DeleteEmployee(id); err != nil {
		return utils.NewErrorResponse(c, "Failed to delete employee", nil)
	}

	return utils.NewSuccessResponse(c, "Employee deleted successfully", nil)
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

