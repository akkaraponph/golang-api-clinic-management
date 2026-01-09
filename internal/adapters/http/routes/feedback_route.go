package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupFeedbackRoutes(app fiber.Router, handler *handlers.FeedbackHandler) {
	feedbacks := app.Group("/feedbacks")
	feedbacks.Get("/", handler.GetAll)
	feedbacks.Get("/:id", handler.GetByID)
	feedbacks.Get("/patient/:patient_id", handler.GetByPatientID)
	feedbacks.Post("/", handler.Create)
	feedbacks.Put("/:id", handler.Update)
	feedbacks.Delete("/:id", handler.Delete)
}
