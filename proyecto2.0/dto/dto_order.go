package dto

type OrderDto struct {
	IDorder         int             `json:"id"`
	Totalprice      float32         `json:"totalprice"`
	State           bool            `json:"state"`
	IDuser          int             `json:"iduser"`
	OrderDetailsDto OrderDetailsDto `json:"orderdetails"`
}

type OrdersDto []OrderDto
