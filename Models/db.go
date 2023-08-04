package Models

import (
  "time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Id  uuid.UUID `gorm:"type:uuid;primary_key"`
  Name  string `gorm:"type:varchar(50);not null"`
  Email string `gorm:"uniqueIndex;not null"`
  Password  string `gorm:"not null"`
  CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
  var users []User
  db.Find(&users)
  return c.Status(200).JSON(users)
}

func NewUser(c *fiber.Ctx, db *gorm.DB) error {
  user := new(User)
  if err := c.BodyParser(user); err != nil {
    return c.Status(503).SendString(err.Error())
  }
  db.Create(&user)
  return c.Status(201).JSON(user)
}
