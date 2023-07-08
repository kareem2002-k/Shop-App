package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password []byte `json:"-"` // hide password from json response
	Email    string `json:"email" gorm:"unique"`
	cart     Cart   // ONE user has one carts

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
