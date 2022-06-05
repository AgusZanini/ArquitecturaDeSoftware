package dto

type ProductDto struct {
	ID_product    int    `json:"id_product"`
	Description   string `json:"description"`
	Internal_code int    `json:"internal_code"`
	Base_price    int    `json:"base_price"`
	Stock         int    `json:"stock"`
}

type ProductsDto []ProductDto
