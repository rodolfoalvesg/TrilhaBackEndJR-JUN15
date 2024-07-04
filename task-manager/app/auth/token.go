package auth

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type InputToken struct {
	PublicID uuid.UUID
	Name     string
	Email    string
}

// GenerateToken generates a JWT token
func GenerateToken(ctx context.Context, input InputToken) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"public_id": input.PublicID,
		"name":      input.Name,
		"email":     input.Email,
		"exp":       expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
