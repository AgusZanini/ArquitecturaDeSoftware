package model

//import "gorm.io/gorm"

type Product struct {
	//gorm.Model
	IDproduct    int     `gorm:"primaryKey;AUTO_INCREMENT"`
	Name         string  `gorm:"type:varchar(20);not null"`
	Description  string  `gorm:"type:varchar(100);not null"`
	Internalcode int     `gorm:"type:int;not null"`
	Baseprice    float32 `gorm:"type:decimal(60,4);not null"`
	Stock        int     `gorm:"type:int;not null"`
	Idcategory   int     `gorm:"foreignKey:IDcategory"`
}

type Products []Product
