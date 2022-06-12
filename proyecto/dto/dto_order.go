package dto

//import "gorm.io/gorm"

type OrderDto struct {
	//gorm.Model
	ID_order    int     `json:"id"`
	Total_price float32 `json:"Total_price"`
	Discount    float32 `json:"Discount"`
	State       bool    `json:"State"`
	Quantity    int     `json:"Quantity"`
	//Client	ClientDto
	ID_client int `json:"id_client"`
}

type OrdersDto []OrderDto
