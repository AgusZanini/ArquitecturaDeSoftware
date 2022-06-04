package model

type Client struct {
	ID_client  int    `gorm:"primarykey"`
	First_name string `gorm:"type:varchar(20);not null"`
	Last_name  string `gorm:"type:varchar(20);not null"`
	Email      string `gorm:"type:varchar(20);not null"`
	Password   string `gorm:"type:varchar(20);not null"`
	Phone      int    `gorm:"type:int;not null"`
}

type Clients []Client
