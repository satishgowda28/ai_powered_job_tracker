package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/satishgowda28/ai_powered_job_tracker/db/generated"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/auth"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/respositories"
)

type AuthService struct {
	userRepo *respositories.UserRepository
	rtknRepo *respositories.RefreshTokenRepository
}

func NewAuthService(repo *respositories.UserRepository, rtknRepo *respositories.RefreshTokenRepository) *AuthService {
	return &AuthService{
		userRepo: repo,
		rtknRepo: rtknRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (generated.User, error) {
	/* Hashing passowrd */
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return generated.User{}, err
	}

	/* creating a new user */
	email = strings.ToLower(email)

	user, err := s.userRepo.CreateUser(ctx, generated.CreateUserParams{Name: name, Email: email, PasswordHash: hashedPassword})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return generated.User{}, errors.New("user already present")
			}
		}
		return generated.User{}, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (generated.User, error) {
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return generated.User{}, errors.New("invalid email or password")
		}
		return generated.User{}, err
	}
	matched, err := auth.ComparePasswordHash(password, user.PasswordHash)
	if err != nil {
		return generated.User{}, err
	}
	if !matched {
		return generated.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *AuthService) NewRefreshToken(ctx context.Context, user_id pgtype.UUID) (generated.UserRefreshToken, error) {
	expiresAt := time.Now().UTC().Add(60 * 24 * time.Hour)
	token, err := auth.GenerateRefreshToken()
	if err != nil {
		return generated.UserRefreshToken{}, err
	}
	return s.rtknRepo.CreateRefreshToken(ctx, generated.CreateRefreshTokenParams{
		Token:     token,
		UserID:    user_id,
		ExpiresAt: pgtype.Timestamp{Time: expiresAt, Valid: true},
	})

}
