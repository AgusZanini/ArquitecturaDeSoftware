package models

type Orderdetail struct {
	Orderdetailid int     `gorm:"primaryKey;auto_increment"`
	Orderid       int     `gorm:"foreignKey:orderid"`
	Productid     int     `gorm:"foreignKey:productid"`
	Amount        int     `gorm:"type:int; not null"`
	Total         float32 `gorm:"type:decimal;not null"`
}

type Orderdetails []Orderdetail
