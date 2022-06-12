package dto

//import "gorm.io/gorm"

type ProductDto struct {
	//gorm.Model
	ID_product    int     `json:"id_product"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Internal_code int     `json:"internal_code"`
	Base_price    float32 `gorm:"type:decimal(60,4);not null"`
	Stock         int     `json:"stock"`
}

type ProductsDto []ProductDto
