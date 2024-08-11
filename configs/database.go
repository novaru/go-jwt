package configs

import (
  "go-jwt/models"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "log"
)

var DB *gorm.DB

func ConnectDB() {
  db, err := gorm.Open(mysql.Open("root:2323@tcp(127.0.0.0:3306)/go_jwt?parseTime=true"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  err = db.AutoMigrate(&models.User{})
  if err != nil {
    return
  }
  DB = db
  log.Println("Successfully connected to database")
}
