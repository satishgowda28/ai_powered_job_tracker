package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

type RegisterParam struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AuthResponse struct {
	User         generated.User `json:"user"`
	AccessToken  string         `json:"accesstoken"`
	RefreshToken string         `json:"refreshtoken"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (authService *AuthHandler) Register(c *fiber.Ctx) error {
	newCreds := new(RegisterParam)
	if err := c.BodyParser(newCreds); err != nil {
		return err
	}
	return c.JSON(AuthResponse{})
}
func (authService *AuthHandler) Login(c *fiber.Ctx) error {
	newCreds := new(RegisterParam)
	if err := c.BodyParser(newCreds); err != nil {
		return err
	}
	return c.JSON(AuthResponse{})
}
