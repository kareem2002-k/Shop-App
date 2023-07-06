package database

import (
	"auth/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// grom database connection
	con, err := gorm.Open(mysql.Open("root:12345678@/shop"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = con

	con.AutoMigrate(&models.User{})
	fmt.Println("Database connected", con)

}
