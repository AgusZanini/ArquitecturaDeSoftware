package models

type Order struct {
	Orderid int     `gorm:"primaryKey;auto_increment"`
	Userid  int     `gorm:"foreignKey:id"`
	Date    string  `gorm:"type:date;not null;"`
	Address string  `gorm:"type:varchar(50); not null"`
	Total   float32 `gorm:"type:decimal; not null"`
}

type Orders []Order
