package services

import (
	productClient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/product"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
)

type productservice struct{}

type productserviceInterface interface {
	GetProductById(id int) (dto.ProductDto, errors.ApiError)
	GetProducts() (dto.ProductsDto, errors.ApiError)
	SearchProducts(name string) (dto.ProductsDto, errors.ApiError)
	InsertProducts(productDto dto.ProductDto) (dto.ProductDto, errors.ApiError)
}

var (
	Productservice productserviceInterface
)

func init() {
	Productservice = &productservice{}
}

func (s *productservice) GetProductById(id int) (dto.ProductDto, errors.ApiError) {
	var product model.Product = productClient.GetProductById(id)
	var productDto dto.ProductDto

	if product.ID_product == 0 {
		return productDto, errors.NewBadRequestApiError("producto no encontrado")
	}

	productDto.ID_product = product.ID_product
	productDto.Name = product.Name
	productDto.Base_price = product.Base_price
	productDto.Description = product.Description
	productDto.Internal_code = product.Internal_code
	productDto.Stock = product.Stock

	return productDto, nil
}

func (s *productservice) GetProducts() (dto.ProductsDto, errors.ApiError) {
	var products model.Products = productClient.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.ID_product = product.ID_product
		productDto.Name = product.Name
		productDto.Base_price = product.Base_price
		productDto.Description = product.Description
		productDto.Internal_code = product.Internal_code
		productDto.Stock = product.Stock

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

func (s *productservice) SearchProducts(name string) (dto.ProductsDto, errors.ApiError) {
	var products model.Products = productClient.SearchProducts(name)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.ID_product = product.ID_product
		productDto.Name = product.Name
		productDto.Base_price = product.Base_price
		productDto.Description = product.Description
		productDto.Internal_code = product.Internal_code
		productDto.Stock = product.Stock

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

func (s *productservice) InsertProducts(productDto dto.ProductDto) (dto.ProductDto, errors.ApiError) {
	var product model.Product

	product.Base_price = productDto.Base_price
	product.Name = productDto.Name
	product.Description = productDto.Description
	product.Internal_code = productDto.Internal_code
	product.Stock = productDto.Stock

	product = productClient.InsertProducts(product)

	productDto.ID_product = product.ID_product

	return productDto, nil
}
