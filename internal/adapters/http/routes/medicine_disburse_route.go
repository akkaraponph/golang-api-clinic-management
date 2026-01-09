package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMedicineDisburseRoutes(app fiber.Router, handler *handlers.MedicineDisburseHandler) {
	disburses := app.Group("/medicine-disburses")
	disburses.Get("/", handler.GetAll)
	disburses.Get("/:id", handler.GetByID)
	disburses.Get("/prescription/:prescription_id", handler.GetByPrescriptionID)
	disburses.Post("/", handler.Create)
	disburses.Put("/:id", handler.Update)
	disburses.Delete("/:id", handler.Delete)
}
