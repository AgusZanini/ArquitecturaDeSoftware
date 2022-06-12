package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID_product    int     `gorm:"primaryKey"`
	Name          string  `gorm:"type:varchar(20);not null"`
	Description   string  `gorm:"type:varchar(100);not null"`
	Internal_code int     `gorm:"type:int;not null"`
	Base_price    float32 `gorm:"type:decimal(60,4);not null"`
	Stock         int     `gorm:"type:int;not null"`
}

type Products []Product
