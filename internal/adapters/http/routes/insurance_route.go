package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupInsuranceRoutes(app fiber.Router, handler *handlers.InsuranceHandler) {
	insurances := app.Group("/insurances")
	insurances.Get("/", handler.GetAll)
	insurances.Get("/:id", handler.GetByID)
	insurances.Get("/patient/:patient_id", handler.GetByPatientID)
	insurances.Post("/", handler.Create)
	insurances.Put("/:id", handler.Update)
	insurances.Delete("/:id", handler.Delete)
}
