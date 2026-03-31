package container

import (
	"github.com/satishgowda28/ai_powered_job_tracker/internal/handlers"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/repositories"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/services"
)

type Container struct {
	/* Respository */
	UserRepository          *repositories.UserRepository
	RefreshTokenRespository *repositories.RefreshTokenRepository
	JobRepository           *repositories.JobRespository

	/* Service */
	AuthService *services.AuthService
	UserService *services.UserService
	JobService  *services.JobService

	/* Handlers */
	AuthHandler *handlers.AuthHandler
	UserHandler *handlers.UserHandler
	JobHandler  *handlers.JobHandler
}

func NewContainer() *Container {
	c := &Container{}
	/* Respository */
	c.UserRepository = repositories.NewUserRepository()
	c.RefreshTokenRespository = repositories.NewRefreshTokenRepository()
	c.JobRepository = repositories.NewJobrepository()

	/* Services */
	c.AuthService = services.NewAuthService(c.UserRepository, c.RefreshTokenRespository)
	c.UserService = services.NewUserService(c.UserRepository)
	c.JobService = services.NewJobSerive(c.JobRepository)

	/* Handlers */
	c.AuthHandler = handlers.NewAuthHandler(c.AuthService)
	c.UserHandler = handlers.NewUserHandler(c.UserService)
	c.JobHandler = handlers.NewJobHandler(c.JobService)

	return c
}
