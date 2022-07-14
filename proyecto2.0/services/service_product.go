package services

import (
	client "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/utils/errors"
)

type productservice struct{}

type productserviceInterface interface {
	GetProductById(id int) (dto.ProductDto, errors.ApiError)
	GetProducts() (dto.ProductsDto, errors.ApiError)
	GetProductByCategory(id int) (dto.ProductsDto, errors.ApiError)
	SearchProduct(name string) (dto.ProductsDto, errors.ApiError)
}

var (
	ProductService productserviceInterface
)

func init() {
	ProductService = &productservice{}
}

func (s *productservice) GetProductById(id int) (dto.ProductDto, errors.ApiError) {
	var product model.Product = client.GetProductById(id)
	var productdto dto.ProductDto

	if product.IDproduct == 0 {
		return productdto, errors.NewBadRequestApiError("bad request")
	}

	productdto.IDproduct = product.IDproduct
	productdto.Baseprice = product.Baseprice
	productdto.Description = product.Description
	productdto.Internalcode = product.Internalcode
	productdto.Idcategory = product.Idcategory
	productdto.Name = product.Name
	productdto.Stock = product.Stock

	return productdto, nil
}

func (s *productservice) Getproducts() (dto.ProductsDto, errors.ApiError) {
	var products model.Products = client.GetProducts()
	var productsdto dto.ProductsDto

	if products == nil {
		return productsdto, errors.NewBadRequestApiError("nil products")
	}

	for _, product := range products {
		var productdto dto.ProductDto

		productdto.IDproduct = product.IDproduct
		productdto.Baseprice = product.Baseprice
		productdto.Description = product.Description
		productdto.Internalcode = product.Internalcode
		productdto.Idcategory = product.Idcategory
		productdto.Name = product.Name
		productdto.Stock = product.Stock

		productsdto = append(productsdto, productdto)
	}

	return productsdto, nil
}

func (s *productservice) GetProductByCategory(id int) (dto.ProductsDto, errors.ApiError) {
	var products model.Products = client.GetProductByCategory(id)
	var productsdto dto.ProductsDto

	if products == nil {
		return productsdto, errors.NewBadRequestApiError("nil products")
	}

	for _, product := range products {
		var productdto dto.ProductDto

		productdto.IDproduct = product.IDproduct
		productdto.Baseprice = product.Baseprice
		productdto.Description = product.Description
		productdto.Internalcode = product.Internalcode
		productdto.Idcategory = product.Idcategory
		productdto.Name = product.Name
		productdto.Stock = product.Stock

		productsdto = append(productsdto, productdto)
	}

	return productsdto, nil
}

func (s *productservice) SearchProduct(name string) (dto.ProductsDto, errors.ApiError) {
	var products model.Products = client.SearchProduct(name)
	var productsdto dto.ProductsDto

	if products == nil {
		return productsdto, errors.NewBadRequestApiError("nil products")
	}

	for _, product := range products {
		var productdto dto.ProductDto

		productdto.IDproduct = product.IDproduct
		productdto.Baseprice = product.Baseprice
		productdto.Description = product.Description
		productdto.Internalcode = product.Internalcode
		productdto.Idcategory = product.Idcategory
		productdto.Name = product.Name
		productdto.Stock = product.Stock

		productsdto = append(productsdto, productdto)
	}

	return productsdto, nil
}
