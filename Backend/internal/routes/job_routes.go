package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/middleware"
)

func RegisterJobRoutes(app *fiber.App, h *handlers.JobHandler) {
	jobRoutes := app.Group("/job", middleware.JWTMiddleware())
	jobRoutes.Post("/", h.CreateJob)
	jobRoutes.Get("/", h.GetJobs)
	jobRoutes.Get("/:id", h.GetJob)
	jobRoutes.Put("/:id", h.UpdateJobStatus)
	/* auth := app.Group("/auth")
	limiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
	})
	auth.Post("/register", limiter, h.Register) */
}
