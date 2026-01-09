package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMaterialDisburseRoutes(app fiber.Router, handler *handlers.MaterialDisburseHandler) {
	disburses := app.Group("/material-disburses")
	disburses.Get("/", handler.GetAll)
	disburses.Get("/:id", handler.GetByID)
	disburses.Post("/", handler.Create)
	disburses.Put("/:id", handler.Update)
	disburses.Delete("/:id", handler.Delete)
}
