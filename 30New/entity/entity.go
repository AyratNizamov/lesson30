package entity

type User struct {
	Name    string
	Age     int
	Friends []int
}

func NewUser() *User {
	var user User
	return &user
}
