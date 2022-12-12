package services

import (
	clients "proyectoarqsoft/clients/categories"
	"proyectoarqsoft/dtos"
	"proyectoarqsoft/models"
	"proyectoarqsoft/utils/errors"
)

type categoryService struct {
	categoryclient clients.CategoryClientInterface
}

type categoryServiceInterface interface {
	InsertCategory(categorydto dtos.Categorydto) errors.ApiError
}

var CategoryService categoryServiceInterface

func initCategoryService(categoryclient clients.CategoryClientInterface) categoryServiceInterface {
	service := new(categoryService)
	service.categoryclient = categoryclient
	return service
}

func init() {
	CategoryService = initCategoryService(clients.CategoryClient)
}

func (c *categoryService) InsertCategory(categorydto dtos.Categorydto) errors.ApiError {
	var category models.Category

	category.Categoryid = categorydto.Categoryid
	category.Description = categorydto.Description
	category.Name = categorydto.Name

	err := c.categoryclient.InsertCategory(category)

	if err != nil {
		return errors.NewBadRequestApiError(err.Error())
	}

	return nil
}
