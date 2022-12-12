package models

type Product struct {
	Productid   int     `gorm:"primaryKey;auto_increment"`
	Name        string  `gorm:"type:varchar(30); not null"`
	Description string  `gorm:"type:varchar(200); not null"`
	Price       float32 `gorm:"type:decimal; not null"`
	Image       string  `gorm:"type:varchar(300); not null"`
	Stock       int     `gorm:"type:int; not null"`
	Categoryid  int     `gorm:"foreignKey:Categoryid"`
}

type Products []Product
