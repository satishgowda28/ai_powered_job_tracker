package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
)

func Register(app *fiber.App) {
	app.Get("/health", handlers.Health)
}
