package model

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	ID_orderdetail int `gorm:"primaryKey;AUTO_INCREMENT"`
	//Product			Product
	//Order				Order
	ID_product int `gorm:"foreignKey:ID_product"`
	ID_order   int `gorm:"foreignKey:ID_order"`
	Quantity   int `gorm:"type:integer;not null"`
}

type OrderDetails []OrderDetail
