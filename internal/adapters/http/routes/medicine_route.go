package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMedicineRoutes(app fiber.Router, handler *handlers.MedicineHandler) {
	medicines := app.Group("/medicines")
	medicines.Get("/", handler.GetAll)
	medicines.Get("/search", handler.Search)
	medicines.Get("/:id", handler.GetByID)
	medicines.Post("/", handler.Create)
	medicines.Put("/:id", handler.Update)
	medicines.Delete("/:id", handler.Delete)
	medicines.Post("/:id/restore", handler.Restore)
}
