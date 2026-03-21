package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/middleware"
)

func RegisterUserRoutes(app *fiber.App) {
	app.Get("/me", middleware.JWTMiddleware(), handlers.Me)
	/* auth := app.Group("/auth")
	limiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
	})
	auth.Post("/register", limiter, h.Register) */
}
