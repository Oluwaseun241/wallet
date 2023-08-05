package config

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

  if err != nil {
    return "", fmt.Errorf("could not hash password %w", err)
  }
  return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, userPassword string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
}
