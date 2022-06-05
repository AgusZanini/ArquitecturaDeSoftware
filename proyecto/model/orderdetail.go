package model

type OrderDetail struct {
	ID_orderdetail		int		`gorm:"primaryKey"`
	Products			Products
	Order				Order
	Quantity			int		`gorm:"type:int"`
}

type OrderDetails []OrderDetail