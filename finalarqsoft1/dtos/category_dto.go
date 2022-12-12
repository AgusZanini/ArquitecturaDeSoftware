package dtos

type Categorydto struct {
	Categoryid  int    `json:"categoryid"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
}

type Categoriesdto []Categorydto
