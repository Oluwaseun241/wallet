package config

import(
  "gorm.io/gorm"
  "fmt"
  Models "github.com/Oluwaseun241/wallet/models"
  "github.com/glebarez/sqlite"
)

var DB *gorm.DB

func InitDB() {
  db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect to DB")
  }
  fmt.Println("DB connected")
  DB = db
  db.AutoMigrate(&Models.User{})
  fmt.Println("DB migrated")
}
