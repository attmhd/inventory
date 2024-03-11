package entity

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	Role     int
}

type UserInput struct {
	Id       int
	Username string
	Password string
	Email    string
	Role     int
}
