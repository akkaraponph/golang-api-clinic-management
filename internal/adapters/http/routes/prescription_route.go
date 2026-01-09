package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPrescriptionRoutes(app fiber.Router, handler *handlers.PrescriptionHandler) {
	prescriptions := app.Group("/prescriptions")
	prescriptions.Get("/", handler.GetAll)
	prescriptions.Get("/:id", handler.GetByID)
	prescriptions.Get("/patient/:patient_id", handler.GetByPatientID)
	prescriptions.Get("/course/:course_id", handler.GetByCourseID)
	prescriptions.Post("/", handler.Create)
	prescriptions.Put("/:id", handler.Update)
	prescriptions.Delete("/:id", handler.Delete)
}
