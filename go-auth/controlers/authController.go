package controlers

import (
	"auth/database"
	"auth/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	// parse body to data map (key-value) string-string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	// hash password with bcrypt
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// create new user with data from request
	user := models.User{
		Username: data["username"],
		Password: password,
		Email:    data["email"],
	}

	database.DB.Create(&user)

	return c.JSON(user)

}

// SecretKey for jwt token creation
var SecretKey = "secret"

func Login(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var user models.User

	// find user by email in database and assign the first outcome to user variable
	database.DB.Where("email = ?", data["email"]).First(&user)

	// if user not found then user.ID will be 0 (default value)
	if user.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// compare password from request with password from database
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// create jwt token with user id as issuer and 1 day expiration time
	// claims is a map[string]interface{} type (key-value)
	// jwt.NewWithClaims creates a new token with the given claims and the default signing method.
	// jwt.StandardClaims is used to add standard claims to the token

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            // convert int to string (int is not allowed)
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	// generate jwt token with secret key
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// create cookie with jwt token
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // 1 day
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

func User(c *fiber.Ctx) error {

	// get jwt token from cookie
	cookie := c.Cookies("jwt")

	// parse jwt token with secret key and get claims
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	// get claims from token
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	// find user by id in database and assign the first outcome to user variable
	// Since we are using the id as the issuer, we can use the id to find the user

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {

	// create empty cookie with same name and set expiration time to 1 hour ago
	// this will delete the cookie from browser since it is expired
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // means expired 1 hour ago
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
