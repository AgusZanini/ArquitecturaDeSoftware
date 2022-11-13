package model

import "gorm.io/gorm"

//revisar lo de los hash y passwords

type Client struct {
	gorm.Model
	ID_client  int    `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string `gorm:"type:varchar(20);not null;unique"`
	First_name string `gorm:"type:varchar(20);not null"`
	Last_name  string `gorm:"type:varchar(20);not null"`
	Email      string `gorm:"type:varchar(20);not null"`
	Password   string `gorm:"type:varchar(20);not null"`
	Phone      int    `gorm:"type:int;not null"`
}

type Clients []Client
