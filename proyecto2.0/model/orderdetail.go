package model

//import "gorm.io/gorm"

type OrderDetail struct {
	//gorm.Model
	IDorderdetail int `gorm:"primaryKey;AUTO_INCREMENT"`
	//Product			Product
	//Order				Order
	IDproduct int `gorm:"foreignKey:IDproduct"`
	IDorder   int `gorm:"foreignKey:IDorder"`
	Quantity  int `gorm:"type:integer;not null"`
}

type OrderDetails []OrderDetail
