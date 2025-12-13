package container

import (
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/respositories"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
)

type Container struct {
	/* Respository */
	UserRepository          *respositories.UserRepository
	RefreshTokenRespository *respositories.RefreshTokenRepository

	/* Service */
	AuthService *services.AuthService

	/* Handlers */
	AuthHandler *handlers.AuthHandler
}

func NewContainer() *Container {
	c := &Container{}
	/* Respository */
	c.UserRepository = respositories.NewUserRepository()
	c.RefreshTokenRespository = respositories.NewRefreshTokenRepository()

	/* Services */
	c.AuthService = services.NewAuthService(c.UserRepository, c.RefreshTokenRespository)

	/* Handlers */
	c.AuthHandler = handlers.NewAuthHandler(c.AuthService)

	return c
}
