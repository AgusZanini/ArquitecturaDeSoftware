package model

type Order struct {
	ID_order		int		`gorm:"primaryKey"`
	Total_price		int		`gorm:"type:int"`
	Discount		float32	`gorm:"type:float32"`
	State			bool	`gorm:"type:bool;default:false"`
	Client			Client
}

type Orders []Order