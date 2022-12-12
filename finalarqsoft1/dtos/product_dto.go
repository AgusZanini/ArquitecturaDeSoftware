package dtos

type Productdto struct {
	Productid   int     `json:"productid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
	Stock       int     `json:"stock"`
	Categoryid  int     `json:"categoryid"`
}

type Productsdto []Productdto
