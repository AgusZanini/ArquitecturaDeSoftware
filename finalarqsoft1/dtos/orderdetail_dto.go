package dtos

type Orderdetaildto struct {
	Orderdetailid int     `json:"orderdetailid"`
	Orderid       int     `json:"orderid"`
	Productid     int     `json:"productid"`
	Amount        int     `json:"amount"`
	Total         float32 `json:"total"`
}

type Orderdetailsdto []Orderdetaildto
