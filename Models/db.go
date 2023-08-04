package Models

import (	
	"time"
  "errors"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Id  uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
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
  switch {
  case u.Name == "":
    db.AddError(errors.New("Name cannot be blank"))
  case len(u.Password) < 8:
    db.AddError(errors.New("Password must be at least 8 characters"))
  }
}
