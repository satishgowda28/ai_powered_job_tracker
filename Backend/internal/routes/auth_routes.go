package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
)

// RegisterAuthRoutes registers authentication endpoints on the given Fiber app.
// It maps POST /auth/register to h.Register and POST /auth/login to h.Login.
func RegisterAuthRoutes(app *fiber.App, h *handlers.AuthHandler) {
	app.Post("/auth/register", h.Register)
	app.Post("/auth/login", h.Login)
}