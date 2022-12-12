package domain

type Credentials struct {
	User		string	`json:"user"`
	Password	string	`json:"password"`
}

type Token struct {
	Token string `json:"token"`
}