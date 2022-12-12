package services

import (
	catclient "proyectoarqsoft/clients/categories"
	clients "proyectoarqsoft/clients/products"
	"proyectoarqsoft/dtos"
	"proyectoarqsoft/models"
	"proyectoarqsoft/utils/errors"
)

type productService struct {
	productClient clients.ProductClientInterface
}

type productServiceInterface interface {
	InsertProduct(productdto dtos.Productdto) errors.ApiError
	SearchByName(name string) (dtos.Productsdto, errors.ApiError)
	GetProductsByCategory(name string) (dtos.Productsdto, errors.ApiError)
}

var (
	ProductService productServiceInterface
)

func initProductService(productClient clients.ProductClientInterface) productServiceInterface {
	service := new(productService)
	service.productClient = productClient
	return service
}

func init() {
	ProductService = initProductService(clients.ProductClient)
}

func (p *productService) InsertProduct(productdto dtos.Productdto) errors.ApiError {

	var product models.Product

	product.Productid = productdto.Productid
	product.Categoryid = productdto.Categoryid
	product.Description = productdto.Description
	product.Image = productdto.Image
	product.Name = productdto.Name
	product.Price = productdto.Price
	product.Stock = productdto.Stock

	err := p.productClient.InsertProduct(product)
	if err != nil {
		return errors.NewBadRequestApiError(err.Error())
	}

	return nil
}

func (p *productService) SearchByName(name string) (dtos.Productsdto, errors.ApiError) {

	products, err := p.productClient.SearchByName(name)
	var productsdto dtos.Productsdto

	if err != nil {
		return dtos.Productsdto{}, errors.NewBadRequestApiError(err.Error())
	}

	for _, product := range products {
		var productdto dtos.Productdto

		productdto.Categoryid = product.Categoryid
		productdto.Description = product.Description
		productdto.Image = product.Image
		productdto.Name = product.Name
		productdto.Price = product.Price
		productdto.Productid = product.Productid
		productdto.Stock = product.Stock

		productsdto = append(productsdto, productdto)
	}

	return productsdto, nil
}

func (p *productService) GetProductsByCategory(name string) (dtos.Productsdto, errors.ApiError) {
	var productsdto dtos.Productsdto
	category, err := catclient.CategoryClient.GetCategoryByName(name)

	if err != nil {
		return dtos.Productsdto{}, errors.NewBadRequestApiError(err.Error())
	}

	products, err := p.productClient.GetProductsByCategory(category.Categoryid)

	if err != nil {
		return dtos.Productsdto{}, errors.NewBadRequestApiError(err.Error())
	}

	for _, product := range products {
		var productdto dtos.Productdto

		productdto.Categoryid = product.Categoryid
		productdto.Description = product.Description
		productdto.Image = product.Image
		productdto.Name = product.Name
		productdto.Price = product.Price
		productdto.Productid = product.Productid
		productdto.Stock = product.Stock

		productsdto = append(productsdto, productdto)
	}

	return productsdto, nil
}
