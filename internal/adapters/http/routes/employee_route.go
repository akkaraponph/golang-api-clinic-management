package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupEmployeeRoutes(app fiber.Router, handler *handlers.EmployeeHandler) {
	employees := app.Group("/employees")
	employees.Get("/", handler.GetAll)
	employees.Get("/:id", handler.GetByID)
	employees.Post("/", handler.Create)
	
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
}

