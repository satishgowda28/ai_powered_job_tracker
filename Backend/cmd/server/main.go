package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/config"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/container"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/database"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/routes"
)

// and starts the Fiber HTTP server on the configured port, logging startup and fatal errors.
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file present")
	}

	cfg := config.LoadConfig()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	database.Connect()

	container := container.NewContainer()

	routes.Register(app)
	routes.RegisterAuthRoutes(app, container.AuthHandler)

	addr := ":" + cfg.Port
	log.Println("Starting server on", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start a server %v", err)
	}
}