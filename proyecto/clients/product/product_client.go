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

	Dbproduct.Where("ID_product = ?", id).First(&producto)
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

	Dbproduct.Where("Name = ?", name).Find(&productos)
	fmt.Println("productos: ", productos)

	return productos
}

func UpdateStock(id int, quantity int) model.Product {
	var product model.Product

	Dbproduct.Where("ID_product = ?", id).First(&product)
	Dbproduct.Model(&product).Where("ID_product = ?", id).Update("Stock", product.Stock-quantity)
	fmt.Println("producto: ", product)

	return product
}

func InsertProducts(producto model.Product) model.Product {
	result := Dbproduct.Create(&producto)

	if result.Error != nil {
		fmt.Println("error")
	}

	fmt.Println("producto agregado", producto.ID_product)
	return producto
}
