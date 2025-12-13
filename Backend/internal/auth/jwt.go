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

// GenerateAccessToken creates a signed JWT access token for the given user ID.
// The token contains issuer, issued-at, subject (user UUID) and expires 15 minutes after issuance.
// It returns the signed token string, or an error if signing with the configured secret fails.
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

// ValidateToken parses and validates a JWT string and returns the user UUID contained in its subject claim.
// On success it returns the parsed UUID. If token parsing fails, the token is not valid, or the subject claim cannot be parsed as a UUID, it returns uuid.Nil and a descriptive error.
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

// GenerateRefreshToken generates a 64-character hex-encoded token derived from 32 bytes of cryptographically secure random data.
// If secure random generation fails it returns an empty string and a nil error.
func GenerateRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", nil
	}
	refreshToken := hex.EncodeToString([]byte(key))
	return refreshToken, nil
}