package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupDoctorRoutes(app fiber.Router, handler *handlers.DoctorHandler) {
	doctors := app.Group("/doctors")
	doctors.Get("/", handler.GetAll)
	doctors.Get("/:id", handler.GetByID)
	doctors.Get("/employee/:employee_id", handler.GetByEmployeeID)
	doctors.Post("/", handler.Create)
	doctors.Put("/:id", handler.Update)
	doctors.Delete("/:id", handler.Delete)
}
