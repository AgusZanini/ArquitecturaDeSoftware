package models

type Category struct {
	Categoryid  int    `gorm:"primaryKey;auto_increment"`
	Name        string `gorm:"type:varchar(20);unique;not null"`
	Description string `gorm:"type:varchar(100);not null"`
}

type Categories []Category
