package controlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"auth/database"
	"auth/models"
	"strconv"
	"time"
)

func AddToCart(c *fiber.Ctx) error {

	var data map[string]string

	// parse request body to get product id

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	token, autheror := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if autheror != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	// get user id from token
	claims := token.Claims.(*jwt.StandardClaims)

	// get product id from request body
	productID, productidconv := strconv.ParseUint(data["product_id"], 10, 64)
	productQuntity, productquntityconv := strconv.ParseUint(data["quntity"], 10, 64)

	if productidconv != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Could not add to cart",
		})
	}

	if productquntityconv != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Could not add to cart",
		})

	}

	// get the product from the database
	var product models.Product

	database.DB.Where("product_id = ?", productID).First(&product)

	// check the product quantity
	if product.Quntity < 1 {
		c.Status(fiber.StatusBadRequest)
		fmt.Println(product.Quntity)
		return c.JSON(fiber.Map{
			"message": "Product is out of stock",
		})

	}

	// check if user quntity is avilable
	if product.Quntity < int(productQuntity) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Product is out of stock",
		})
	}

	// add cart item to cart
	var cart models.Cart

	database.DB.Where("user_id = ?", claims.Issuer).Preload("Items").First(&cart)

	uID, _ := strconv.ParseUint(claims.Issuer, 10, 64)

	// check if the user has a cart
	if cart.UserID == 0 {
		cart = models.Cart{
			UserID:       uint(uID),
			Creationdate: time.Now().Format("2006-01-02"), // date of creation only for the first time
		}
		database.DB.Create(&cart)

	}

	// check if the product is already in the cart
	var cartItem models.CartItem

	database.DB.Where("cart_id = ? AND product_id = ?", cart.CartID, productID).First(&cartItem)

	// if the product is already in the cart
	if cartItem.CartID != 0 {
		// update the cart item quntity
		cartItem.Quantity = int(cartItem.Quantity) + int(productQuntity)

		// update the product quntity

		database.DB.Model(&cartItem).Where("cart_id = ? AND product_id = ?", cartItem.CartID, productID).Updates(&cartItem)
	} else {
		// add the product to the cart
		cartItem = models.CartItem{
			CartID:    cart.CartID,
			ProductID: uint(productID),
			Quantity:  int(productQuntity),
		}
		database.DB.Create(&cartItem)
	}

	return c.JSON(fiber.Map{
		"message": "Product added to cart",
	})
}
