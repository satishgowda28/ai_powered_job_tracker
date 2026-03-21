package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/auth"
)

const authHeader = "authorization"

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Get(authHeader)
		fmt.Println(authToken)
		if authToken == "" {
			return unauthorized(c, "missing auth token")
		}
		parts := strings.SplitN(authToken, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return unauthorized(c, "invalid format for token")
		}
		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			return unauthorized(c, "token missing")
		}
		userId, err := auth.ValidateToken(tokenString)
		if err != nil {
			return unauthorized(c, "invald or expired token")
		}
		c.Locals("userID", userId.String())
		return c.Next()
	}
}
func unauthorized(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized", "message": msg})
}
