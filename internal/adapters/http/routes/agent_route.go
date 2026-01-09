package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAgentRoutes(app fiber.Router, handler *handlers.AgentHandler) {
	agents := app.Group("/agents")
	agents.Get("/", handler.GetAll)
	agents.Get("/:id", handler.GetByID)
	agents.Post("/", handler.Create)
	agents.Put("/:id", handler.Update)
	agents.Delete("/:id", handler.Delete)
}
