package config

import (
	"fmt"
	"time"
  "github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID uuid.UUID) (string, error) { 
  claims := jwt.MapClaims{
    "sub": userID.String(),
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  secretKey := []byte("manchester1")
  tokenString, err := token.SignedString(secretKey)
  if err != nil {
    return "", fmt.Errorf("Generating JWT Token failed: %w", err)
  }
  return tokenString, nil
}

// func ValidateToken(token string) {

