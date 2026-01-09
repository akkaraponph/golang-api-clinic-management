package routes

import (
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupReceiptRoutes(app fiber.Router, handler *handlers.ReceiptHandler) {
	receipts := app.Group("/receipts")
	receipts.Get("/", handler.GetAll)
	receipts.Get("/:id", handler.GetByID)
	receipts.Post("/", handler.Create)
	receipts.Put("/:id", handler.Update)
	receipts.Get("/:id/download", handler.DownloadPDF)
}

