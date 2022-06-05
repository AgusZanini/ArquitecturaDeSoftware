package dto

type OrderDto struct {
	ID_cart     int     `json:"id"`
	Total_price int     `json:"Total_price"`
	Discount    float32 `json:"Discount"`
	State       bool    `json:"State"`
	Quantity    int     `json:"Quantity"`
	Client      ClientDto
}

type OrdersDto []OrderDto
