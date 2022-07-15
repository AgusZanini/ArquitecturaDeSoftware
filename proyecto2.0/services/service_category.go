package services

import (
	client "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/utils/errors"
)

type categoryservice struct{}

type categoryserviceinterface interface {
	GetCategories() (dto.CategoriesDto, errors.ApiError)
}

var (
	CategoryService categoryserviceinterface
)

func init() {
	CategoryService = &categoryservice{}
}

func (s *categoryservice) GetCategories() (dto.CategoriesDto, errors.ApiError) {
	var categories model.Categories = client.GetCategories()
	var categoriesdto dto.CategoriesDto

	if categories == nil {
		return categoriesdto, errors.NewBadRequestApiError("nil categories")
	}

	for _, category := range categories {
		var categorydto dto.CategoryDto

		categorydto.IDcategory = category.IDcategory
		categorydto.Name = category.Name

		categoriesdto = append(categoriesdto, categorydto)
	}

	return categoriesdto, nil
}
