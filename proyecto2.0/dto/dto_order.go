package dto

type Order struct {
	IDorder    int     `json:"id"`
	Totalprice float32 `json:"totalprice"`
	Discount   float32 `json:"discount"`
	State      bool    `json:"state"`
	//Client		Client
	IDuser int `json:"iduser"`
}

type Orders []Order
