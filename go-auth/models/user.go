package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}
