package dtos

type Orderdto struct {
	Orderid int     `json:"orderid"`
	Userid  int     `json:"userid"`
	Date    string  `json:"date"`
	Address string  `json:"address"`
	Total   float32 `json:"total"`
}

type Ordersdto []Orderdto
