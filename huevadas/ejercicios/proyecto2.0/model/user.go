package model

//import "gorm.io/gorm"

type User struct {
	//gorm.Model
	IDuser    int    `gorm:"primarykey;AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(20);not null;unique"`
	Firstname string `gorm:"type:varchar(20);not null"`
	Lastname  string `gorm:"type:varchar(20);not null"`
	Email     string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"type:varchar(20);not null"`
	Phone     int    `gorm:"type:int;not null"`
}

type Users []User
