package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAppointmentRoutes(app fiber.Router, handler *handlers.AppointmentHandler) {
	appointments := app.Group("/appointments")
	appointments.Get("/", handler.GetAll)
	appointments.Get("/:id", handler.GetByID)
	appointments.Get("/patient/:patient_id", handler.GetByPatientID)
	appointments.Post("/", handler.Create)
	appointments.Put("/:id", handler.Update)
	appointments.Delete("/:id", handler.Delete)
}
