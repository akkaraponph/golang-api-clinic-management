package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMedicineOrderRoutes(app fiber.Router, handler *handlers.MedicineOrderHandler) {
	orders := app.Group("/medicine-orders")
	orders.Get("/", handler.GetAll)
	orders.Get("/:id", handler.GetByID)
	orders.Post("/", handler.Create)
	orders.Put("/:id", handler.Update)
	orders.Delete("/:id", handler.Delete)
}
