package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/utils"
)

type UserHandler struct {
	usrService *services.UserService
}

func NewUserHandler(uService *services.UserService) *UserHandler {
	return &UserHandler{
		usrService: uService,
	}
}

func (usrHandler *UserHandler) Me(c *fiber.Ctx) error {
	userId, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server_error",
			"message": "user id missing from context"})
	}
	uId := utils.ToPgtypeUUID(userId)
	user, err := usrHandler.usrService.GetUserDetails(c.Context(), uId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not_found",
			"message": "user not found"})
	}

	return c.JSON(fiber.Map{"Status": "OK", "data": fiber.Map{"name": user.Name, "email": user.Email, "userId": user.ID}})
}
