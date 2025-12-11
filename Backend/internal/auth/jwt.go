package auth

import "github.com/google/uuid"

func GenerateAccessToken(userID uuid.UUID) (string, error) {
	return "", nil
}

func GenerateRefreshToken(userID uuid.UUID) (string, error) {
	return "", nil
}

func ValidateToken(token string) (uuid.UUID, error) {
	return uuid.Nil, nil
}
