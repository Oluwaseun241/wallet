package controllers

import (
	"strings"

	"github.com/Oluwaseun241/wallet/auths"
	db "github.com/Oluwaseun241/wallet/config"
	Models "github.com/Oluwaseun241/wallet/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

func Demo(c *fiber.Ctx) error {
  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "status": "ok",
  })
}

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
  // user := &Models.User{
  //   ID: uuid.New(),
  // }

  var input Models.SignUpInput
  if err := c.BodyParser(&input); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Invalid request format",
    })
  }
  
  // Check if the password and password confirmation match
  if input.Password != input.PasswordConfirm {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Password and password confirmation do not match",
    })
  }

  // Hashing
  hashedPassword, err := auth.HashPassword(input.Password)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to hash password",
    })
  }

  newUser := Models.User {
    ID: uuid.New(),
    Name: input.Name,
    Email: input.Email,
    Password: hashedPassword,
    Verified: false,
  }

  // Validate data(db)
  newUser.Validate(db.DB)
  if db.DB.Error != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Validation failed",
      "details": db.DB.Error.Error(),
    })
  }

  var existingUser Models.User
  if err := db.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
    return c.Status(fiber.StatusConflict).JSON(fiber.Map{
      "error": "User with this email already exists",
    })
  }

  result := db.DB.Create(&newUser)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": "Failed to create user",
    })
  }
  
  // Generate verification code
  code := randstr.String(20)

  verification_code := auth.Encode(code)

  newUser.VerificationCode = verification_code

  db.DB.Save(newUser)

  var firstName = newUser.Name
  if strings.Contains(firstName, " ") {
    firstName = strings.Split(firstName, " ")[1]
  }

  // Send Email

  emailData := auth.EmailData {
    URL: "/verifyemail/",
    FirstName: firstName,
    Subject: "Your account verification code",
  }
  
  auth.SendEmail(&newUser, &emailData)

  userResponse := Models.UserResponse{
    ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
  }
  return c.Status(fiber.StatusCreated).JSON(userResponse)
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
