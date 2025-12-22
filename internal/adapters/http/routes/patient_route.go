package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPatientRoutes(app fiber.Router, handler *handlers.PatientHandler) {
	patients := app.Group("/patients")
	patients.Get("/", handler.GetAll)
	patients.Get("/search", handler.Search)
	patients.Get("/:id", handler.GetByID)
	patients.Post("/", handler.Create)
	patients.Put("/:id", handler.Update)
	patients.Delete("/:id", handler.Delete)
}

