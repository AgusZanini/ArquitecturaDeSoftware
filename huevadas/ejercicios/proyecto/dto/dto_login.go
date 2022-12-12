package dto

type Credentials struct {
	Client   string `json:"client"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
