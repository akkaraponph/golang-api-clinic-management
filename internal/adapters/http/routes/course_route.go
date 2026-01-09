package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupCourseRoutes(app fiber.Router, handler *handlers.CourseHandler) {
	courses := app.Group("/courses")
	courses.Get("/", handler.GetAll)
	courses.Get("/trash", handler.GetTrash)
	courses.Get("/:id", handler.GetByID)
	courses.Get("/patient/:patient_id", handler.GetByPatientID)
	courses.Post("/", handler.Create)
	courses.Put("/:id", handler.Update)
	courses.Delete("/:id", handler.Delete)
	courses.Post("/:id/restore", handler.Restore)
}
