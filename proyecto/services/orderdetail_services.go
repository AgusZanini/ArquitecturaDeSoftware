package services

import (
	orderdetailClient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/orderdetail"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
)

type orderdetailservice struct{}

type orderdetailserviceInterface interface {
	GetOrderDetailById(id int) (dto.OrderdetailDto, errors.ApiError)
	GetOrderdetails() (dto.OrderdetailsDto, errors.ApiError)
	InsertOrderDetail(orderdetailDto dto.OrderdetailDto) (dto.OrderdetailDto, errors.ApiError)
}

var (
	orderdetailService orderdetailserviceInterface
)

func init() {
	orderdetailService = &orderdetailService{}
}

func (s *orderdetailservice) GetOrderDetailById(id int) (dto.OrderdetailDto, errors.ApiError) {
	var orderdetail model.OrderDetail = orderdetailClient.GetOrderDetailById(id)
	var orderdetailDto dto.OrderdetailDto

	if orderdetail.ID_orderdetail == 0 {
		return orderdetailDto, errors.NewBadRequestApiError("detalle de orden no encontrada")
	}

	orderdetailDto.ID_orderdetail = orderdetail.ID_orderdetail
	//orderdetailDto.Products = orderdetail.Products
	for _, product := range orderdetail.Products {
		var productDto dto.ProductDto

		productDto.ID_product = product.ID_product
		productDto.Base_price = product.Base_price
		productDto.Description = product.Description
		productDto.Internal_code = product.Internal_code
		productDto.Stock = product.Stock

		orderdetailDto.Products = append(orderdetailDto.Products, productDto)
	}
	orderdetailDto.Quantity = orderdetail.Quantity
	orderdetailDto.Order = orderdetailDto.Order

	return orderdetailDto, nil
}

func (s *orderdetailservice) GetOrderDetails() (dto.OrderdetailsDto, errors.ApiError) {
	var orderdetails model.OrderDetails = orderdetailClient.GetOrderDetails()
	var orderdetailsDto dto.OrderdetailsDto

	for _, orderdetail := range orderdetails {
		var orderdetailDto dto.OrderdetailDto
		orderdetailDto.ID_orderdetail = orderdetail.ID_orderdetail
		//orderdetailDto.Products = orderdetail.Products
		for _, product := range orderdetail.Products {
			var productDto dto.ProductDto

			productDto.ID_product = product.ID_product
			productDto.Base_price = product.Base_price
			productDto.Description = product.Description
			productDto.Internal_code = product.Internal_code
			productDto.Stock = product.Stock

			orderdetailDto.Products = append(orderdetailDto.Products, productDto)
		}
		orderdetailDto.Quantity = orderdetail.Quantity
		orderdetailDto.Order = orderdetailDto.Order

		orderdetailsDto = append(orderdetailsDto, orderdetailDto)
	}
}

func (s *orderdetailservice) InsertOrderDetail(orderdetailDto dto.OrderdetailDto) (dto.OrderdetailDto, errors.ApiError) {
	var orderdetail model.OrderDetail

}
