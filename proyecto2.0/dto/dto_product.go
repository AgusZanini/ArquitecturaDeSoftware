package dto

type ProductDto struct {
	IDproduct    int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Internalcode int     `json:"internalcode"`
	Baseprice    float32 `json:"baseprice"`
	Stock        int     `json:"stock"`
	Idcategory   int     `json:"idcategory"`
}

type ProductsDto []ProductDto
