package clients

import (
	"fmt"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DbProduct *gorm.DB

func GetProductById(id int) model.Product {
	var producto model.Product

	if DbProduct == nil {
		fmt.Println("nil product")
	}

	if result := DbProduct.Where("IDproduct = ?", id).First(&producto); result.Error != nil {
		fmt.Println("couldn't find product")
	}
	fmt.Println("Producto", producto)

	return producto
}

func GetProducts() model.Products {
	var productos model.Products

	if DbProduct == nil {
		fmt.Println("nil product")
	}

	if result := DbProduct.Find(&productos); result.Error != nil {
		fmt.Println("couldn't find products")
	}
	fmt.Println("productos", productos)

	return productos
}

func SearchProduct(name string) model.Products {
	var productos model.Products

	if DbProduct == nil {
		fmt.Println("nil product")
	}

	if result := DbProduct.Where("Name = ?", "%"+name+"%").Find(&productos); result.Error != nil {
		fmt.Println("couldn't find products")
	}
	fmt.Println("Productos", productos)

	return productos
}

func CheckStock(id int, req int) bool {
	var producto model.Product
	available := true
	if DbProduct == nil {
		fmt.Println("nil product")
	}

	if result := DbProduct.Where("Idproduct = ?", id).First(&producto); result.Error != nil {
		fmt.Println("couldn't find product")
	}
	fmt.Println("Stock", producto.Stock)

	if req > producto.Stock {
		available = false
	}

	return available
}

func UpdateStock(id int, sub int) {
	var producto model.Product

	if DbProduct == nil {
		fmt.Println("nil product")
	}

	if result := DbProduct.Model(&model.Product{}).Where("IDproduct = ?", id).Update("Stock", producto.Stock-sub); result.Error != nil {
		fmt.Println("couldn't finde product")
	}

	fmt.Println("Stock updated", producto.Stock)
}

func GetProductByCategory(id int) model.Products {
	var productos model.Products

	if DbProduct == nil {
		fmt.Println("nil products")
	}

	if result := DbProduct.Where("Idcategory", id).Find(&productos); result.Error != nil {
		fmt.Println("coulnd't find products by category")
	}
	fmt.Println("productos", productos)

	return productos
}
