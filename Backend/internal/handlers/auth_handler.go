package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/auth"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/utils"
)

type BaseAuthParam struct {
	Email    string `json:"email"    form:"email"`
	Password string `json:"password" form:"password"`
}
type AuthHandler struct {
	authService *services.AuthService
}

// type RegisterParam struct {
// 	BaseAuthParam
// 	Name string `json:"name" form:"name"`
// }
// type LoginParam struct {
// 	BaseAuthParam
// }

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

func (h *AuthHandler) RegisterRoutes(router fiber.Router) {
	authRoute := router.Group("/auth")
	authRoute.Post("/register", h.Register)
	authRoute.Post("/auth/login", h.Login)
	authRoute.Post("/auth/refresh", h.Refresh)
	authRoute.Post("/logout", h.Logout)
}

func (authHandler *AuthHandler) Register(c *fiber.Ctx) error {
	var newCreds struct {
		BaseAuthParam
		Name string `json:"name" form:"name"`
	}
	if err := c.BodyParser(&newCreds); err != nil {
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
	user, err := authHandler.authService.Register(
		c.Context(),
		newCreds.Name,
		newCreds.Email,
		newCreds.Password,
	)
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
	var loginCreds struct {
		BaseAuthParam
	}
	if err := c.BodyParser(&loginCreds); err != nil {
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

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var rfTknBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BodyParser(&rfTknBody); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "invalid request body",
		})
	}
	if rfTknBody.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": "refresh token required",
		})
	}
	accessToken, err := h.authService.Refresh(c.Context(), rfTknBody.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "bad_request",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"access_token": accessToken})

}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BodyParser(&body); err != nil {
		return utils.BadRequest(c, "invalid request body")
	}
	if body.RefreshToken == "" {
		return utils.Unauthorized(c, "refresh token is required")
	}
	if err := h.authService.HandleLogout(c.Context(), body.RefreshToken); err != nil {
		return utils.BadRequest(c, "Something went wrong")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out successfully",
	})
}
