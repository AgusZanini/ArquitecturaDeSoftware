package dto

type ClientDto struct {
	ID_client  int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      int    `json:"phone"`
}

type ClientsDto []ClientDto
