package dto

type OrderDetailDto struct {
	IDorderdetail int `json:"id"`
	//Product			Product
	//Order				Order
	IDproduct int `json:"idproduct"`
	IDorder   int `json:"idorder"`
	Quantity  int `json:"quantity"`
}

type OrderDetailsDto []OrderDetailDto
