package clients

import (
	"fmt"
	"proyectoarqsoft/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

type productClient struct{}

type ProductClientInterface interface {
	InsertProduct(product models.Product) error
	SearchByName(name string) (models.Products, error)
	GetProductsByCategory(id int) (models.Products, error)
	UpdateStock(id int, amount int) (models.Product, error)
	GetproductById(id int) (models.Product, error)
}

var ProductClient ProductClientInterface

func init() {
	ProductClient = &productClient{}
}

func (p *productClient) InsertProduct(product models.Product) error {

	result := Db.Create(&product)

	if result.Error != nil {
		fmt.Printf("no se pudo cargar el producto nro %v", product.Productid)
		return result.Error
	}

	fmt.Println("Producto cargado con exito", product)

	return nil
}

func (p *productClient) SearchByName(name string) (models.Products, error) {
	var products models.Products

	result := Db.Where("name LIKE ?", "%"+name+"%").Find(&products)

	if result.Error != nil {
		return models.Products{}, result.Error
	}

	return products, nil
}

func (p *productClient) GetProductsByCategory(id int) (models.Products, error) {
	var products models.Products

	result := Db.Where("Categoryid = ?", id).Find(&products)

	if result.Error != nil {
		return models.Products{}, result.Error
	}

	return products, nil
}

func (p *productClient) GetproductById(id int) (models.Product, error) {
	var product models.Product

	result := Db.Where("Productid = ?", id).First(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

func (p *productClient) UpdateStock(id int, amount int) (models.Product, error) {

	var product models.Product

	if result := Db.Where("Productid = ?", id).First(&product); result.Error != nil {
		fmt.Println("No se pudo traer el product")
		return models.Product{}, result.Error
	}

	result := Db.Model(&product).Where("Productid = ?", product.Productid).Update("Stock", product.Stock-amount)

	if result.Error != nil {
		fmt.Println("error al actualizar el producto")
		return models.Product{}, result.Error
	}

	return product, nil

}
