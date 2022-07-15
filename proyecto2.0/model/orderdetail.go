package model

//import "gorm.io/gorm"

type OrderDetail struct {
	//gorm.Model
	IDorderdetail int     `gorm:"primaryKey;AUTO_INCREMENT"`
	IDproduct     int     `gorm:"foreignKey:IDproduct"`
	IDorder       int     `gorm:"foreignKey:IDorder"`
	Quantity      int     `gorm:"type:integer;not null"`
	Unitaryprice  float32 `gorm:"type;varchar(10)"`
	Totalprice    float32 `gorm:"type;varchar(10)"`
}

type OrderDetails []OrderDetail
