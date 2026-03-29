package utils

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	error   string
	message string
}

func HandleErrorResponse(
	c *fiber.Ctx,
	status int,
	error string,
	message string,
) error {
	return c.Status(status).JSON(ErrorResponse{
		error,
		message,
	})
}

func BadRequest(c *fiber.Ctx, message string) error {
	return HandleErrorResponse(c, fiber.StatusBadRequest, "bad_request", message)
}

func ServerError(c *fiber.Ctx, message string) error {
	return HandleErrorResponse(c, fiber.StatusInternalServerError, "server_erro", message)
}
func NotFound(c *fiber.Ctx, message string) error {
	return HandleErrorResponse(c, fiber.StatusNotFound, "not_found", message)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return HandleErrorResponse(c, fiber.StatusUnauthorized, "unauthorized", message)
}
