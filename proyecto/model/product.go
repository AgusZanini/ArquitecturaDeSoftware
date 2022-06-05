package model

type Product struct {
	ID_product		int		`gorm:"primaryKey"`
	Description		string	`gorm:"type:varchar(100);not null"`
	Internal_code	int		`gorm:"type:int;not null"`
	Base_price		int		`gorm:"type:int;not null"`
	Stock			int		`gorm:"type:int;not null"`
}

type Products []Product