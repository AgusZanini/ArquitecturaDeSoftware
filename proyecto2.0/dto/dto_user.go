package dto

type UserDto struct {
	IDuser    int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     int    `json:"phone"`
}

type UsersDto []UserDto
