package dto

//import "gorm.io/gorm"

type ClientDto struct {
	//gorm.Model
	ID_client  int    `json:"id"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      int    `json:"phone"`
}

type ClientsDto []ClientDto

//agregar gorm.model
