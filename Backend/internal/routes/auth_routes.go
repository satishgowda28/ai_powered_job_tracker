package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
)

func RegisterAuthRoutes(app *fiber.App, h *handlers.AuthHandler) {
	app.Post("/auth/register", h.Register)
	app.Post("/auth/login", h.Login)
	/* auth := app.Group("/auth")
	limiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
	})
	auth.Post("/register", limiter, h.Register) */
}
