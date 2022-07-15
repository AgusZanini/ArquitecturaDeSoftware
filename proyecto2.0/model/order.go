package model

type Order struct {
	IDorder    int     `gorm:"primaryKey;AUTO_INCREMENT"`
	Totalprice float32 `gorm:"type:float32"`
	State      bool    `gorm:"type:bool;default:false"`
	IDuser     int     `gorm:"type:foreignKey:IDuser"`
}

type Orders []Order
