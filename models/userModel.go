package models

import (
	"errors"
	"time"

	"github.com/Oluwaseun241/wallet/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
  gorm.Model
  ID  uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
  Name  string `gorm:"type:varchar(50);not null"`
  Email string `gorm:"uniqueIndex;not null"`
  Password  string `gorm:"not null"`
  CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type SignInInput struct {
  Email string  `json:"email" binding:"required"`
  Password  string  `json:"password" binding:"required"`
}

// type UserResponse struct {
// 	ID        uuid.UUID `json:"id,omitempty"`
// 	Name      string    `json:"name,omitempty"`
// 	Email     string    `json:"email,omitempty"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

func (User) TableName() string {
  return "users"
}

//Custom validation
func (u *User) Validate(db *gorm.DB) {
  switch {
  case u.Name == "":
    db.AddError(errors.New("Name cannot be blank"))
  case len(u.Password) < 8:
    db.AddError(errors.New("Password must be at least 8 characters"))
  }
}


func Authenticate(db *gorm.DB, email, password string) (*User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// Validate password
	if err := auth.VerifyPassword(user.Password, password); err != nil {
		return nil, err
	}

	return &user, nil
}

