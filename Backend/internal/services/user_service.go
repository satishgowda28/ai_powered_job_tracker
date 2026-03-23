package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/respositories"
)

type UserService struct {
	userRepo *respositories.UserRepository
}

func NewUserService(userResp *respositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userResp,
	}
}

func (usrService *UserService) GetUserDetails(ctx context.Context, userId pgtype.UUID) (generated.User, error) {
	user, err := usrService.userRepo.GetUser(ctx, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return generated.User{}, errors.New("No user found")
		}
		return generated.User{}, err
	}
	return user, nil
}
