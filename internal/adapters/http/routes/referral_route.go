package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupReferralRoutes(app fiber.Router, handler *handlers.ReferralHandler) {
	referrals := app.Group("/referrals")
	referrals.Get("/", handler.GetAll)
	referrals.Get("/:id", handler.GetByID)
	referrals.Get("/patient/:patient_id", handler.GetByPatientID)
	referrals.Post("/", handler.Create)
	referrals.Put("/:id", handler.Update)
	referrals.Delete("/:id", handler.Delete)
}
