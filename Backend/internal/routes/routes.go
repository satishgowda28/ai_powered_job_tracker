package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/container"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
)

// func Register(app *fiber.App) {
// 	app.Get("/health", handlers.Health)
// }

func SetupRoutes(app *fiber.App, c *container.Container) {

	api := app.Group("/api")

	api.Get("/health", handlers.Health)

	c.AuthHandler.RegisterRoutes(api)
	c.UserHandler.RegisterRoutes(api)
	c.JobHandler.RegisterRoutes(api)
}
