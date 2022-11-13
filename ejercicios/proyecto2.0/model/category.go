package model

//import "gorm.io/gorm"

type Category struct {
	//gorm.Model
	IDcategory int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Name       string `gorm:"type:varchar(20);not null;unique"`
}

type Categories []Category
