package dtos

type Orderandetailsdto struct {
	Orderdto        Orderdto        `json:"order"`
	Orderdetailsdto Orderdetailsdto `json:"orderdetails"`
}

type Ordersandetailsdto []Orderandetailsdto
