package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMedicineStockRoutes(app fiber.Router, handler *handlers.MedicineStockHandler) {
	stocks := app.Group("/medicine-stocks")
	stocks.Get("/", handler.GetAll)
	stocks.Get("/remaining", handler.GetRemaining)
	stocks.Get("/:id", handler.GetByID)
	stocks.Get("/medicine/:medicine_id", handler.GetByMedicineID)
	stocks.Post("/", handler.Create)
	stocks.Put("/:id", handler.Update)
	stocks.Delete("/:id", handler.Delete)
}
