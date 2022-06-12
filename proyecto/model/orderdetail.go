package model

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	ID_orderdetail int `gorm:"primaryKey"`
	//Product			Product
	//Order				Order
	ID_product int `gorm:"type:integer;not null"`
	ID_order   int `gorm:"type:integer;not null"`
	Quantity   int `gorm:"type:integer;not null"`
}

type OrderDetails []OrderDetail
