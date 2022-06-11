package dto

type Credentials struct {
	Client   string `json:"client"`
	Password string `json:"password"`
}

type Hash struct {
	Hash string `json:"token"`
}
