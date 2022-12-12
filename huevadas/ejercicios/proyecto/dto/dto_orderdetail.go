package dto

//import "gorm.io/gorm"

type OrderdetailDto struct {
	//gorm.Model
	ID_orderdetail int `json:"id_orderdetail"`
	//Product	       ProductDto
	//Order          OrderDto
	ID_product int `json:"id_product"`
	ID_order   int `json:"id_order"`
	Quantity   int `json:"quantity"`
}

type OrderdetailsDto []OrderdetailDto
