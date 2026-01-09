package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMaterialRoutes(app fiber.Router, handler *handlers.MaterialHandler) {
	materials := app.Group("/materials")
	materials.Get("/", handler.GetAll)
	materials.Get("/:id", handler.GetByID)
	materials.Post("/", handler.Create)
	materials.Put("/:id", handler.Update)
	materials.Delete("/:id", handler.Delete)
}
