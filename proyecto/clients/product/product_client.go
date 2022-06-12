package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Dbproduct *gorm.DB

func GetProductById(id int) model.Product {
	var producto model.Product

	Dbproduct.Where("id = ?", id).First(&producto)
	fmt.Println("producto: ", producto)

	return producto
}

func GetProducts() model.Products {
	var productos model.Products

	Dbproduct.Find(&productos)
	fmt.Println("productos: ", productos)

	return productos
}

func SearchProducts(name string) model.Products {
	var productos model.Products

	Dbproduct.Where("name = ?", name).Find(&productos)
	fmt.Println("productos: ", productos)

	return productos
}

func InsertProducts(producto model.Product) model.Product {
	result := Dbproduct.Create(&producto)

	if result.Error != nil {
		fmt.Println("error")
	}

	fmt.Println("producto agregado", producto.ID_product)
	return producto
}
