package Models

import (
  "time"
	"github.com/gofiber/fiber/v2"
	//"github.com/google/uuid"
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Id  uint `gorm:"primaryKey;autoIncrement;not null"`
  Name  string `gorm:"type:varchar(50);not null"`
  Email string `gorm:"uniqueIndex;not null"`
  Password  string `gorm:"not null"`
  CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

func (User) TableName() string {
  return "users"
}

//Custom validation
func (u *User) Validate(db *gorm.DB) {
  if u.Name == "" {
    db.AddError(gorm.ErrInvalidValue)
  }
  if len(u.Password) < 8 {
    db.AddError(gorm.ErrModelValueRequired)
  }
}

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
  var users []User
  db.Find(&users)
  return c.Status(200).JSON(users)
}

func NewUser(c *fiber.Ctx, db *gorm.DB) error {
  user := new(User)
  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  // Validate
  user.Validate(db)
  if db.Error != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Validation failed",
      "details": db.Error.Error(),
    })
  }

  result := db.Create(&user)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create user",
    })
  }
  return c.Status(fiber.StatusCreated).JSON(user)
}
