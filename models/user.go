package models

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}
