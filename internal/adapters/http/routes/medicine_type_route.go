package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMedicineTypeRoutes(app fiber.Router, handler *handlers.MedicineTypeHandler) {
	medicineTypes := app.Group("/medicine-types")
	medicineTypes.Get("/", handler.GetAll)
	medicineTypes.Get("/:id", handler.GetByID)
	medicineTypes.Get("/agent/:agent_id", handler.GetByAgentID)
	medicineTypes.Post("/", handler.Create)
	medicineTypes.Put("/:id", handler.Update)
	medicineTypes.Delete("/:id", handler.Delete)
}
