package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	userId, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server_error",
			"message": "user id missing from context"})
	}
	println(userId)
	return c.JSON(fiber.Map{"Status": "OK", "Message": userId})
}
