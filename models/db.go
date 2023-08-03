package models

import (
	"github.com/google/uuid"
)

type User struct {
  Id  uuid.UUID
  FirstName  string
  LastName  string
  Email string
  Password  string
}
