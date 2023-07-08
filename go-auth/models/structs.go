package models

type User struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Username string  `json:"username"`
	Password []byte  `json:"-"` // hide password from json response
	Email    string  `json:"email" gorm:"unique"`
	cart     Cart    // ONE user has one carts
	Orders   []Order // ONE user has many orders
	address  Address // ONE user has one address
}

type Product struct {
	ProductID   uint   `json:"product_id" gorm:"primary_key"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quntity     int    `json:"quntity"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Categories  string `json:"categories"`
}

type Cart struct {
	CartID       uint       `json:"cart_id" gorm:"primary_key"`
	UserID       uint       `json:"user_id"`
	Creationdate string     `json:"creationdate"`
	Items        []CartItem `json:"cart_items"` // one to many relationship means one cart has many cart items
}

type CartItem struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	CartID    uint    `json:"cart_id"`
	Product   Product `json:"product"` // one to one relationship means one cart item has one product
}

type Order struct {
	OrderID      uint          `json:"order_id" gorm:"primary_key"`
	UserID       uint          `json:"user_id"`
	Creationdate string        `json:"creationdate"`
	Items        []OrderedItem `json:"ordered_items"` // one to many relationship means one order has many ordered items
	Status       string        `json:"status"`
	Total        float64       `json:"total"`
	Address      string        `json:"address"`
}

type OrderedItem struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderID   uint    `json:"order_id"`
	UserID    uint    `json:"user_id"`
	Product   Product `json:"product"` // one to one relationship means one ordered item has one product

}

type Address struct {
	AddressID uint   `json:"address_id" gorm:"primary_key"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	Building  string `json:"building"`
	Apartment string `json:"apartment"`
	UserID    uint   `json:"user_id"`
}
