package config

import(
  "gorm.io/gorm"
  "fmt"
  "github.com/Oluwaseun241/wallet/Models"
  "github.com/glebarez/sqlite"
)
func initDB() *gorm.DB{
  db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
  //validations.RegisterCallbacks(db)
  if err != nil {
    panic("failed to connect to DB")
  }
  fmt.Println("DB connected")
  db.AutoMigrate(&Models.User{})
  fmt.Println("DB migrated")
  return db
}
