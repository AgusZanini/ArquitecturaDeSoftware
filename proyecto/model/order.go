package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID_order    int     `gorm:"primaryKey;AUTO_INCREMENT"`
	Total_price float32 `gorm:"type:float32"`
	Discount    float32 `gorm:"type:float32"`
	State       bool    `gorm:"type:bool;default:false"`
	//Client		Client
	ID_client int `gorm:"type:foreignKey:ID_client"`
}

type Orders []Order
