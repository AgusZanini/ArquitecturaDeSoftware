package model

//import "gorm.io/gorm"

type Order struct {
	//gorm.Model
	IDorder    int     `gorm:"primaryKey;AUTO_INCREMENT"`
	Totalprice float32 `gorm:"type:float32"`
	Discount   float32 `gorm:"type:float32"`
	State      bool    `gorm:"type:bool;default:false"`
	//Client		Client
	IDuser int `gorm:"type:foreignKey:IDuser"`
}

type Orders []Order
