package dto

type OrderdetailDto struct {
	ID_orderdetail int `json:"id_orderdetail"`
	Products       ProductsDto
	Order          OrderDto
	Quantity       int `json:"quantity"`
}

type OrderdetailsDto []OrderdetailDto
