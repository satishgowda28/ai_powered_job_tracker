package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/auth"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/respositories"
)

type AuthService struct {
	repo *respositories.UserRespository
}

func NewAuthService(repo *respositories.UserRespository) *AuthService {
	return &AuthService{
		repo: respositories.NewUserRepository(),
	}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (generated.User, error) {
	/* Hashing passowrd */
	hashedPassword, err := auth.Hashpassword(password)
	if err != nil {
		return generated.User{}, err
	}

	/* creating a new user */
	email = strings.ToLower(email)
	user, err := s.repo.CreateUser(ctx, generated.CreateUserParams{Name: name, Email: email, PasswordHash: hashedPassword})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return generated.User{}, errors.New("user alerady present")
			}
		}
		return generated.User{}, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (generated.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return generated.User{}, errors.New("user not found")
		}
		return generated.User{}, err
	}
	matched, err := auth.ComparePasswordHash(user.PasswordHash, password)
	if err != nil {
		return generated.User{}, err
	}
	if !matched {
		return generated.User{}, errors.New("email or password is wrong")
	}
	return user, nil
}
