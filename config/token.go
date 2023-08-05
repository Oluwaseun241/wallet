package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)

  now := time.Now().UTC()
  claims := token.Claims.(jwt.MapClaims)

  claims["sub"] = payload
  claims["exp"] = now.Add(ttl).Unix()
  claims["iat"] = now.Unix()
  claims["nbf"] = now.Unix()

  tokenString, err := token.SignedString([]byte(secretJWTKey))

  if err != nil {
    return "", fmt.Errorf("Generating JWT Token failed: %w", err)
  }
  return tokenString, nil
}

func ValidateToken(token string) {

}
