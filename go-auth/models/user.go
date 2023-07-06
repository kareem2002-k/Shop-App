package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"` // hide password from json response
	Email    string `json:"email" gorm:"unique"`
}
