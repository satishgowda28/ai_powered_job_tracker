package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/auth"
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

func (authHandler *AuthHandler) Register(c *fiber.Ctx) error {
	newCreds := new(RegisterParam)
	if err := c.BodyParser(newCreds); err != nil {
		return err
	}
	/* Errors */
	if newCreds.Name == "" {
		return errors.New("user name is required")
	}
	if newCreds.Email == "" {
		return errors.New("user email is required")
	}
	if newCreds.Password == "" {
		return errors.New("user password is required")
	}

	/* register user */
	user, err := authHandler.authService.Register(c.Context(), newCreds.Name, newCreds.Email, newCreds.Password)
	if err != nil {
		return err
	}
	/* generate refreshtoke */
	rfToken, err := authHandler.authService.NewRefreshToken(c.Context(), user.ID)
	if err != nil {
		return err
	}
	/* Access token */
	token, err := auth.GenerateAccessToken(user.ID.Bytes)
	if err != nil {
		return err
	}

	return c.JSON(AuthResponse{
		User:         user,
		RefreshToken: rfToken.Token,
		AccessToken:  token,
	})
}
func (authHandler *AuthHandler) Login(c *fiber.Ctx) error {
	newCreds := new(RegisterParam)
	if err := c.BodyParser(newCreds); err != nil {
		return err
	}
	return c.JSON(AuthResponse{})
}
