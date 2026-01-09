package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMaterialOrderRoutes(app fiber.Router, handler *handlers.MaterialOrderHandler) {
	orders := app.Group("/material-orders")
	orders.Get("/", handler.GetAll)
	orders.Get("/:id", handler.GetByID)
	orders.Post("/", handler.Create)
	orders.Put("/:id", handler.Update)
	orders.Delete("/:id", handler.Delete)
}
