package main

import (
	"auth/database"
	"auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to the database
	database.Connect()

	// Create a new Fiber app
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowOrigins:     "http://localhost:3000", // Replace with your allowed origin(s)

	}))

	// Routes
	routes.Setup(app)

	app.Listen(":8000")
}
