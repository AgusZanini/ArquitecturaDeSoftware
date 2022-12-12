package dtos

type Orderrequestdto struct {
	Userid       int             `json:"userid"`
	Orderdetails Orderdetailsdto `json:"orderdetails"`
	Address      string          `json:"address"`
}
