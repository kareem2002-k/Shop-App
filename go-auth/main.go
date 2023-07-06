package main

import (
	"auth/database"
	"auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to database
	database.Connect()

	// Create new fiber app
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Routes
	routes.Setup(app)

	app.Listen(":8000")

}
