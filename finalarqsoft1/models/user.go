package models

type User struct {
	Id        int    `gorm:"primaryKey;auto_increment"`
	Username  string `gorm:"type:varchar(20);unique; not null"`
	Firstname string `gorm:"type:varchar(20); not null"`
	Lastname  string `gorm:"type:varchar(20); not null"`
	Password  string `gorm:"type:varchar(200); not null"`
}

type Users []User
