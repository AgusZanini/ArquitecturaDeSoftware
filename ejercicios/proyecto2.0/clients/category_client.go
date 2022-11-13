package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"

	"gorm.io/gorm"
)

var Dbcategory *gorm.DB

func GetCategories() model.Categories {
	var categories model.Categories

	if result := Dbcategory.Find(&categories); result.Error != nil {
		fmt.Println("couldn't find categories")
	}

	return categories
}
