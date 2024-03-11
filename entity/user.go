package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}

type UserInput struct {
	Id       int
	Username string
	Password string
	Email    string
	Role     int
}
