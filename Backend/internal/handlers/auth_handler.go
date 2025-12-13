package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/auth"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

type BaseAuthParam struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type RegisterParam struct {
	BaseAuthParam
	Name string `json:"name" form:"name"`
}
type LoginParam struct {
	BaseAuthParam
}

/* Keep adding aditional info for use if required */
type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	User         UserData `json:"user"`
	AccessToken  string   `json:"accesstoken"`
	RefreshToken string   `json:"refreshtoken"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (authHandler *AuthHandler) Register(c *fiber.Ctx) error {
	newCreds := new(RegisterParam)
	if err := c.BodyParser(newCreds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "invalid JSON body",
		})
	}
	/* Errors */
	if newCreds.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "user name is required",
		})
	}
	if newCreds.Email == "" || newCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "Email and Password is required",
		})
	}

	/* register user */
	user, err := authHandler.authService.Register(c.Context(), newCreds.Name, newCreds.Email, newCreds.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}
	/* generate refreshtoke */
	rfToken, err := authHandler.authService.NewRefreshToken(c.Context(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}
	/* Access token */
	token, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}

	return c.JSON(AuthResponse{
		User: UserData{
			Name:  user.Name,
			Email: user.Email,
		},
		RefreshToken: rfToken.Token,
		AccessToken:  token,
	})
}

func (authHandler *AuthHandler) Login(c *fiber.Ctx) error {
	loginCreds := new(LoginParam)
	if err := c.BodyParser(loginCreds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "invalid JSON body",
		})
	}
	/* errors */
	if loginCreds.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "User Email is required",
		})
	}
	if loginCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "Password is required",
		})
	}
	/* check user */
	user, err := authHandler.authService.Login(c.Context(), loginCreds.Email, loginCreds.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}
	/* generate refreshtoken */
	rfToken, err := authHandler.authService.NewRefreshToken(c.Context(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}
	/* Access token */
	token, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}

	return c.JSON(AuthResponse{User: UserData{
		Name:  user.Name,
		Email: user.Email,
	},
		RefreshToken: rfToken.Token,
		AccessToken:  token})
}
