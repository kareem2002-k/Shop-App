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

	// assign connection to global variable TO USE IT IN OTHER PACKAGES
	DB = con

	// migrate models to create tables in database
	con.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderedItem{}, &models.Address{})
	// create foreign key for user table
	DB.Model(&models.User{}).Association("Address")

	// create foreign key for cart table
	DB.Model(&models.User{}).Association("Cart")
	DB.Model(&models.Cart{}).Association("Items")

	// create foreign key for order table
	DB.Model(&models.User{}).Association("Orders")
	DB.Model(&models.Order{}).Association("Items")

	// print message if connection successful

	fmt.Println("Database connected")

}
