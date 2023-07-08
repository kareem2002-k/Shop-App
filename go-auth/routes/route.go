package routes

import (
	"github.com/gofiber/fiber/v2"

	"auth/controlers"
)

func Setup(app *fiber.App) {

	// Routes for auth controlers
	app.Post("/register", controlers.Register)
	app.Post("/login", controlers.Login)
	app.Get("/user", controlers.User)
	app.Post("/logout", controlers.Logout)
	app.Post("/addtocart", controlers.AddToCart)
	app.Get("/cart", controlers.GetCart)
	app.Post("/deletecart", controlers.RemoveFromCart)
	app.Get("/products", controlers.GetProducts)
	app.Get("/product/:id", controlers.GetProduct)
	app.Post("/order", controlers.CheckOut)
}
