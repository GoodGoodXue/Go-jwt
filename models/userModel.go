package models

import (
	"Practice/Go-Projects/jwt/initializes"

	"gorm.io/gorm"
)

type MyUser struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func SyncDatabase() {
	initializes.DB.AutoMigrate(&MyUser{})
}
