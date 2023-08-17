package controllers

import (
	"github.com/Oluwaseun241/wallet/auth"
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//Get all users
func GetUser(c *fiber.Ctx) error {
  var users []Models.User
  db.DB.Find(&users)
  var userResponses []Models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, Models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
  return c.Status(fiber.StatusOK).JSON(userResponses)
}

// Create New User
func NewUser(c *fiber.Ctx) error {
  user := &Models.User{
    ID: uuid.New(),
  }

  if err := c.BodyParser(user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }
  
  // Hashing
  hashedPassword, err := auth.HashPassword(user.Password)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to hash password",
    })
  }
  user.Password = hashedPassword

  // Validate data(db)
  user.Validate(db.DB)
  if db.DB.Error != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Validation failed",
      "details": db.DB.Error.Error(),
    })
  }

  result := db.DB.Create(&user)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create user",
    })
  }
  return c.Status(fiber.StatusCreated).JSON(user)
}

//Get by ID
func GetUserId(c* fiber.Ctx) error {
  userId := c.Params("userId")
  var user Models.User
  if err := db.DB.Select("*").Where("id=?",userId).First(&user).Error; err != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "User not found",
    })
  }
  userResponse := Models.UserResponse{
    ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
  }
  return c.Status(fiber.StatusOK).JSON(userResponse)
}

//Update User
func UpdateUser(c *fiber.Ctx) error {
  userId := c.Params("userId")
  var user Models.User

  result := db.DB.Find(&user, "id=?", userId)
  if result.Error != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "User not found",
    })
  }

  var updatedUser Models.User
  if err := c.BodyParser(&updatedUser); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }

  user.Name = updatedUser.Name
  user.Email = updatedUser.Email

  result = db.DB.Save(&user)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "User updated successfully",
  })
}
