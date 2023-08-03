package Models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
  "gorm.io/gorm"
)

type User struct {
  Id  uuid.UUID
  FirstName  string
  LastName  string
  Email string
  Password  string
}

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
  var users []User
  db.Find(&users)
  return c.JSON(users)
}
