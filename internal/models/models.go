package models

// User is a bank user with some balance in wallet
type User struct {
	ID      int `json:"user_id"`
	Balance int `json:"balance"`
}
