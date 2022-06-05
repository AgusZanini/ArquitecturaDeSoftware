package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetProductById(id int) model.Product {
	var producto model.Product

	Db.Where("id = ?", id).First(&producto)
	fmt.Println("Client: ", producto)

	return producto
}

func GetProducts() model.Products {
	var productos model.Products

	Db.Find(&productos)
	fmt.Println("productos: ", productos)

	return productos
}

func InsertProducts(producto model.Product) model.Product {
	result := Db.Create(&producto)

	if result.Error != nil {
		fmt.Println("error")
	}

	fmt.Println("producto agregado")
	return producto
}