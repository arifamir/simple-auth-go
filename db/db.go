package db

import (
	"github.com/arifamir/simple-auth-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:trade12#@/simple_auth"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Mysql database.")
	}

	DB = connection

	// Create a user table
	connection.AutoMigrate(&models.User{})
}
