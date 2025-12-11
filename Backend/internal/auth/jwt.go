package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/config"
)

func GenerateAccessToken(userID uuid.UUID) (string, error) {
	currentTime := time.Now().UTC()
	cfg := config.Get()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    cfg.Issuer,
		IssuedAt:  jwt.NewNumericDate(currentTime),
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(15 * time.Minute)),
		Subject:   userID.String(),
	})
	jwtToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func ValidateToken(token string) (uuid.UUID, error) {
	claim := &jwt.RegisteredClaims{}
	cfg := config.Get()
	jwtToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (any, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	if !jwtToken.Valid {
		return uuid.Nil, errors.New("not authorized")
	}
	userId, err := uuid.Parse(claim.Subject)
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}

func GenerateRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", nil
	}
	refreshToken := hex.EncodeToString([]byte(key))
	return refreshToken, nil
}
