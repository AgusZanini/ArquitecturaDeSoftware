package clients

import (
	"fmt"
	"proyectoarqsoft/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

type categoryClient struct{}

type CategoryClientInterface interface {
	GetCategoryByName(name string) (models.Category, error)
	GetCategories() (models.Categories, error)
	InsertCategory(category models.Category) error
}

var CategoryClient CategoryClientInterface

func init() {
	CategoryClient = &categoryClient{}
}

func (c *categoryClient) InsertCategory(category models.Category) error {

	result := Db.Create(&category)

	if result.Error != nil {
		fmt.Println("no se pudo insertar la categoria", category)
		return result.Error
	}

	fmt.Println("se pudo insertar la categoria con exito", category)
	return nil
}

func (c *categoryClient) GetCategories() (models.Categories, error) {
	var categories models.Categories

	result := Db.Find(&categories)

	if result.Error != nil {
		return models.Categories{}, result.Error
	}

	return categories, nil
}

func (c *categoryClient) GetCategoryByName(name string) (models.Category, error) {

	var category models.Category

	result := Db.Where("name = ?", name).First(&category)

	if result.Error != nil {
		return models.Category{}, result.Error
	}

	return category, nil
}
