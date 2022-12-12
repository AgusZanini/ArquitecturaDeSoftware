package dto

type OrderDetailDto struct {
	IDorderdetail int     `json:"id"`
	IDproduct     int     `json:"idproduct"`
	IDorder       int     `json:"idorder"`
	Quantity      int     `json:"quantity"`
	Unitaryprice  float32 `json:"unitaryprice"`
	Totalprice    float32 `json:"totalprice"`
}

type OrderDetailsDto []OrderDetailDto
