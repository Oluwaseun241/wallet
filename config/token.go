package config

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)
const secretKey = "manchester1"

// Middleware
func AuthMiddleware(c *fiber.Ctx) error {
  token := c.Get("Authorization")
  if token == "" {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "error": "Unauthorized",
    })
  }

  claims, err := ValidateToken(token)
  if err != nil {
    fmt.Println(err)
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "error": "Unauthorized",
    })
  }

  userID, ok := claims["sub"].(string)
  if !ok {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
  }

  c.Locals("userID", userID)
  return c.Next()
}


func GenerateToken(userID uuid.UUID) (string, error) { 
  claims := jwt.MapClaims{
    "sub": userID.String(),
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
  tokenString, err := token.SignedString([]byte(secretKey))
  if err != nil {
    return "", fmt.Errorf("Generating JWT Token failed: %w", err)
  }
  return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(secretKey), nil
  })
  if err != nil{
    return nil, fmt.Errorf("token validation failed: %w", err)
  }
  
  // Check if the token is valid and has valid claims
  if !token.Valid {
    return nil, fmt.Errorf("token is invalid")
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    return nil, fmt.Errorf("invalid token claims")
  }
  return claims, nil
}
