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
			Creationdate: time.Now().Format("2006-01-02"),
			// date of creation only for the first time
		}
		database.DB.Create(&cart)

	} else {

		// check if the product is already in the cart
		var cartItem models.CartItem
		// get the cart item from the database
		getItem := database.DB.Where("cart_id = ? AND product_id = ?", cart.CartID, productID).First(&cartItem)

		// if the product is already in the cart
		if getItem.Error != nil {

			if getItem.Error.Error() == "record not found" {
				// add the product to the cart
				cartItem = models.CartItem{
					CartID:    cart.CartID,
					ProductID: uint(productID),
					Quantity:  int(productQuntity),
				}

				fmt.Println(cartItem)
				fmt.Println(cartItem.CartID)
				fmt.Println("i am here")
				database.DB.Create(&cartItem)
			}
			// update the cart item quntity

		} else {

			cartItem.Quantity = int(cartItem.Quantity) + int(productQuntity)

			// update the product quntity

			database.DB.Model(&cartItem).Where("cart_id = ? AND product_id = ?", cartItem.CartID, productID).Updates(&cartItem)

		}

	}

	return c.JSON(fiber.Map{
		"message": "Product added to cart",
	})
}

func GetCart(c *fiber.Ctx) error {

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

	// get the cart from the database
	var cart models.Cart

	database.DB.Where("user_id = ?", claims.Issuer).Preload("Items").First(&cart)

	return c.JSON(cart)
}

func RemoveFromCart(c *fiber.Ctx) error {

	var data map[string]string

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
	productID, _ := strconv.ParseUint(data["product_id"], 10, 64)
	removedQuantity, _ := strconv.ParseUint(data["quntity"], 10, 64)

	// get the cart from the database
	var cart models.Cart

	database.DB.Where("user_id = ?", claims.Issuer).Preload("Items").First(&cart)

	// get the cart item from the database
	var cartItem models.CartItem

	database.DB.Where("cart_id = ? AND product_id = ?", cart.CartID, productID).First(&cartItem)

	// check if the product is in the cart
	if cartItem.CartID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Product is not in the cart",
		})
	}

	// if the quantity removed is less than or equal the quantity in the cart

	if removedQuantity <= uint64(cartItem.Quantity) {
		// update the cart item quntity
		cartItem.Quantity = cartItem.Quantity - int(removedQuantity)

		// update the cart item in the database
		if cartItem.Quantity == 0 {
			database.DB.Where("cart_id = ? AND product_id = ?", cart.CartID, productID).Delete(&cartItem)
		} else {
			database.DB.Model(&cartItem).Where("cart_id = ? AND product_id = ?", cart.CartID, productID).Updates(&cartItem)
		}

	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Quantity removed is greater than the quantity in the cart",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product removed from cart",
	})
}

func GetProducts(c *fiber.Ctx) error {

	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	var product models.Product

	database.DB.Where("product_id = ?", id).First(&product)

	return c.JSON(product)
}

func CheckOut(c *fiber.Ctx) error {

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

	// get the cart from the database
	var cart models.Cart

	database.DB.Where("user_id = ?", claims.Issuer).Preload("Items").First(&cart)

	// check if the cart is empty
	if len(cart.Items) == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Cart is empty",
		})
	}

	// Check if the user has valid address
	var address models.Address

	database.DB.Where("user_id = ?", claims.Issuer).First(&address)

	if address.UserID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Address not found",
		})
	}

	UserID, _ := strconv.ParseInt(claims.Issuer, 10, 64)

	// Initialize a new order
	order := models.Order{
		UserID:       uint(UserID),
		Creationdate: time.Now().Format("2006-01-02 15:04:05"),
		Status:       "Order Placed",
	}

	var TotalPrice float64 = 0

	for _, cartItem := range cart.Items {
		orderItem := models.OrderedItem{
			ProductID: uint(cartItem.ProductID),
			Quantity:  cartItem.Quantity,
			Product:   cartItem.Product,
			UserID:    uint(UserID),
		}
		order.Items = append(order.Items, orderItem)
		TotalPrice = TotalPrice + (float64(cartItem.Quantity) * float64(cartItem.Product.Price))

	}
	order.Total = TotalPrice
	order.Address = address.Country + ", " + address.City + ", " + address.Street + ", " + address.Building + ", " + address.Apartment

	// save the order in the database
	database.DB.Create(&order)

	// Reset the cart and delete the cart items
	var cartItems models.CartItem

	database.DB.Where("cart_id = ?", cart.CartID).Delete(&cartItems)

	return c.JSON(fiber.Map{
		"message": "Order Placed",
	})

}
